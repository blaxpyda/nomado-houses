package service

import (
	"fmt"
	"nomado-houses/internal/models"
	"nomado-houses/internal/repository"
	"nomado-houses/internal/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService interface defines methods for authentication
type AuthService interface {
	Register(req *models.RegisterRequest) (*models.AuthResponse, error)
	Login(req *models.LoginRequest) (*models.AuthResponse, error)
	ValidateToken(tokenString string) (int, error)
	VerifyEmail(req *models.VerifyEmailRequest) error
	ResendVerification(req *models.ResendVerificationRequest) error
}

// authService implements AuthService
type authService struct {
	userRepo     repository.UserRepository
	emailService EmailService
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo:     userRepo,
		emailService: NewEmailService(),
	}
}

// Register registers a new user
func (s *authService) Register(req *models.RegisterRequest) (*models.AuthResponse, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.GetUserByEmail(req.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("user with this email already exists")
	}

	// Validate input
	if err := utils.ValidateEmail(req.Email); err != nil {
		return nil, err
	}
	if err := utils.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &models.User{
		Email:            req.Email,
		Password:         string(hashedPassword),
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Phone:            req.Phone,
		Role:             req.Role, // Set role from request
		EmailVerified:    false,
		VerificationCode: s.emailService.GenerateVerificationCode(),
	}

	// Validate role
	if !user.Role.IsValid() {
		user.Role = models.RoleUser // Default to user role
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Send verification email
	if err := s.emailService.SendVerificationEmail(user.Email, user.FirstName, user.VerificationCode); err != nil {
		// Log the error but don't fail registration
		fmt.Printf("Failed to send verification email: %v\n", err)
	}

	// Generate JWT token (user can login but some features may be restricted)
	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Clear password and verification code from response
	user.Password = ""
	user.VerificationCode = ""

	return &models.AuthResponse{
		Token:         token,
		User:          *user,
		EmailVerified: user.EmailVerified,
	}, nil
}

// Login authenticates a user
func (s *authService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Clear password from response
	user.Password = ""

	return &models.AuthResponse{
		Token:         token,
		User:          *user,
		EmailVerified: user.EmailVerified,
	}, nil
}

// ValidateToken validates a JWT token and returns user ID
func (s *authService) ValidateToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(float64); ok {
			return int(userID), nil
		}
	}

	return 0, fmt.Errorf("invalid token claims")
}

// generateToken generates a JWT token for a user
func (s *authService) generateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iat":     time.Now().Unix(), // Issued at
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // 1 hour
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// VerifyEmail verifies a user's email with the provided code
func (s *authService) VerifyEmail(req *models.VerifyEmailRequest) error {
	// Verify the email with the code
	if err := s.userRepo.VerifyEmail(req.Email, req.VerificationCode); err != nil {
		return fmt.Errorf("email verification failed: %w", err)
	}

	// Get user details for welcome email
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		// Email is verified but we couldn't get user details for welcome email
		return nil
	}

	// Send welcome email
	if err := s.emailService.SendWelcomeEmail(user.Email, user.FirstName); err != nil {
		// Log the error but don't fail verification
		fmt.Printf("Failed to send welcome email: %v\n", err)
	}

	return nil
}

// ResendVerification resends the verification email
func (s *authService) ResendVerification(req *models.ResendVerificationRequest) error {
	// Get user by email
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Check if already verified
	if user.EmailVerified {
		return fmt.Errorf("email is already verified")
	}

	// Generate new verification code
	newCode := s.emailService.GenerateVerificationCode()

	// Update verification code in database
	if err := s.userRepo.UpdateVerificationCode(req.Email, newCode); err != nil {
		return fmt.Errorf("failed to update verification code: %w", err)
	}

	// Send verification email
	if err := s.emailService.SendVerificationEmail(user.Email, user.FirstName, newCode); err != nil {
		return fmt.Errorf("failed to send verification email: %w", err)
	}

	return nil
}

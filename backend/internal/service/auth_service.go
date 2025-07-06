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
}

// authService implements AuthService
type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
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
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate JWT token
	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Clear password from response
	user.Password = ""

	return &models.AuthResponse{
		Token: token,
		User:  *user,
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
		Token: token,
		User:  *user,
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
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

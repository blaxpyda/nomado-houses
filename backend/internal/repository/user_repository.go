package repository

import (
	"database/sql"
	"fmt"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/models"
)

// UserRepository interface defines methods for user operations
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	UpdateVerificationCode(email, code string) error
	VerifyEmail(email, code string) error
}

// userRepository implements UserRepository
type userRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB, logger *logger.Logger) UserRepository {
	return &userRepository{db: db, logger: logger}
}

// CreateUser creates a new user
func (r *userRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (email, password, first_name, last_name, phone, email_verified, verification_code)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query, user.Email, user.Password, user.FirstName, user.LastName, user.Phone, false, user.VerificationCode).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Set default values for newly created user
	user.EmailVerified = false
	return nil
}

// GetUserByEmail retrieves a user by email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password, first_name, last_name, phone, email_verified, verification_code, created_at, updated_at
		FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.FirstName,
		&user.LastName, &user.Phone, &user.EmailVerified, &user.VerificationCode, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password, first_name, last_name, phone, email_verified, verification_code, created_at, updated_at
		FROM users WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.FirstName,
		&user.LastName, &user.Phone, &user.EmailVerified, &user.VerificationCode, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// UpdateUser updates a user
func (r *userRepository) UpdateUser(user *models.User) error {
	query := `
		UPDATE users 
		SET first_name = $1, last_name = $2, phone = $3
		WHERE id = $4`
	_, err := r.db.Exec(query, user.FirstName, user.LastName, user.Phone, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// DeleteUser deletes a user
func (r *userRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// UpdateVerificationCode updates the verification code for a user
func (r *userRepository) UpdateVerificationCode(email, code string) error {
	query := `
		UPDATE users 
		SET verification_code = $1, updated_at = CURRENT_TIMESTAMP
		WHERE email = $2`
	_, err := r.db.Exec(query, code, email)
	if err != nil {
		return fmt.Errorf("failed to update verification code: %w", err)
	}
	return nil
}

// VerifyEmail verifies a user's email with the provided code
func (r *userRepository) VerifyEmail(email, code string) error {
	query := `
		UPDATE users 
		SET email_verified = TRUE, verification_code = NULL, updated_at = CURRENT_TIMESTAMP
		WHERE email = $1 AND verification_code = $2`

	result, err := r.db.Exec(query, email, code)
	if err != nil {
		return fmt.Errorf("failed to verify email: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("invalid verification code or email")
	}

	return nil
}

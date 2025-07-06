package repository

import (
	"database/sql"
	"fmt"
	"nomado-houses/internal/models"
	"time"
)

// UserRepository interface defines methods for user operations
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

// userRepository implements UserRepository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser creates a new user
func (r *userRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (email, password, first_name, last_name, phone, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	now := time.Now()
	err := r.db.QueryRow(query, user.Email, user.Password, user.FirstName, user.LastName, user.Phone, now, now).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

// GetUserByEmail retrieves a user by email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password, first_name, last_name, phone, created_at, updated_at
		FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.FirstName,
		&user.LastName, &user.Phone, &user.CreatedAt, &user.UpdatedAt,
	)
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
		SELECT id, email, password, first_name, last_name, phone, created_at, updated_at
		FROM users WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.FirstName,
		&user.LastName, &user.Phone, &user.CreatedAt, &user.UpdatedAt,
	)
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
		SET first_name = $1, last_name = $2, phone = $3, updated_at = $4
		WHERE id = $5`

	now := time.Now()
	_, err := r.db.Exec(query, user.FirstName, user.LastName, user.Phone, now, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	user.UpdatedAt = now
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

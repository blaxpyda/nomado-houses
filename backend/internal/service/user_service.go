package service

import (
	"nomado-houses/internal/models"
	"nomado-houses/internal/repository"
)

// UserService interface defines methods for user operations
type UserService interface {
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	UpdateUserRole(userID int, role models.UserRole) error
	GetAllUsers() ([]models.User, error)
	GetUsersByRole(role models.UserRole) ([]models.User, error)
	DeleteUser(id int) error
}

// userService implements UserService
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

// GetUserByEmail retrieves a user by email
func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

// UpdateUser updates user information
func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}

// UpdateUserRole updates a user's role
func (s *userService) UpdateUserRole(userID int, role models.UserRole) error {
	return s.userRepo.UpdateUserRole(userID, role)
}

// GetAllUsers retrieves all users (admin only)
func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

// GetUsersByRole retrieves users by role
func (s *userService) GetUsersByRole(role models.UserRole) ([]models.User, error) {
	return s.userRepo.GetUsersByRole(role)
}

// DeleteUser deletes a user
func (s *userService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}

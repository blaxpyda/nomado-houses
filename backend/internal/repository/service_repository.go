package repository

import (
	"database/sql"
	"fmt"
	"nomado-houses/internal/models"
)

// ServiceRepository interface defines methods for service operations
type ServiceRepository interface {
	GetAllServices() ([]models.Service, error)
	GetServicesByCategory(category string) ([]models.Service, error)
	GetServiceByID(id int) (*models.Service, error)
	CreateService(service *models.Service) error
	UpdateService(service *models.Service) error
	DeleteService(id int) error
}

// serviceRepository implements ServiceRepository
type serviceRepository struct {
	db *sql.DB
}

// NewServiceRepository creates a new service repository
func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

// GetAllServices retrieves all services
func (r *serviceRepository) GetAllServices() ([]models.Service, error) {
	query := `
		SELECT id, name, category, description, image_url, is_active, created_at, updated_at
		FROM services
		WHERE is_active = true
		ORDER BY category, name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(
			&service.ID, &service.Name, &service.Category,
			&service.Description, &service.ImageURL, &service.IsActive,
			&service.CreatedAt, &service.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan service: %w", err)
		}
		services = append(services, service)
	}

	return services, nil
}

// GetServicesByCategory retrieves services by category
func (r *serviceRepository) GetServicesByCategory(category string) ([]models.Service, error) {
	query := `
		SELECT id, name, category, description, image_url, is_active, created_at, updated_at
		FROM services
		WHERE category = $1 AND is_active = true
		ORDER BY name`

	rows, err := r.db.Query(query, category)
	if err != nil {
		return nil, fmt.Errorf("failed to get services by category: %w", err)
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(
			&service.ID, &service.Name, &service.Category,
			&service.Description, &service.ImageURL, &service.IsActive,
			&service.CreatedAt, &service.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan service: %w", err)
		}
		services = append(services, service)
	}

	return services, nil
}

// GetServiceByID retrieves a service by ID
func (r *serviceRepository) GetServiceByID(id int) (*models.Service, error) {
	service := &models.Service{}
	query := `
		SELECT id, name, category, description, image_url, is_active, created_at, updated_at
		FROM services WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&service.ID, &service.Name, &service.Category,
		&service.Description, &service.ImageURL, &service.IsActive,
		&service.CreatedAt, &service.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("service not found")
		}
		return nil, fmt.Errorf("failed to get service: %w", err)
	}

	return service, nil
}

// CreateService creates a new service
func (r *serviceRepository) CreateService(service *models.Service) error {
	query := `
		INSERT INTO services (name, category, description, image_url, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query, service.Name, service.Category,
		service.Description, service.ImageURL, service.IsActive).Scan(
		&service.ID, &service.CreatedAt, &service.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	return nil
}

// UpdateService updates a service
func (r *serviceRepository) UpdateService(service *models.Service) error {
	query := `
		UPDATE services 
		SET name = $1, category = $2, description = $3, image_url = $4, is_active = $5, updated_at = CURRENT_TIMESTAMP
		WHERE id = $6`

	_, err := r.db.Exec(query, service.Name, service.Category,
		service.Description, service.ImageURL, service.IsActive, service.ID)
	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}

	return nil
}

// DeleteService deletes a service
func (r *serviceRepository) DeleteService(id int) error {
	query := `DELETE FROM services WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}
	return nil
}

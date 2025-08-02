package repository

import (
	"database/sql"
	"fmt"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/models"
)

// ServiceRepository interface defines methods for service operations
type ServiceRepository interface {
	GetAllServices() ([]models.Service, error)
	GetServicesByServiceType(serviceTypeID int) ([]models.Service, error)
	GetServicesByCategory(category string) ([]models.Service, error)
	GetServiceByID(id int) (*models.Service, error)
	CreateService(service *models.Service) error
	UpdateService(service *models.Service) error
	DeleteService(id int) error
}

// serviceRepository implements ServiceRepository
type serviceRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

// NewServiceRepository creates a new service repository
func NewServiceRepository(db *sql.DB, logger *logger.Logger) ServiceRepository {
	return &serviceRepository{db: db, logger: logger}
}

// GetAllServices retrieves all services
func (r *serviceRepository) GetAllServices() ([]models.Service, error) {
	query := `
		SELECT id, provider_id, service_type_id, name, description, price, availability, created_at, updated_at
		FROM services
		WHERE availability = true
		ORDER BY name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(
			&service.ID, &service.ProviderID, &service.ServiceTypeID,
			&service.Name, &service.Description, &service.Price, &service.Availability,
			&service.CreatedAt, &service.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan service: %w", err)
		}
		services = append(services, service)
	}

	return services, nil
}

// GetServicesByServiceType retrieves services by service type
func (r *serviceRepository) GetServicesByServiceType(serviceTypeID int) ([]models.Service, error) {
	query := `
		SELECT id, provider_id, service_type_id, name, description, price, availability, created_at, updated_at
		FROM services
		WHERE service_type_id = $1 AND availability = true
		ORDER BY name`

	rows, err := r.db.Query(query, serviceTypeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get services by service type: %w", err)
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(
			&service.ID, &service.ProviderID, &service.ServiceTypeID,
			&service.Name, &service.Description, &service.Price, &service.Availability,
			&service.CreatedAt, &service.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan service: %w", err)
		}
		services = append(services, service)
	}

	return services, nil
}

// GetServicesByCategory retrieves services by category name
func (r *serviceRepository) GetServicesByCategory(category string) ([]models.Service, error) {
	query := `
		SELECT s.id, s.provider_id, s.service_type_id, s.name, s.description, s.price, s.availability, s.created_at, s.updated_at
		FROM services s
		JOIN service_types st ON s.service_type_id = st.id
		WHERE st.name = $1 AND s.availability = true
		ORDER BY s.name`

	rows, err := r.db.Query(query, category)
	if err != nil {
		return nil, fmt.Errorf("failed to get services by category: %w", err)
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(
			&service.ID, &service.ProviderID, &service.ServiceTypeID,
			&service.Name, &service.Description, &service.Price, &service.Availability,
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
		SELECT id, provider_id, service_type_id, name, description, price, availability, created_at, updated_at
		FROM services WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&service.ID, &service.ProviderID, &service.ServiceTypeID,
		&service.Name, &service.Description, &service.Price, &service.Availability,
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
		INSERT INTO services (provider_id, service_type_id, name, description, price, availability)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query, service.ProviderID, service.ServiceTypeID,
		service.Name, service.Description, service.Price, service.Availability).Scan(
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
		SET provider_id = $1, service_type_id = $2, name = $3, description = $4, price = $5, availability = $6, updated_at = CURRENT_TIMESTAMP
		WHERE id = $7`

	_, err := r.db.Exec(query, service.ProviderID, service.ServiceTypeID,
		service.Name, service.Description, service.Price, service.Availability, service.ID)
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

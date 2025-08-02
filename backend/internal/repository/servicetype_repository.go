package repository

import (
	"database/sql"
	"fmt"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/models"
)

// ServiceTypeRepository interface defines methods for service type operations
type ServiceTypeRepository interface {
	GetAllServiceTypes() ([]models.ServiceType, error)
	GetServiceTypeByID(id int) (*models.ServiceType, error)
	CreateServiceType(serviceType *models.ServiceType) error
	UpdateServiceType(serviceType *models.ServiceType) error
	DeleteServiceType(id int) error
}

// serviceTypeRepository implements ServiceTypeRepository
type serviceTypeRepository struct {
	db *sql.DB
	logger *logger.Logger
}

// NewServiceTypeRepository creates a new service type repository
func NewServiceTypeRepository(db *sql.DB, logger *logger.Logger) ServiceTypeRepository {
	return &serviceTypeRepository{db: db, logger: logger}
}

// GetAllServiceTypes retrieves all service types
func (r *serviceTypeRepository) GetAllServiceTypes() ([]models.ServiceType, error) {
	query := `
	    SELECT id, name, description
		FROM service_types
		ORDER BY name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get service types: %w", err)
	}
	defer rows.Close()

	var serviceTypes []models.ServiceType
	for rows.Next() {
		var serviceType models.ServiceType
		err := rows.Scan(&serviceType.ID, &serviceType.Name, &serviceType.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan service type: %w", err)
		}
		serviceTypes = append(serviceTypes, serviceType)
	}
	return serviceTypes, nil
}

// GetServiceTypeByID retrieves a service type by ID
func (r *serviceTypeRepository) GetServiceTypeByID(id int) (*models.ServiceType, error) {
	serviceType := &models.ServiceType{}
	query := `
		SELECT id, name, description
		FROM service_types WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&serviceType.ID, &serviceType.Name, &serviceType.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("service type not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get service type: %w", err)
	}
	return serviceType, nil
}

// CreateServiceType creates a new service type
func (r *serviceTypeRepository) CreateServiceType(serviceType *models.ServiceType) error {
	query := `
		INSERT INTO service_types (name, description)
		VALUES ($1, $2)
		RETURNING id`

	err := r.db.QueryRow(query, serviceType.Name, serviceType.Description).Scan(&serviceType.ID)
	if err != nil {
		return fmt.Errorf("failed to create service type: %w", err)
	}
	return nil
}

// UpdateServiceType updates an existing service type
func (r *serviceTypeRepository) UpdateServiceType(serviceType *models.ServiceType) error {
	query := `
		UPDATE service_types 
		SET name = $1, description = $2
		WHERE id = $3`

	_, err := r.db.Exec(query, serviceType.Name, serviceType.Description, serviceType.ID)
	if err != nil {
		return fmt.Errorf("failed to update service type: %w", err)
	}
	return nil
}

// DeleteServiceType deletes a service type by ID
func (r *serviceTypeRepository) DeleteServiceType(id int) error {
	query := `
		DELETE FROM service_types WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete service type: %w", err)
	}
	return nil
}
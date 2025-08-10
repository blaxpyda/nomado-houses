package service

import (
	"nomado-houses/internal/models"
	"nomado-houses/internal/repository"
)

// ServiceTypeService interface defines methods for service type operations
type ServiceTypeService interface {
	GetAllServiceTypes() ([]models.ServiceType, error)
	GetServiceTypeByID(id int) (*models.ServiceType, error)
	CreateServiceType(serviceType *models.ServiceType) error
	UpdateServiceType(serviceType *models.ServiceType) error
	DeleteServiceType(id int) error
}

// serviceTypeService implements ServiceTypeService
type serviceTypeService struct {
	serviceTypeRepo repository.ServiceTypeRepository
}

// NewServiceTypeService creates a new service type service
func NewServiceTypeService(serviceTypeRepo repository.ServiceTypeRepository) ServiceTypeService {
	return &serviceTypeService{serviceTypeRepo: serviceTypeRepo}
}

// GetAllServiceTypes retrieves all service types
func (s *serviceTypeService) GetAllServiceTypes() ([]models.ServiceType, error) {
	return s.serviceTypeRepo.GetAllServiceTypes()
}

// GetServiceTypeByID retrieves a service type by ID
func (s *serviceTypeService) GetServiceTypeByID(id int) (*models.ServiceType, error) {
	return s.serviceTypeRepo.GetServiceTypeByID(id)
}

// CreateServiceType creates a new service type
func (s *serviceTypeService) CreateServiceType(serviceType *models.ServiceType) error {
	return s.serviceTypeRepo.CreateServiceType(serviceType)
}

// UpdateServiceType updates a service type
func (s *serviceTypeService) UpdateServiceType(serviceType *models.ServiceType) error {
	return s.serviceTypeRepo.UpdateServiceType(serviceType)
}

// DeleteServiceType deletes a service type
func (s *serviceTypeService) DeleteServiceType(id int) error {
	return s.serviceTypeRepo.DeleteServiceType(id)
}

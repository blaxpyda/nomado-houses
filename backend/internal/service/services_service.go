package service

import (
	"nomado-houses/internal/models"
	"nomado-houses/internal/repository"
)

// ServiceService interface defines methods for service operations
type ServiceService interface {
	GetAllServices() ([]models.Service, error)
	GetServicesByCategory(category string) ([]models.Service, error)
	GetServiceByID(id int) (*models.Service, error)
	CreateService(service *models.Service) error
	UpdateService(service *models.Service) error
	DeleteService(id int) error
}

// serviceService implements ServiceService
type serviceService struct {
	serviceRepo repository.ServiceRepository
}

// NewServiceService creates a new service service
func NewServiceService(serviceRepo repository.ServiceRepository) ServiceService {
	return &serviceService{serviceRepo: serviceRepo}
}

// GetAllServices retrieves all services
func (s *serviceService) GetAllServices() ([]models.Service, error) {
	return s.serviceRepo.GetAllServices()
}

// GetServicesByCategory retrieves services by category
func (s *serviceService) GetServicesByCategory(category string) ([]models.Service, error) {
	return s.serviceRepo.GetServicesByCategory(category)
}

// GetServiceByID retrieves a service by ID
func (s *serviceService) GetServiceByID(id int) (*models.Service, error) {
	return s.serviceRepo.GetServiceByID(id)
}

// CreateService creates a new service
func (s *serviceService) CreateService(service *models.Service) error {
	return s.serviceRepo.CreateService(service)
}

// UpdateService updates a service
func (s *serviceService) UpdateService(service *models.Service) error {
	return s.serviceRepo.UpdateService(service)
}

// DeleteService deletes a service
func (s *serviceService) DeleteService(id int) error {
	return s.serviceRepo.DeleteService(id)
}

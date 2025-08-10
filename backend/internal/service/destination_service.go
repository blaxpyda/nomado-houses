package service

import (
	"nomado-houses/internal/models"
	"nomado-houses/internal/repository"
)

type DestinationService interface {
	GetAllDestinations() ([]models.Destination, error)
	GetDestinationByID(id int) (*models.Destination, error)
	CreateDestination(destination *models.Destination) error
	UpdateDestination(destination *models.Destination) error
	DeleteDestination(id int) error
}

// destinationService implements DestinationService
type destinationService struct {
	destinationRepo repository.DestinationRepository
}

// NewDestinationService creates a new destination service
func NewDestinationService(destinationRepo repository.DestinationRepository) DestinationService {
	return &destinationService{destinationRepo: destinationRepo}
}

// GetAllDestinations retrieves all destinations
func (s *destinationService) GetAllDestinations() ([]models.Destination, error) {
	return s.destinationRepo.GetAllDestinations()
}

// GetDestinationByID retrieves a destination by ID
func (s *destinationService) GetDestinationByID(id int) (*models.Destination, error) {
	return s.destinationRepo.GetDestinationByID(id)
}

// CreateDestination creates a new destination
func (s *destinationService) CreateDestination(destination *models.Destination) error {
	return s.destinationRepo.CreateDestination(destination)
}

// UpdateDestination updates a destination
func (s *destinationService) UpdateDestination(destination *models.Destination) error {
	return s.destinationRepo.UpdateDestination(destination)
}

// DeleteDestination deletes a destination
func (s *destinationService) DeleteDestination(id int) error {
	return s.destinationRepo.DeleteDestination(id)
}

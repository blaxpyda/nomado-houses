package service

import (
	"nomado-houses/internal/models"
	"nomado-houses/internal/repository"
)

// DestinationService interface defines methods for destination operations
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

// BookingService interface defines methods for booking operations
type BookingService interface {
	CreateBooking(booking *models.Booking) error
	GetBookingsByUserID(userID int) ([]models.Booking, error)
	GetBookingByID(id int) (*models.Booking, error)
	UpdateBookingStatus(id int, status string) error
	DeleteBooking(id int) error
}

// bookingService implements BookingService
type bookingService struct {
	bookingRepo repository.BookingRepository
}

// NewBookingService creates a new booking service
func NewBookingService(bookingRepo repository.BookingRepository) BookingService {
	return &bookingService{bookingRepo: bookingRepo}
}

// CreateBooking creates a new booking
func (s *bookingService) CreateBooking(booking *models.Booking) error {
	return s.bookingRepo.CreateBooking(booking)
}

// GetBookingsByUserID retrieves bookings by user ID
func (s *bookingService) GetBookingsByUserID(userID int) ([]models.Booking, error) {
	return s.bookingRepo.GetBookingsByUserID(userID)
}

// GetBookingByID retrieves a booking by ID
func (s *bookingService) GetBookingByID(id int) (*models.Booking, error) {
	return s.bookingRepo.GetBookingByID(id)
}

// UpdateBookingStatus updates booking status
func (s *bookingService) UpdateBookingStatus(id int, status string) error {
	return s.bookingRepo.UpdateBookingStatus(id, status)
}

// DeleteBooking deletes a booking
func (s *bookingService) DeleteBooking(id int) error {
	return s.bookingRepo.DeleteBooking(id)
}

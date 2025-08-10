package service

import (
	"nomado-houses/internal/models"
	"nomado-houses/internal/repository"
)

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

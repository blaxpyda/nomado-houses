package repository

import (
	"database/sql"
	"fmt"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/models"
)

// BookingRepository interface defines methods for booking operations
type BookingRepository interface {
	CreateBooking(booking *models.Booking) error
	GetBookingsByUserID(userID int) ([]models.Booking, error)
	GetBookingByID(id int) (*models.Booking, error)
	UpdateBookingStatus(id int, status string) error
	DeleteBooking(id int) error
}

// bookingRepository implements BookingRepository
type bookingRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

// NewBookingRepository creates a new booking repository
func NewBookingRepository(db *sql.DB, logger *logger.Logger) BookingRepository {
	return &bookingRepository{db: db, logger: logger}
}

// CreateBooking creates a new booking
func (r *bookingRepository) CreateBooking(booking *models.Booking) error {
	query := `
		INSERT INTO bookings (user_id, service_id, booking_date_start, booking_date_end, total_price, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query, booking.UserID, booking.ServiceID,
		booking.BookingDateStart, booking.BookingDateEnd, booking.TotalPrice, booking.Status).Scan(
		&booking.ID, &booking.CreatedAt, &booking.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create booking: %w", err)
	}

	return nil
}

// GetBookingsByUserID retrieves bookings by user ID
func (r *bookingRepository) GetBookingsByUserID(userID int) ([]models.Booking, error) {
	query := `
		SELECT id, user_id, service_id, booking_date_start, booking_date_end, total_price, status, created_at, updated_at
		FROM bookings
		WHERE user_id = $1
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get bookings: %w", err)
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		err := rows.Scan(
			&booking.ID, &booking.UserID, &booking.ServiceID,
			&booking.BookingDateStart, &booking.BookingDateEnd,
			&booking.TotalPrice, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan booking: %w", err)
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// GetBookingByID retrieves a booking by ID
func (r *bookingRepository) GetBookingByID(id int) (*models.Booking, error) {
	booking := &models.Booking{}
	query := `
		SELECT id, user_id, service_id, booking_date_start, booking_date_end, total_price, status, created_at, updated_at
		FROM bookings WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&booking.ID, &booking.UserID, &booking.ServiceID,
		&booking.BookingDateStart, &booking.BookingDateEnd,
		&booking.TotalPrice, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("booking not found")
		}
		return nil, fmt.Errorf("failed to get booking: %w", err)
	}

	return booking, nil
}

// UpdateBookingStatus updates booking status
func (r *bookingRepository) UpdateBookingStatus(id int, status string) error {
	query := `UPDATE bookings SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`

	_, err := r.db.Exec(query, status, id)
	if err != nil {
		return fmt.Errorf("failed to update booking status: %w", err)
	}

	return nil
}

// DeleteBooking deletes a booking
func (r *bookingRepository) DeleteBooking(id int) error {
	query := `DELETE FROM bookings WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete booking: %w", err)
	}
	return nil
}

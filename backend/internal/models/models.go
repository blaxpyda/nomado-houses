package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        int    `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"-" db:"password"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Phone     string `json:"phone" db:"phone"`
}

// Service represents a category of services offered by the platform
type Service struct {
	ID            int       `json:"id" db:"id"`
	ProviderID    int       `json:"provider_id" db:"provider_id"`
	ServiceTypeID int       `json:"service_type_id" db:"service_type_id"`
	Name          string    `json:"name" db:"name"`
	Description   string    `json:"description" db:"description"`
	Price         float64   `json:"price" db:"price"`
	Availability  bool      `json:"availability" db:"availability"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// ServiceType represents a type of service offered by the platform
type ServiceType struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Booking represents a booking made by a user for any services offered by the platform
type Booking struct {
	ID               int       `json:"id" db:"id"`
	UserID           int       `json:"user_id" db:"user_id"`
	ServiceID        int       `json:"service_id" db:"service_id"`
	ProviderID       int       `json:"provider_id" db:"provider_id"`
	BookingDateStart time.Time `json:"booking_date_start" db:"booking_date_start"`
	BookingDateEnd   time.Time `json:"booking_date_end" db:"booking_date_end"`
	TotalPrice       float64   `json:"total_price" db:"total_price"`
	Status           string    `json:"status" db:"status"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

// Payment represents a payment made by a user for a booking
type Payment struct {
	ID            int       `json:"id" db:"id"`
	UserID        int       `json:"user_id" db:"user_id"`
	BookingID     int       `json:"booking_id" db:"booking_id"`
	Amount        float64   `json:"amount" db:"amount"`
	PaymentDate   time.Time `json:"payment_date" db:"payment_date"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	Status        bool      `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// Destination represents a destination for travel or service
type Destination struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Location    string    `json:"location" db:"location"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	Rating      float64   `json:"rating" db:"rating"`
	Reviews     int       `json:"reviews" db:"reviews"`
	Price       float64   `json:"price" db:"price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}


// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest represents the registration request payload
type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// CreateDestinationRequest represents the request to create a destination
type CreateDestinationRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
}

// UpdateDestinationRequest represents the request to update a destination
type UpdateDestinationRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
}

// CreateServiceRequest represents the request to create a service
type CreateServiceRequest struct {
	ProviderID    int     `json:"provider_id" validate:"required"`
	ServiceTypeID int     `json:"service_type_id" validate:"required"`
	Name          string  `json:"name" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	Price         float64 `json:"price" validate:"required,gt=0"`
	Availability  bool    `json:"availability"`
}

// UpdateServiceRequest represents the request to update a service
type UpdateServiceRequest struct {
	ProviderID    int     `json:"provider_id" validate:"required"`
	ServiceTypeID int     `json:"service_type_id" validate:"required"`
	Name          string  `json:"name" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	Price         float64 `json:"price" validate:"required,gt=0"`
	Availability  bool    `json:"availability"`
}

// CreateServiceTypeRequest represents the request to create a service type
type CreateServiceTypeRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// UpdateServiceTypeRequest represents the request to update a service type
type UpdateServiceTypeRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// CreateBookingRequest represents the request to create a booking
type CreateBookingRequest struct {
	ServiceID        int       `json:"service_id" validate:"required"`
	ProviderID       int       `json:"provider_id" validate:"required"`
	BookingDateStart time.Time `json:"booking_date_start" validate:"required"`
	BookingDateEnd   time.Time `json:"booking_date_end" validate:"required"`
	TotalPrice       float64   `json:"total_price" validate:"required,gt=0"`
}

// UpdateBookingStatusRequest represents the request to update booking status
type UpdateBookingStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=pending confirmed cancelled completed"`
}

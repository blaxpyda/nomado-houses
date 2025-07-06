package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Phone     string    `json:"phone" db:"phone"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Destination represents a travel destination
type Destination struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Country     string    `json:"country" db:"country"`
	City        string    `json:"city" db:"city"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	Rating      float64   `json:"rating" db:"rating"`
	DealsCount  int       `json:"deals_count" db:"deals_count"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Hotel represents a hotel accommodation
type Hotel struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	DestinationID int       `json:"destination_id" db:"destination_id"`
	Address       string    `json:"address" db:"address"`
	Description   string    `json:"description" db:"description"`
	ImageURL      string    `json:"image_url" db:"image_url"`
	Rating        float64   `json:"rating" db:"rating"`
	PricePerNight float64   `json:"price_per_night" db:"price_per_night"`
	Amenities     string    `json:"amenities" db:"amenities"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// Service represents a travel service
type Service struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Category    string    `json:"category" db:"category"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Booking represents a booking made by a user
type Booking struct {
	ID           int       `json:"id" db:"id"`
	UserID       int       `json:"user_id" db:"user_id"`
	ServiceType  string    `json:"service_type" db:"service_type"`
	ServiceID    int       `json:"service_id" db:"service_id"`
	CheckInDate  time.Time `json:"check_in_date" db:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date" db:"check_out_date"`
	TotalAmount  float64   `json:"total_amount" db:"total_amount"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
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

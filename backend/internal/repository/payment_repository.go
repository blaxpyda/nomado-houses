package repository

import (
	"database/sql"
	"fmt"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/models"
)

// PaymentRepository defines the interface for payment-related database operations
type PaymentRepository interface {
	GetAllPayments() ([]models.Payment, error)
	GetPaymentsByUserID(userID int) ([]models.Payment, error)
	GetPaymentByID(id int) (*models.Payment, error)
	CreatePayment(payment *models.Payment) error
	UpdatePayment(payment *models.Payment) error
	DeletePayment(id int) error
}

// paymentRepository implements PaymentRepository
type paymentRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

// NewPaymentRepository creates a new payment repository
func NewPaymentRepository(db *sql.DB, logger *logger.Logger) PaymentRepository {
	return &paymentRepository{db: db, logger: logger}
}

// GetAllPayments retrieves all payments
func (r *paymentRepository) GetAllPayments() ([]models.Payment, error) {
	query := `
		SELECT id, user_id, booking_id, amount, payment_date, payment_method, status, created_at, updated_at
		FROM payments
		ORDER BY payment_date DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get payments: %w", err)
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(
			&payment.ID, &payment.UserID, &payment.BookingID,
			&payment.Amount, &payment.PaymentDate, &payment.PaymentMethod,
			&payment.Status, &payment.CreatedAt, &payment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan payment: %w", err)
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

// GetPaymentsByUserID retrieves payments by user ID
func (r *paymentRepository) GetPaymentsByUserID(userID int) ([]models.Payment, error) {
	query := `
		SELECT id, user_id, booking_id, amount, payment_date, payment_method, status, created_at, updated_at
		FROM payments
		WHERE user_id = $1
		ORDER BY payment_date DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payments by user ID: %w", err)
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(
			&payment.ID, &payment.UserID, &payment.BookingID,
			&payment.Amount, &payment.PaymentDate, &payment.PaymentMethod,
			&payment.Status, &payment.CreatedAt, &payment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan payment: %w", err)
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

// GetPaymentByID retrieves a payment by ID
func (r *paymentRepository) GetPaymentByID(id int) (*models.Payment, error) {
	payment := &models.Payment{}
	query := `
		SELECT id, user_id, booking_id, amount, payment_date, payment_method, status, created_at, updated_at
		FROM payments WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&payment.ID, &payment.UserID, &payment.BookingID,
		&payment.Amount, &payment.PaymentDate, &payment.PaymentMethod,
		&payment.Status, &payment.CreatedAt, &payment.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("payment not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}
	return payment, nil
}

// CreatePayment creates a new payment
func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	query := `
		INSERT INTO payments (user_id, booking_id, amount, payment_date, payment_method, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query, payment.UserID, payment.BookingID,
		payment.Amount, payment.PaymentDate, payment.PaymentMethod,
		payment.Status).Scan(&payment.ID, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}
	return nil
}

// UpdatePayment updates an existing payment
func (r *paymentRepository) UpdatePayment(payment *models.Payment) error {
	query := `
		UPDATE payments 
		SET amount = $1, payment_date = $2, payment_method = $3, status = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $5`

	_, err := r.db.Exec(query, payment.Amount, payment.PaymentDate,
		payment.PaymentMethod, payment.Status, payment.ID)
	if err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}
	return nil
}

// DeletePayment deletes a payment
func (r *paymentRepository) DeletePayment(id int) error {
	query := `DELETE FROM payments WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete payment: %w", err)
	}
	return nil
}

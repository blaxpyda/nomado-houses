package repository

import (
	"database/sql"
	"fmt"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/models"
)

// DestinationRepository interface defines methods for destination operations
type DestinationRepository interface {
	GetAllDestinations() ([]models.Destination, error)
	GetDestinationByID(id int) (*models.Destination, error)
	CreateDestination(destination *models.Destination) error
	UpdateDestination(destination *models.Destination) error
	DeleteDestination(id int) error
}

// destinationRepository implements DestinationRepository
type destinationRepository struct {
	db *sql.DB
	logger *logger.Logger
}

// NewDestinationRepository creates a new destination repository
func NewDestinationRepository(db *sql.DB, logger *logger.Logger) DestinationRepository {
	return &destinationRepository{db: db, logger: logger}
}

// GetAllDestinations retrieves all destinations
func (r *destinationRepository) GetAllDestinations() ([]models.Destination, error) {
	query := `
		SELECT id, name, description, location, image_url, rating, reviews, price, created_at, updated_at
		FROM destinations
		ORDER BY rating DESC, reviews DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get destinations: %w", err)
	}
	defer rows.Close()

	var destinations []models.Destination
	for rows.Next() {
		var dest models.Destination
		err := rows.Scan(
			&dest.ID, &dest.Name, &dest.Description, &dest.Location,
			&dest.ImageURL, &dest.Rating, &dest.Reviews, &dest.Price,
			 &dest.CreatedAt, &dest.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan destination: %w", err)
		}
		destinations = append(destinations, dest)
	}

	return destinations, nil
}

// GetDestinationByID retrieves a destination by ID
func (r *destinationRepository) GetDestinationByID(id int) (*models.Destination, error) {
	destination := &models.Destination{}
	query := `
		SELECT id, name, description, location, image_url, rating, reviews, price, created_at, updated_at
		FROM destinations WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&destination.ID, &destination.Name, &destination.Description, &destination.Location,
		&destination.ImageURL, &destination.Rating, &destination.Reviews, &destination.Price,
		&destination.CreatedAt, &destination.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("destination not found")
		}
		return nil, fmt.Errorf("failed to get destination: %w", err)
	}

	return destination, nil
}

// CreateDestination creates a new destination
func (r *destinationRepository) CreateDestination(destination *models.Destination) error {
	query := `
		INSERT INTO destinations (name, description, location, image_url, rating, reviews, price)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query, destination.Name, destination.Description, destination.Location, destination.ImageURL, destination.Rating, destination.Reviews, destination.Price).Scan(
		&destination.ID, &destination.CreatedAt, &destination.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create destination: %w", err)
	}

	return nil
}

// UpdateDestination updates a destination
func (r *destinationRepository) UpdateDestination(destination *models.Destination) error {
	query := `
		UPDATE destinations 
		SET name = $1, description = $2, location = $3, image_url = $4, rating = $5, reviews = $6, price = $7, updated_at = CURRENT_TIMESTAMP
		WHERE id = $8`

	_, err := r.db.Exec(query, destination.Name, destination.Description, destination.Location, destination.ImageURL, destination.Rating, destination.Reviews, destination.Price, destination.ID)
	if err != nil {
		return fmt.Errorf("failed to update destination: %w", err)
	}

	return nil
}

// DeleteDestination deletes a destination
func (r *destinationRepository) DeleteDestination(id int) error {
	query := `DELETE FROM destinations WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete destination: %w", err)
	}
	return nil
}

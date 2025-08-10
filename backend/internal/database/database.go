package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSL_MODE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// // CreateTables creates the necessary database tables
// func CreateTables() error {
// 	queries := []string{
// 		`CREATE TABLE IF NOT EXISTS users (
// 			id SERIAL PRIMARY KEY,
// 			email VARCHAR(255) UNIQUE NOT NULL,
// 			password VARCHAR(255) NOT NULL,
// 			first_name VARCHAR(100) NOT NULL,
// 			last_name VARCHAR(100) NOT NULL,
// 			phone VARCHAR(20),
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)`,
// 		`CREATE TABLE IF NOT EXISTS destinations (
// 			id SERIAL PRIMARY KEY,
// 			name VARCHAR(255) NOT NULL,
// 			country VARCHAR(100) NOT NULL,
// 			city VARCHAR(100) NOT NULL,
// 			description TEXT,
// 			image_url VARCHAR(500),
// 			rating DECIMAL(3,2) DEFAULT 0.0,
// 			deals_count INTEGER DEFAULT 0,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)`,
// 		`CREATE TABLE IF NOT EXISTS hotels (
// 			id SERIAL PRIMARY KEY,
// 			name VARCHAR(255) NOT NULL,
// 			destination_id INTEGER REFERENCES destinations(id),
// 			address TEXT NOT NULL,
// 			description TEXT,
// 			image_url VARCHAR(500),
// 			rating DECIMAL(3,2) DEFAULT 0.0,
// 			price_per_night DECIMAL(10,2) NOT NULL,
// 			amenities TEXT,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)`,
// 		`CREATE TABLE IF NOT EXISTS services (
// 			id SERIAL PRIMARY KEY,
// 			name VARCHAR(255) NOT NULL,
// 			category VARCHAR(100) NOT NULL,
// 			description TEXT,
// 			image_url VARCHAR(500),
// 			is_active BOOLEAN DEFAULT TRUE,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)`,
// 		`CREATE TABLE IF NOT EXISTS bookings (
// 			id SERIAL PRIMARY KEY,
// 			user_id INTEGER REFERENCES users(id),
// 			service_type VARCHAR(50) NOT NULL,
// 			service_id INTEGER NOT NULL,
// 			check_in_date DATE NOT NULL,
// 			check_out_date DATE NOT NULL,
// 			total_amount DECIMAL(10,2) NOT NULL,
// 			status VARCHAR(20) DEFAULT 'pending',
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)`,
// 	}

// 	for _, query := range queries {
// 		if _, err := DB.Exec(query); err != nil {
// 			return fmt.Errorf("failed to create table: %w", err)
// 		}
// 	}

// 	log.Println("Database tables created successfully")
// 	return nil
// }

// // SeedData inserts sample data into the database
// func SeedData() error {
// 	// Insert sample destinations
// 	destinations := []struct {
// 		name, country, city, description, imageURL string
// 		rating                                     float64
// 		dealsCount                                 int
// 	}{
// 		{"Dubai, UAE", "UAE", "Dubai", "A modern metropolis with luxury shopping and stunning architecture", "https://images.unsplash.com/photo-1582719478250-c89cae4dc85b", 4.8, 15},
// 		{"Cape Town, SA", "South Africa", "Cape Town", "Beautiful coastal city with Table Mountain and wine regions", "https://images.unsplash.com/photo-1580060839134-75a5edca2e99", 4.8, 12},
// 		{"Nairobi, Kenya", "Kenya", "Nairobi", "Gateway to East Africa with wildlife safaris and cultural experiences", "https://images.unsplash.com/photo-1516026672322-bc52d61a55d5", 4.8, 18},
// 		{"Lagos, Nigeria", "Nigeria", "Lagos", "Vibrant economic hub with rich culture and bustling markets", "https://images.unsplash.com/photo-1581833971358-2c8b550f87b3", 4.8, 9},
// 	}

// 	for _, dest := range destinations {
// 		_, err := DB.Exec(`
// 			INSERT INTO destinations (name, country, city, description, image_url, rating, deals_count)
// 			VALUES ($1, $2, $3, $4, $5, $6, $7)
// 			ON CONFLICT DO NOTHING`,
// 			dest.name, dest.country, dest.city, dest.description, dest.imageURL, dest.rating, dest.dealsCount)
// 		if err != nil {
// 			return fmt.Errorf("failed to insert destination: %w", err)
// 		}
// 	}

// 	// Insert sample services
// 	services := []struct {
// 		name, category, description, imageURL string
// 		isActive                              bool
// 	}{
// 		{"Hotels & Guesthouses", "accommodation", "Comfortable stays across Africa", "https://images.unsplash.com/photo-1566073771259-6a8506099945", true},
// 		{"Flights", "transportation", "International and regional flights", "https://images.unsplash.com/photo-1436491865332-7a61a109cc05", true},
// 		{"Bus Travel", "transportation", "Affordable intercity transportation", "https://images.unsplash.com/photo-1544620347-c4fd4a3d5957", true},
// 		{"Visa Assistance", "documentation", "Professional visa application help", "https://images.unsplash.com/photo-1578662996442-48f60103fc96", true},
// 		{"Car Rental & Rides", "transportation", "Self-drive and chauffeur options", "https://images.unsplash.com/photo-1449965408869-eaa3f722e40d", true},
// 		{"Nomado Love", "special", "Romantic getaways for couples", "https://images.unsplash.com/photo-1518621012382-8e77fbbaf6e9", true},
// 		{"Little Nomads", "special", "Educational school trips", "https://images.unsplash.com/photo-1503454537195-1dcabb73ffb9", true},
// 		{"Events & Retreats", "experience", "Transformative experiences", "https://images.unsplash.com/photo-1492684223066-81342ee5ff30", true},
// 		{"Nomado Jobs", "services", "International job opportunities", "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d", true},
// 		{"Nomado Shop", "retail", "Travel gear and African souvenirs", "https://images.unsplash.com/photo-1472851294608-062f824d29cc", true},
// 		{"Nomado Lux", "luxury", "Ultra-luxury travel experiences", "https://images.unsplash.com/photo-1582719478250-c89cae4dc85b", true},
// 		{"Nomado Forex", "financial", "Currency exchange and conversion", "https://images.unsplash.com/photo-1559526324-593bc073d938", true},
// 		{"Nomado Eats", "food", "Local dining and food experiences", "https://images.unsplash.com/photo-1414235077428-338989a2e8c0", true},
// 		{"Events & Social Travel", "social", "Join travel groups and local events", "https://images.unsplash.com/photo-1511632765486-a01980e01a18", true},
// 		{"eSIM & Translation", "technology", "Stay connected and communicate", "https://images.unsplash.com/photo-1512941937669-90a1b58e7e9c", true},
// 		{"Nomado Logistics", "logistics", "Freight and delivery services", "https://images.unsplash.com/photo-1566576912321-d58ddd7a6088", true},
// 	}

// 	for _, service := range services {
// 		_, err := DB.Exec(`
// 			INSERT INTO services (name, category, description, image_url, is_active)
// 			VALUES ($1, $2, $3, $4, $5)
// 			ON CONFLICT DO NOTHING`,
// 			service.name, service.category, service.description, service.imageURL, service.isActive)
// 		if err != nil {
// 			return fmt.Errorf("failed to insert service: %w", err)
// 		}
// 	}

// 	log.Println("Sample data inserted successfully")
// 	return nil
// }

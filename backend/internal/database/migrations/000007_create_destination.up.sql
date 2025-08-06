-- This migration creates a new table for destinations
-- The table will store destination information including name, description, location,
-- and timestamps for creation and updates.
CREATE TABLE IF NOT EXISTS destinations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    location VARCHAR(255) NOT NULL,
    image_url VARCHAR(255),
    rating FLOAT CHECK (rating >= 0 AND rating <= 5),
    reviews INTEGER DEFAULT 0,
    price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Insert default destinations
-- These are sample destinations for initial setup.
-- Insert default destinations
INSERT INTO destinations (name, description, location, image_url, rating, reviews, price)
VALUES
('Cape Town', 'A vibrant city known for its stunning landscapes and cultural diversity.', 'South Africa', '/home/diesel/Desktop/nomado-houses/backend/public/images/destinations/capetown.jpeg', 4.5, 120, 150.00),
('Dubai', 'A modern city known for luxury shopping, ultramodern architecture, and a lively nightlife.', 'United Arab Emirates', '/home/diesel/Desktop/nomado-houses/backend/public/images/destinations/dubai.jpeg', 4.8, 200, 250.00),
('Lagos', 'A bustling metropolis with a rich cultural scene and beautiful beaches.', 'Nigeria', '/home/diesel/Desktop/nomado-houses/backend/public/images/destinations/lagos.jpeg', 4.2, 150, 100.00),
('Nairobi', 'The capital city of Kenya, known for its national park and vibrant culture.', 'Kenya', '/home/diesel/Desktop/nomado-houses/backend/public/images/destinations/nairobi.jpeg', 4.6, 180, 120.00),
('Zanzibar', 'An archipelago off the coast of Tanzania known for its beautiful beaches and rich history.', 'Tanzania', '/home/diesel/Desktop/nomado-houses/backend/public/images/destinations/zanzibar.webp', 4.7, 90, 200.00);

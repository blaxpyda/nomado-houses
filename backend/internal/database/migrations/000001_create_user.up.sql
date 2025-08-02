-- Create the users table
-- This migration creates the users table with necessary fields
-- and constraints to store user information securely.
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15)
);

-- Insert a default admin user
-- This is a sample admin user for initial setup.
INSERT INTO users (first_name, last_name, email, password_hash, phone_number)
VALUES
('Admin', 'User', 'admin@example.com', '$2b$10$EIXZQ1z5Q5Z5Q5Z5Q5Z5QO', '1234567890');
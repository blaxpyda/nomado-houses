-- This creates a new table for providers
-- The table will store provider information including their name, email, and contact details.
CREATE TABLE IF NOT EXISTS providers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    company_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15)
);
-- Insert a default provider
-- This is a sample provider for initial setup.
INSERT INTO providers (first_name, last_name, company_name, email, password_hash, phone_number)
VALUES
('John', 'Doe', 'Example Company', 'john.doe@example.com', '$2b$10$EIXZQ1z5Q5Z5Q5Z5Q5Z5QO', '1234567890');
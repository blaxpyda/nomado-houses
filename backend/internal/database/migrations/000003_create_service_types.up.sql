-- This migration creates a new table for service types
-- The table will store service type information including their name and description.
CREATE TABLE IF NOT EXISTS service_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Insert default service types
-- These are sample service types for initial setup.
INSERT INTO service_types (name, description)
VALUES
('Cleaning', 'Professional cleaning services for homes and offices.'),
('Maintenance', 'General maintenance services for properties.'),
('Landscaping', 'Landscaping and gardening services for outdoor spaces.'),
('Security', 'Security services for residential and commercial properties.'),
('Transportation', 'Transportation services including logistics and moving.'),
('Catering', 'Catering services for events and gatherings.'),
('IT Support', 'IT support services for businesses and individuals.'),
('Consulting', 'Consulting services for business and personal needs.'),
('Legal', 'Legal services including advice and representation.'),
('Financial', 'Financial services including accounting and investment advice.');

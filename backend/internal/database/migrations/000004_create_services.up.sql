-- This migration creates a new table for services
-- The table will store service information including the provider, service type, name, description, price,
-- availability status, and timestamps for creation and updates.
CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    provider_id INTEGER NOT NULL REFERENCES providers(id) ON DELETE CASCADE,
    service_type_id INTEGER NOT NULL REFERENCES service_types(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    availability BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Insert default services
-- These are sample services for initial setup.
INSERT INTO services (provider_id, service_type_id, name, description, price, availability)
VALUES
(1, 1, 'Home Cleaning', 'Professional home cleaning service including dusting, vacuuming, and sanitizing.', 99.99, TRUE),
(1, 2, 'Property Maintenance', 'General maintenance services for residential properties including plumbing and electrical work.', 149.99, TRUE),
(1, 3, 'Lawn Care', 'Comprehensive lawn care services including mowing, trimming, and fertilization.', 79.99, TRUE),
(1, 4, 'Security Patrol', '24/7 security patrol services for residential and commercial properties.', 199.99, TRUE),
(1, 5, 'Logistics and Moving', 'Reliable logistics and moving services for residential and commercial needs.', 299.99, TRUE),
(1, 6, 'Event Catering', 'Catering services for events including weddings, parties, and corporate gatherings.', 499.99, TRUE),
(1, 7, 'IT Support', 'Comprehensive IT support services for businesses including network setup and troubleshooting.', 89.99, TRUE),
(1, 8, 'Business Consulting', 'Expert consulting services for business strategy and operations.', 199.99, TRUE),
(1, 9, 'Legal Advice', 'Professional legal advice and representation services.', 249.99, TRUE),
(1, 10, 'Financial Planning', 'Financial planning and investment advice services for individuals and businesses.', 349.99, TRUE);
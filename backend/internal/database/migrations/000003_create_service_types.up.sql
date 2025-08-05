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
('Hotels & Guesthouses', 'Comfortable stays across Africa'),
('Flights', 'International and regional flights'),
('Bus Travel', 'Affordable intercity transportation.'),
('Visa Assistance', 'Professional visa apllication help'),
('Car Rental & Rides', 'Self-drive and chauffeur options'),
('Nomado Love', 'Romantic getaways for couples'),
('Little Nomads', 'Educational school trips'),
('Events & Retreats', 'Transformative experiences'),
('Nomado Jobs', 'International job placements'),
('Nomado Shop', 'Travel gear and African souvenirs'),
('Nomado Lux', 'Ultra-luxury travel experiences'),
('Nomado Forex', 'Currency exchange and conversions'),
('Nomado Eats', 'Local dining and food experiences'),
('Events & Social Travel', 'Join travel groups and local events'),
('eSIM & Translation', 'Stay connected and communicate'),
('Nomado Logistics', 'Freight and delivery services');

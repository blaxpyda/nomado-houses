-- This migration creates a new table for bookings
-- The table will store booking information including user, service, provider, booking dates, total price,
-- status, and timestamps for creation and updates.
CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    service_id INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    provider_id INTEGER NOT NULL REFERENCES providers(id) ON DELETE CASCADE,
    booking_date_start TIMESTAMP NOT NULL,
    booking_date_end TIMESTAMP NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Insert default bookings
-- These are sample bookings for initial setup.
INSERT INTO bookings (user_id, service_id, provider_id, booking_date_start, booking_date_end, total_price, status)
VALUES
(1, 1, 1, '2023-10-01 10:00:00', '2023-10-01 12:00:00', 99.99, 'confirmed'),
(1, 2, 1, '2023-10-02 09:00:00', '2023-10-02 11:00:00', 149.99, 'pending'),
(1, 3, 1, '2023-10-03 08:00:00', '2023-10-03 10:00:00', 79.99, 'cancelled'),
(1, 4, 1, '2023-10-04 14:00:00', '2023-10-04 16:00:00', 199.99, 'confirmed');
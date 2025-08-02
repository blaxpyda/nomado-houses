-- This migration creates a new table for payments
-- The table will store payment information including user, booking, amount, payment date,
-- payment method, status, and timestamps for creation and updates.
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    booking_id INT NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    amount DECIMAL(10, 2) NOT NULL,
    payment_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    payment_method VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Insert default payments
-- These are sample payments for initial setup.
INSERT INTO payments (user_id, booking_id, amount, payment_date, payment_method, status)
VALUES
(1, 1, 99.99, '2023-10-01 10:00:00', 'credit_card', 'completed'),
(1, 2, 149.99, '2023-10-02 09:00:00', 'paypal', 'pending'),
(1, 3, 79.99, '2023-10-03 08:00:00', 'bank_transfer', 'cancelled'),
(1, 4, 199.99, '2023-10-04 14:00:00', 'credit_card', 'completed');

-- This migration creates a new table for destinations
-- The table will store destination information including name, description, location,
-- and timestamps for creation and updates.
CREATE TABLE IF NOT EXISTS destinations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Insert default destinations
-- These are sample destinations for initial setup.
INSERT INTO destinations (name, description, location)
VALUES
('Paris', 'The capital city of France, known for its art, fashion, and culture.',
 '48.8566° N, 2.3522° E'),
('New York', 'A bustling metropolis in the USA, famous for its skyline and cultural diversity.',
 '40.7128° N, 74.0060° W'),
('Tokyo', 'The capital city of Japan, known for its modernity and traditional culture.',
 '35.6762° N, 139.6503° E'),
('Sydney', 'A major city in Australia, known for its Sydney Opera House and Harbour Bridge.',
 '33.8688° S, 151.2093° E'),
('Cape Town', 'A coastal city in South Africa, known for its stunning landscapes and Table Mountain.',
 '33.9249° S, 18.4241° E'),
('Rio de Janeiro', 'A vibrant city in Brazil, famous for its beaches and Carnival festival.',
 '22.9068° S, 43.1729° W'),
('Dubai', 'A modern city in the UAE, known for luxury shopping and ultramodern architecture.',
 '25.276987° N, 55.296249° E'),
('Istanbul', 'A transcontinental city in Turkey, known for its rich history and cultural heritage.',
 '41.0082° N, 28.9784° E'),
('London', 'The capital city of the UK, known for its history, culture, and landmarks.',
 '51.5074° N, 0.1278° W');
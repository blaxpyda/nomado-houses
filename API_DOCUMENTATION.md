# Nomado Houses API Documentation

## Overview
This is a REST API for the Nomado Houses travel booking platform built with Go, following the repository pattern architecture.

## Base URL
```
http://localhost:8080/api
```

## Authentication
The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

## Swagger Documentation
Interactive API documentation is available at:
```
http://localhost:8080/swagger/index.html
```

## API Endpoints

### Authentication

#### Register
- **POST** `/auth/register`
- **Description**: Register a new user account
- **Body**:
```json
{
  "email": "user@example.com",
  "password": "password123",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890"
}
```
- **Response**: `201 Created`
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "token": "jwt-token-here",
    "user": {
      "id": 1,
      "email": "user@example.com",
      "first_name": "John",
      "last_name": "Doe",
      "phone": "+1234567890"
    }
  }
}
```

#### Login
- **POST** `/auth/login`
- **Description**: Login with existing credentials
- **Body**:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```
- **Response**: `200 OK`
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "jwt-token-here",
    "user": {
      "id": 1,
      "email": "user@example.com",
      "first_name": "John",
      "last_name": "Doe",
      "phone": "+1234567890"
    }
  }
}
```

### Destinations

#### Get All Destinations
- **GET** `/destinations`
- **Description**: Get all available destinations
- **Response**: `200 OK`
```json
{
  "success": true,
  "message": "Destinations retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Bali, Indonesia",
      "description": "Beautiful tropical paradise",
      "location": "Indonesia",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### Get Destination by ID
- **GET** `/destinations/{id}`
- **Description**: Get specific destination by ID
- **Response**: `200 OK`

#### Create Destination (Protected)
- **POST** `/destinations`
- **Description**: Create a new destination (Admin only)
- **Authentication**: Required
- **Body**:
```json
{
  "name": "Tokyo, Japan",
  "description": "Modern city with rich culture",
  "location": "Japan"
}
```

#### Update Destination (Protected)
- **PUT** `/destinations/{id}`
- **Description**: Update destination (Admin only)
- **Authentication**: Required

#### Delete Destination (Protected)
- **DELETE** `/destinations/{id}`
- **Description**: Delete destination (Admin only)
- **Authentication**: Required

### Services

#### Get All Services
- **GET** `/services`
- **Description**: Get all available services
- **Query Parameters**:
  - `category` (optional): Filter by service category
- **Response**: `200 OK`

#### Get Service by ID
- **GET** `/services/{id}`
- **Description**: Get specific service by ID
- **Response**: `200 OK`

#### Create Service (Protected)
- **POST** `/services`
- **Description**: Create a new service (Admin only)
- **Authentication**: Required
- **Body**:
```json
{
  "provider_id": 1,
  "service_type_id": 1,
  "name": "Luxury Villa Rental",
  "description": "Beautiful villa with ocean view",
  "price": 299.99,
  "availability": true
}
```

#### Update Service (Protected)
- **PUT** `/services/{id}`
- **Description**: Update service (Admin only)
- **Authentication**: Required

#### Delete Service (Protected)
- **DELETE** `/services/{id}`
- **Description**: Delete service (Admin only)
- **Authentication**: Required

### Service Types

#### Get All Service Types
- **GET** `/service-types`
- **Description**: Get all service categories
- **Response**: `200 OK`

#### Get Service Type by ID
- **GET** `/service-types/{id}`
- **Description**: Get specific service type by ID
- **Response**: `200 OK`

#### Create Service Type (Protected)
- **POST** `/service-types`
- **Description**: Create a new service type (Admin only)
- **Authentication**: Required
- **Body**:
```json
{
  "name": "Accommodation",
  "description": "Various types of accommodations"
}
```

#### Update Service Type (Protected)
- **PUT** `/service-types/{id}`
- **Description**: Update service type (Admin only)
- **Authentication**: Required

#### Delete Service Type (Protected)
- **DELETE** `/service-types/{id}`
- **Description**: Delete service type (Admin only)
- **Authentication**: Required

### Bookings (All Protected)

#### Create Booking
- **POST** `/bookings`
- **Description**: Create a new booking
- **Authentication**: Required
- **Body**:
```json
{
  "service_id": 1,
  "provider_id": 1,
  "booking_date_start": "2024-03-01T00:00:00Z",
  "booking_date_end": "2024-03-07T00:00:00Z",
  "total_price": 1299.99
}
```

#### Get User Bookings
- **GET** `/bookings`
- **Description**: Get all bookings for authenticated user
- **Authentication**: Required
- **Response**: `200 OK`

#### Get Booking by ID
- **GET** `/bookings/{id}`
- **Description**: Get specific booking by ID
- **Authentication**: Required
- **Response**: `200 OK`

#### Update Booking Status
- **PUT** `/bookings/{id}/status`
- **Description**: Update booking status
- **Authentication**: Required
- **Body**:
```json
{
  "status": "confirmed"
}
```
- **Valid statuses**: `pending`, `confirmed`, `cancelled`, `completed`

## Error Responses

All endpoints return consistent error responses:

```json
{
  "success": false,
  "message": "Error description",
  "error": "Detailed error message (optional)"
}
```

### Common HTTP Status Codes
- `200` - OK
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `404` - Not Found
- `500` - Internal Server Error

## Data Models

### User
```json
{
  "id": 1,
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890"
}
```

### Destination
```json
{
  "id": 1,
  "name": "Destination Name",
  "description": "Destination description",
  "location": "Location",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Service
```json
{
  "id": 1,
  "provider_id": 1,
  "service_type_id": 1,
  "name": "Service Name",
  "description": "Service description",
  "price": 99.99,
  "availability": true,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### ServiceType
```json
{
  "id": 1,
  "name": "Service Type Name",
  "description": "Service type description"
}
```

### Booking
```json
{
  "id": 1,
  "user_id": 1,
  "service_id": 1,
  "provider_id": 1,
  "booking_date_start": "2024-03-01T00:00:00Z",
  "booking_date_end": "2024-03-07T00:00:00Z",
  "total_price": 299.99,
  "status": "pending",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## Getting Started

1. **Start the server**:
   ```bash
   cd backend
   go run main.go
   ```

2. **Register a new user**:
   ```bash
   curl -X POST http://localhost:8080/api/auth/register \
     -H "Content-Type: application/json" \
     -d '{
       "email": "test@example.com",
       "password": "password123",
       "first_name": "Test",
       "last_name": "User",
       "phone": "+1234567890"
     }'
   ```

3. **Login to get JWT token**:
   ```bash
   curl -X POST http://localhost:8080/api/auth/login \
     -H "Content-Type: application/json" \
     -d '{
       "email": "test@example.com",
       "password": "password123"
     }'
   ```

4. **Use the token for authenticated requests**:
   ```bash
   curl -X POST http://localhost:8080/api/bookings \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_JWT_TOKEN" \
     -d '{
       "service_id": 1,
       "provider_id": 1,
       "booking_date_start": "2024-03-01T00:00:00Z",
       "booking_date_end": "2024-03-07T00:00:00Z",
       "total_price": 299.99
     }'
   ```

## Architecture

The API follows the repository pattern with clean architecture:

- **handlers/**: HTTP request handlers
- **service/**: Business logic layer
- **repository/**: Data access layer
- **models/**: Data structures and DTOs
- **database/**: Database connection and setup
- **logger/**: Logging utility

## Environment Variables

Create a `.env` file in the backend directory with:
```
PORT=8080
JWT_SECRET=your-secret-key-here
DB_HOST=localhost
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=nomado_houses
```

## Testing with Postman

Import the API into Postman using the Swagger documentation URL or create a collection with the endpoints listed above.

## Contributing

1. Follow the existing code structure
2. Add proper Swagger annotations for new endpoints
3. Maintain the repository pattern
4. Add proper error handling
5. Update this documentation for any changes
# Nomado Houses - Travel Platform

A comprehensive travel platform built with Go backend following the repository pattern and vanilla JavaScript frontend. Inspired by Nomado Global Adventures, this application provides travel services including destinations, hotels, bookings, and user management.

## Features

### Backend (Go)
- **Clean Architecture**: Repository pattern with service layer
- **RESTful API**: Well-structured endpoints for all resources
- **JWT Authentication**: Secure user authentication and authorization
- **PostgreSQL Database**: Robust data persistence
- **CORS Support**: Cross-origin resource sharing enabled
- **Environment Configuration**: Flexible configuration management

### Frontend (Vanilla JavaScript)
- **Modern UI**: Responsive design with smooth animations
- **Authentication**: User registration and login
- **Destinations**: Browse popular travel destinations
- **Services**: Explore various travel services
- **Booking System**: Quick booking forms for hotels, buses, and visa assistance
- **Notifications**: Real-time user feedback system

## Project Structure

```
nomado-houses/
├── backend/
│   ├── internal/
│   │   ├── database/       # Database connection and setup
│   │   ├── models/         # Data models and structures
│   │   ├── repository/     # Data access layer
│   │   ├── service/        # Business logic layer
│   │   ├── handlers/       # HTTP handlers
│   │   └── utils/          # Utility functions
│   ├── main.go            # Application entry point
│   ├── go.mod             # Go module dependencies
│   └── .env               # Environment variables
├── web/
│   ├── index.html         # Main HTML file
│   ├── styles.css         # Styling
│   └── script.js          # Frontend logic
├── docker-compose.yml     # PostgreSQL container
├── Makefile              # Build and run commands
└── README.md             # This file
```

## API Endpoints

### Authentication
- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login

### Destinations
- `GET /api/destinations` - Get all destinations
- `GET /api/destinations/{id}` - Get destination by ID

### Services
- `GET /api/services` - Get all services
- `GET /api/services?category={category}` - Get services by category
- `GET /api/services/{id}` - Get service by ID

### Bookings (Protected)
- `POST /api/bookings` - Create new booking
- `GET /api/bookings` - Get user's bookings
- `GET /api/bookings/{id}` - Get booking by ID

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL (or use Docker)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd nomado-houses
   ```

2. **Start the database**
   ```bash
   make db-up
   ```

3. **Install Go dependencies**
   ```bash
   make install
   ```

4. **Start the development server**
   ```bash
   make dev
   ```

5. **Serve the frontend** (in a new terminal)
   ```bash
   make frontend
   ```

### Alternative Setup

**Manual setup:**

1. **Start PostgreSQL**
   ```bash
   docker-compose up -d postgres
   ```

2. **Run the backend**
   ```bash
   cd backend
   go run main.go
   ```

3. **Serve the frontend**
   ```bash
   cd web
   python3 -m http.server 3000
   ```

## Configuration

### Environment Variables (.env)
```env
DB_HOST=localhost
DB_PORT=5432
DB_NAME=nomado_db
DB_USER=postgres
DB_PASSWORD=password
DB_SSL_MODE=disable
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
PORT=8080
```

## Database Schema

### Users Table
- `id` - Primary key
- `email` - User email (unique)
- `password` - Hashed password
- `first_name` - User's first name
- `last_name` - User's last name
- `phone` - Contact number
- `created_at` - Account creation timestamp
- `updated_at` - Last update timestamp

### Destinations Table
- `id` - Primary key
- `name` - Destination name
- `country` - Country
- `city` - City
- `description` - Description
- `image_url` - Image URL
- `rating` - Average rating
- `deals_count` - Number of available deals

### Services Table
- `id` - Primary key
- `name` - Service name
- `category` - Service category
- `description` - Service description
- `image_url` - Service image
- `is_active` - Active status

### Bookings Table
- `id` - Primary key
- `user_id` - Foreign key to users
- `service_type` - Type of service
- `service_id` - ID of the service
- `check_in_date` - Check-in date
- `check_out_date` - Check-out date
- `total_amount` - Total booking amount
- `status` - Booking status

## Available Make Commands

```bash
make help         # Show help message
make build        # Build the Go application
make run          # Run the application
make test         # Run tests
make clean        # Clean build artifacts
make db-up        # Start PostgreSQL database
make db-down      # Stop PostgreSQL database
make dev          # Start development environment
make install      # Install dependencies
make frontend     # Serve frontend files
```

## Usage

1. **Access the application**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080

2. **Register a new account**
   - Click "Sign Up" in the navigation
   - Fill in your details
   - You'll be automatically logged in

3. **Explore destinations**
   - Browse popular destinations with ratings and deals
   - View detailed information for each destination

4. **Browse services**
   - Explore various travel services
   - Services are categorized for easy navigation

5. **Make bookings**
   - Use the quick booking forms
   - Login required for bookings
   - View your bookings in the user dashboard

## Technologies Used

### Backend
- **Go**: Main programming language
- **Gorilla Mux**: HTTP router and URL matcher
- **PostgreSQL**: Database
- **JWT**: Authentication tokens
- **bcrypt**: Password hashing

### Frontend
- **HTML5**: Markup
- **CSS3**: Styling with modern features
- **Vanilla JavaScript**: No frameworks, pure JS
- **Font Awesome**: Icons
- **Google Fonts**: Typography

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Support

For support, please open an issue in the GitHub repository or contact the development team.

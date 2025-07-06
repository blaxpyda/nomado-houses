package main

import (
	"log"
	"net/http"
	"os"

	"nomado-houses/internal/database"
	appHandlers "nomado-houses/internal/handlers"
	"nomado-houses/internal/repository"
	"nomado-houses/internal/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CloseDB()

	// Create tables and seed data
	if err := database.CreateTables(); err != nil {
		log.Fatal("Failed to create tables:", err)
	}

	if err := database.SeedData(); err != nil {
		log.Fatal("Failed to seed data:", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(database.DB)
	destinationRepo := repository.NewDestinationRepository(database.DB)
	serviceRepo := repository.NewServiceRepository(database.DB)
	bookingRepo := repository.NewBookingRepository(database.DB)

	// Initialize services
	authService := service.NewAuthService(userRepo)
	destinationService := service.NewDestinationService(destinationRepo)
	serviceService := service.NewServiceService(serviceRepo)
	bookingService := service.NewBookingService(bookingRepo)

	// Initialize handlers
	authHandler := appHandlers.NewAuthHandler(authService)
	destinationHandler := appHandlers.NewDestinationHandler(destinationService)
	serviceHandler := appHandlers.NewServiceHandler(serviceService)
	bookingHandler := appHandlers.NewBookingHandler(bookingService)

	// Setup routes
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Public routes
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	api.HandleFunc("/destinations", destinationHandler.GetAllDestinations).Methods("GET")
	api.HandleFunc("/destinations/{id}", destinationHandler.GetDestinationByID).Methods("GET")
	api.HandleFunc("/services", serviceHandler.GetAllServices).Methods("GET")
	api.HandleFunc("/services/{id}", serviceHandler.GetServiceByID).Methods("GET")

	// Protected routes
	protected := api.PathPrefix("").Subrouter()
	protected.Use(authHandler.AuthMiddleware)
	protected.HandleFunc("/bookings", bookingHandler.CreateBooking).Methods("POST")
	protected.HandleFunc("/bookings", bookingHandler.GetUserBookings).Methods("GET")
	protected.HandleFunc("/bookings/{id}", bookingHandler.GetBookingByID).Methods("GET")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	// Setup CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))

	log.Println("Server stopped")
}

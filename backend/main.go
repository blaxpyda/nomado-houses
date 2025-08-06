package main

import (
	"log"
	"net/http"
	"os"

	_ "nomado-houses/docs"
	"nomado-houses/internal/database"
	appHandlers "nomado-houses/internal/handlers"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/repository"
	"nomado-houses/internal/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("logs/nomado.log")
	if err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	return logInstance
}

// @title Nomado Houses API
// @version 1.0
// @description API for Nomado Houses travel booking platform
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @schemes http

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize logger
	logInstance := initializeLogger()
	defer logInstance.Close()

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CloseDB()

	// Initialize repositories
	userRepo := repository.NewUserRepository(database.DB, logInstance)
	destinationRepo := repository.NewDestinationRepository(database.DB, logInstance)
	serviceRepo := repository.NewServiceRepository(database.DB, logInstance)
	serviceTypeRepo := repository.NewServiceTypeRepository(database.DB, logInstance)
	bookingRepo := repository.NewBookingRepository(database.DB, logInstance)

	// Initialize services
	authService := service.NewAuthService(userRepo)
	destinationService := service.NewDestinationService(destinationRepo)
	serviceService := service.NewServiceService(serviceRepo)
	serviceTypeService := service.NewServiceTypeService(serviceTypeRepo)
	bookingService := service.NewBookingService(bookingRepo)

	// Initialize handlers
	authHandler := appHandlers.NewAuthHandler(authService, logInstance)
	destinationHandler := appHandlers.NewDestinationHandler(destinationService, logInstance)
	serviceHandler := appHandlers.NewServiceHandler(serviceService, logInstance)
	serviceTypeHandler := appHandlers.NewServiceTypeHandler(serviceTypeService, logInstance)
	bookingHandler := appHandlers.NewBookingHandler(bookingService, logInstance)

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
	api.HandleFunc("/service-types", serviceTypeHandler.GetAllServiceTypes).Methods("GET")
	api.HandleFunc("/service-types/{id}", serviceTypeHandler.GetServiceTypeByID).Methods("GET")

	// Protected routes
	protected := api.PathPrefix("").Subrouter()
	protected.Use(authHandler.AuthMiddleware)

	// User routes
	protected.HandleFunc("/bookings", bookingHandler.CreateBooking).Methods("POST")
	protected.HandleFunc("/bookings", bookingHandler.GetUserBookings).Methods("GET")
	protected.HandleFunc("/bookings/{id}", bookingHandler.GetBookingByID).Methods("GET")
	protected.HandleFunc("/bookings/{id}/status", bookingHandler.UpdateBookingStatus).Methods("PUT")

	// Admin routes (TODO: Add admin role middleware)
	protected.HandleFunc("/destinations", destinationHandler.CreateDestination).Methods("POST")
	protected.HandleFunc("/destinations/{id}", destinationHandler.UpdateDestination).Methods("PUT")
	protected.HandleFunc("/destinations/{id}", destinationHandler.DeleteDestination).Methods("DELETE")

	protected.HandleFunc("/services", serviceHandler.CreateService).Methods("POST")
	protected.HandleFunc("/services/{id}", serviceHandler.UpdateService).Methods("PUT")
	protected.HandleFunc("/services/{id}", serviceHandler.DeleteService).Methods("DELETE")

	protected.HandleFunc("/service-types", serviceTypeHandler.CreateServiceType).Methods("POST")
	protected.HandleFunc("/service-types/{id}", serviceTypeHandler.UpdateServiceType).Methods("PUT")
	protected.HandleFunc("/service-types/{id}", serviceTypeHandler.DeleteServiceType).Methods("DELETE")

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	// Setup CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000", "http://127.0.0.1:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
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

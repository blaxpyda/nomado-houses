package main

import (
	"log"
	"net/http"
	"os"

	_ "nomado-houses/docs"
	"nomado-houses/internal/database"
	appHandlers "nomado-houses/internal/handlers"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/middleware"
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
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	destinationService := service.NewDestinationService(destinationRepo)
	serviceService := service.NewServiceService(serviceRepo)
	serviceTypeService := service.NewServiceTypeService(serviceTypeRepo)
	bookingService := service.NewBookingService(bookingRepo)
	travelPayoutsService := service.NewTravelPayoutsService()

	// Initialize middleware
	roleMiddleware := middleware.NewRoleMiddleware(authService, userService)

	// Initialize handlers
	authHandler := appHandlers.NewAuthHandler(authService, logInstance)
	userHandler := appHandlers.NewUserHandler(userService, logInstance)
	destinationHandler := appHandlers.NewDestinationHandler(destinationService, logInstance)
	serviceHandler := appHandlers.NewServiceHandler(serviceService, logInstance)
	serviceTypeHandler := appHandlers.NewServiceTypeHandler(serviceTypeService, logInstance)
	bookingHandler := appHandlers.NewBookingHandler(bookingService, logInstance)
	// hotelHandler := appHandlers.NewHotelHandler(travelPayoutsService, logInstance)
	flightHandler := appHandlers.NewFlightHandler(travelPayoutsService, logInstance)

	// Setup routes
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Public routes
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	api.HandleFunc("/auth/verify-email", authHandler.VerifyEmail).Methods("POST")
	api.HandleFunc("/auth/resend-verification", authHandler.ResendVerification).Methods("POST")
	api.HandleFunc("/destinations", destinationHandler.GetAllDestinations).Methods("GET")
	api.HandleFunc("/destinations/{id}", destinationHandler.GetDestinationByID).Methods("GET")
	api.HandleFunc("/services", serviceHandler.GetAllServices).Methods("GET")
	api.HandleFunc("/services/{id}", serviceHandler.GetServiceByID).Methods("GET")
	api.HandleFunc("/service-types", serviceTypeHandler.GetAllServiceTypes).Methods("GET")
	api.HandleFunc("/service-types/{id}", serviceTypeHandler.GetServiceTypeByID).Methods("GET")

	// // Hotel routes
	// api.HandleFunc("/hotels/search", hotelHandler.SearchHotels).Methods("GET")
	// api.HandleFunc("/hotels/popular/{cityId}", hotelHandler.GetPopularHotels).Methods("GET")
	// api.HandleFunc("/hotels/destinations", hotelHandler.GetDestinations).Methods("GET")

	// Flight routes
	api.HandleFunc("/flights/search", flightHandler.SearchFlights).Methods("GET")
	api.HandleFunc("/flights/popular-routes", flightHandler.GetPopularRoutes).Methods("GET")
	api.HandleFunc("/flights/airports", flightHandler.GetAirports).Methods("GET")

	// Protected routes (any authenticated user)
	protected := api.PathPrefix("").Subrouter()
	protected.Use(authHandler.AuthMiddleware)

	// User profile routes (any authenticated user)
	protected.HandleFunc("/user/profile", userHandler.GetProfile).Methods("GET")
	protected.HandleFunc("/user/profile", userHandler.UpdateProfile).Methods("PUT")

	// User booking routes (any authenticated user)
	protected.HandleFunc("/bookings", bookingHandler.CreateBooking).Methods("POST")
	protected.HandleFunc("/bookings", bookingHandler.GetUserBookings).Methods("GET")
	protected.HandleFunc("/bookings/{id}", bookingHandler.GetBookingByID).Methods("GET")

	// Provider routes (provider or admin only)
	providerRoutes := api.PathPrefix("/provider").Subrouter()
	providerRoutes.Use(roleMiddleware.RequireAdminOrProvider())

	// Services management (providers can create/manage their services)
	providerRoutes.HandleFunc("/services", serviceHandler.CreateService).Methods("POST")
	providerRoutes.HandleFunc("/services/{id}", serviceHandler.UpdateService).Methods("PUT")
	providerRoutes.HandleFunc("/services/{id}", serviceHandler.DeleteService).Methods("DELETE")

	// Admin routes (admin only)
	adminRoutes := api.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(roleMiddleware.RequireAdmin())

	// User management
	adminRoutes.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	adminRoutes.HandleFunc("/users/role/{role}", userHandler.GetUsersByRole).Methods("GET")
	adminRoutes.HandleFunc("/users/{id}/role", userHandler.UpdateUserRole).Methods("PUT")

	// Destinations management (admin only)
	adminRoutes.HandleFunc("/destinations", destinationHandler.CreateDestination).Methods("POST")
	adminRoutes.HandleFunc("/destinations/{id}", destinationHandler.UpdateDestination).Methods("PUT")
	adminRoutes.HandleFunc("/destinations/{id}", destinationHandler.DeleteDestination).Methods("DELETE")

	// Service types management (admin only)
	adminRoutes.HandleFunc("/service-types", serviceTypeHandler.CreateServiceType).Methods("POST")
	adminRoutes.HandleFunc("/service-types/{id}", serviceTypeHandler.UpdateServiceType).Methods("PUT")
	adminRoutes.HandleFunc("/service-types/{id}", serviceTypeHandler.DeleteServiceType).Methods("DELETE")

	// Booking status management (admin only)
	adminRoutes.HandleFunc("/bookings/{id}/status", bookingHandler.UpdateBookingStatus).Methods("PUT")

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Auth page routes
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/login.html")
	})
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/register.html")
	})
	r.HandleFunc("/verify-email", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/verify-email.html")
	})
	// r.HandleFunc("/hotels", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./public/hotels.html")
	// })
	r.HandleFunc("/flights", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/flights.html")
	})

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

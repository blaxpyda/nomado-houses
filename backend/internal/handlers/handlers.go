package handlers

import (
	"encoding/json"
	"net/http"
	"nomado-houses/internal/models"
	"nomado-houses/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

// DestinationHandler handles destination requests
type DestinationHandler struct {
	destinationService service.DestinationService
}

// NewDestinationHandler creates a new destination handler
func NewDestinationHandler(destinationService service.DestinationService) *DestinationHandler {
	return &DestinationHandler{destinationService: destinationService}
}

// GetAllDestinations handles GET /api/destinations
func (h *DestinationHandler) GetAllDestinations(w http.ResponseWriter, r *http.Request) {
	destinations, err := h.destinationService.GetAllDestinations()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Destinations retrieved successfully",
		Data:    destinations,
	})
}

// GetDestinationByID handles GET /api/destinations/{id}
func (h *DestinationHandler) GetDestinationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid destination ID")
		return
	}

	destination, err := h.destinationService.GetDestinationByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Destination retrieved successfully",
		Data:    destination,
	})
}

// ServiceHandler handles service requests
type ServiceHandler struct {
	serviceService service.ServiceService
}

// NewServiceHandler creates a new service handler
func NewServiceHandler(serviceService service.ServiceService) *ServiceHandler {
	return &ServiceHandler{serviceService: serviceService}
}

// GetAllServices handles GET /api/services
func (h *ServiceHandler) GetAllServices(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	var services []models.Service
	var err error

	if category != "" {
		services, err = h.serviceService.GetServicesByCategory(category)
	} else {
		services, err = h.serviceService.GetAllServices()
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Services retrieved successfully",
		Data:    services,
	})
}

// GetServiceByID handles GET /api/services/{id}
func (h *ServiceHandler) GetServiceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid service ID")
		return
	}

	service, err := h.serviceService.GetServiceByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Service retrieved successfully",
		Data:    service,
	})
}

// BookingHandler handles booking requests
type BookingHandler struct {
	bookingService service.BookingService
}

// NewBookingHandler creates a new booking handler
func NewBookingHandler(bookingService service.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

// CreateBooking handles POST /api/bookings
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Get user ID from header (set by auth middleware)
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	booking.UserID = userID
	booking.Status = "pending"

	if err := h.bookingService.CreateBooking(&booking); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Booking created successfully",
		Data:    booking,
	})
}

// GetUserBookings handles GET /api/bookings
func (h *BookingHandler) GetUserBookings(w http.ResponseWriter, r *http.Request) {
	// Get user ID from header (set by auth middleware)
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	bookings, err := h.bookingService.GetBookingsByUserID(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Bookings retrieved successfully",
		Data:    bookings,
	})
}

// GetBookingByID handles GET /api/bookings/{id}
func (h *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid booking ID")
		return
	}

	booking, err := h.bookingService.GetBookingByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Booking retrieved successfully",
		Data:    booking,
	})
}

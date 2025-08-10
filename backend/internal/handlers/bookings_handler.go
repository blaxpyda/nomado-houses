package handlers

import (
	"encoding/json"
	"net/http"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/models"
	"nomado-houses/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

// BookingHandler handles booking requests
type BookingHandler struct {
	bookingService service.BookingService
	logger         *logger.Logger
}

// NewBookingHandler creates a new booking handler
func NewBookingHandler(bookingService service.BookingService, logger *logger.Logger) *BookingHandler {
	return &BookingHandler{bookingService: bookingService, logger: logger}
}

// CreateBooking handles POST /api/bookings
// @Summary Create booking
// @Description Create a new booking for a service
// @Tags Bookings
// @Accept json
// @Produce json
// @Param request body models.CreateBookingRequest true "Create booking request"
// @Security Bearer
// @Success 201 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /bookings [post]
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	booking := &models.Booking{
		UserID:           userID,
		ServiceID:        req.ServiceID,
		ProviderID:       req.ProviderID,
		BookingDateStart: req.BookingDateStart,
		BookingDateEnd:   req.BookingDateEnd,
		TotalPrice:       req.TotalPrice,
		Status:           "pending",
	}

	if err := h.bookingService.CreateBooking(booking); err != nil {
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
// @Summary Get user bookings
// @Description Get all bookings for the authenticated user
// @Tags Bookings
// @Produce json
// @Security Bearer
// @Success 200 {object} models.APIResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /bookings [get]
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

// UpdateBookingStatus handles PUT /api/bookings/{id}/status
func (h *BookingHandler) UpdateBookingStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid booking ID")
		return
	}

	var req models.UpdateBookingStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.bookingService.UpdateBookingStatus(id, req.Status); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Booking status updated successfully",
	})
}


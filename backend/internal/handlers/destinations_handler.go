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

// DestinationsHandler handles destination-related requests
type DestinationsHandler struct {
	destinationService service.DestinationService
	logger             *logger.Logger
}

// NewDestinationsHandler creates a new destinations handler
func NewDestinationsHandler(destinationService service.DestinationService, logger *logger.Logger) *DestinationsHandler {
	return &DestinationsHandler{
		destinationService: destinationService,
		logger:             logger,
	}
}

// respondWithError sends an error response
func (h *DestinationsHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJSON(w, code, models.ErrorResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// respondWithJSON sends a JSON response
func (h *DestinationsHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		h.logger.Error("Failed to encode JSON response", err)
	}
}

// GetAllDestinations handles GET /api/destinations
// @Summary Get all destinations
// @Description Get all available destinations from the database
// @Tags Destinations
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /destinations [get]
func (h *DestinationsHandler) GetAllDestinations(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Fetching all destinations from database")

	destinations, err := h.destinationService.GetAllDestinations()
	if err != nil {
		h.logger.Error("Failed to get destinations", err)
		h.respondWithError(w, http.StatusInternalServerError, "Failed to retrieve destinations")
		return
	}

	h.logger.Info("Successfully retrieved destinations from database")
	h.respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Destinations retrieved successfully",
		Data:    destinations,
	})
}

// GetDestinationByID handles GET /api/destinations/{id}
// @Summary Get destination by ID
// @Description Get a specific destination from the database by its ID
// @Tags Destinations
// @Produce json
// @Param id path int true "Destination ID"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /destinations/{id} [get]
func (h *DestinationsHandler) GetDestinationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid destination ID", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid destination ID")
		return
	}

	h.logger.Info("Fetching destination with ID: " + vars["id"])

	destination, err := h.destinationService.GetDestinationByID(id)
	if err != nil {
		h.logger.Error("Failed to get destination", err)
		h.respondWithError(w, http.StatusNotFound, "Destination not found")
		return
	}

	h.logger.Info("Successfully retrieved destination from database")
	h.respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Destination retrieved successfully",
		Data:    destination,
	})
}

// CreateDestination handles POST /api/destinations
// @Summary Create destination
// @Description Create a new destination in the database (Admin only)
// @Tags Destinations
// @Accept json
// @Produce json
// @Param request body models.CreateDestinationRequest true "Create destination request"
// @Security Bearer
// @Success 201 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /destinations [post]
func (h *DestinationsHandler) CreateDestination(w http.ResponseWriter, r *http.Request) {
	var req models.CreateDestinationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request payload", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	h.logger.Info("Creating new destination: " + req.Name)

	destination := &models.Destination{
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
	}

	if err := h.destinationService.CreateDestination(destination); err != nil {
		h.logger.Error("Failed to create destination", err)
		h.respondWithError(w, http.StatusInternalServerError, "Failed to create destination")
		return
	}

	h.logger.Info("Successfully created destination in database")
	h.respondWithJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Destination created successfully",
		Data:    destination,
	})
}

// UpdateDestination handles PUT /api/destinations/{id}
// @Summary Update destination
// @Description Update an existing destination in the database (Admin only)
// @Tags Destinations
// @Accept json
// @Produce json
// @Param id path int true "Destination ID"
// @Param request body models.UpdateDestinationRequest true "Update destination request"
// @Security Bearer
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /destinations/{id} [put]
func (h *DestinationsHandler) UpdateDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid destination ID", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid destination ID")
		return
	}

	var req models.UpdateDestinationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request payload", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	h.logger.Info("Updating destination with ID: " + vars["id"])

	destination := &models.Destination{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
	}

	if err := h.destinationService.UpdateDestination(destination); err != nil {
		h.logger.Error("Failed to update destination", err)
		h.respondWithError(w, http.StatusInternalServerError, "Failed to update destination")
		return
	}

	h.logger.Info("Successfully updated destination in database")
	h.respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Destination updated successfully",
		Data:    destination,
	})
}

// DeleteDestination handles DELETE /api/destinations/{id}
// @Summary Delete destination
// @Description Delete a destination from the database (Admin only)
// @Tags Destinations
// @Produce json
// @Param id path int true "Destination ID"
// @Security Bearer
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /destinations/{id} [delete]
func (h *DestinationsHandler) DeleteDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid destination ID", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid destination ID")
		return
	}

	h.logger.Info("Deleting destination with ID: " + vars["id"])

	if err := h.destinationService.DeleteDestination(id); err != nil {
		h.logger.Error("Failed to delete destination", err)
		h.respondWithError(w, http.StatusInternalServerError, "Failed to delete destination")
		return
	}

	h.logger.Info("Successfully deleted destination from database")
	h.respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Destination deleted successfully",
	})
}

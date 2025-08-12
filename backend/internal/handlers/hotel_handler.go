package handlers

import (
	"net/http"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/models"
	"nomado-houses/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

// HotelHandler handles hotel-related requests
type HotelHandler struct {
	travelService service.TravelPayoutsService
	logger        *logger.Logger
}

// NewHotelHandler creates a new hotel handler
func NewHotelHandler(travelService service.TravelPayoutsService, logger *logger.Logger) *HotelHandler {
	return &HotelHandler{
		travelService: travelService,
		logger:        logger,
	}
}

// SearchHotels handles hotel search requests
// @Summary Search hotels
// @Description Search for hotels in a destination
// @Tags Hotels
// @Accept json
// @Produce json
// @Param destination query string true "Destination"
// @Param checkIn query string true "Check-in date (YYYY-MM-DD)"
// @Param checkOut query string true "Check-out date (YYYY-MM-DD)"
// @Param adults query int false "Number of adults" default(2)
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /hotels/search [get]
func (h *HotelHandler) SearchHotels(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	destination := r.URL.Query().Get("destination")
	checkIn := r.URL.Query().Get("checkIn")
	checkOut := r.URL.Query().Get("checkOut")
	adultsStr := r.URL.Query().Get("adults")

	// Validate required parameters
	if destination == "" {
		respondWithError(w, http.StatusBadRequest, "Destination is required")
		return
	}

	if checkIn == "" {
		respondWithError(w, http.StatusBadRequest, "Check-in date is required")
		return
	}

	if checkOut == "" {
		respondWithError(w, http.StatusBadRequest, "Check-out date is required")
		return
	}

	// Parse adults parameter
	adults := 2 // default value
	if adultsStr != "" {
		if parsed, err := strconv.Atoi(adultsStr); err == nil && parsed > 0 {
			adults = parsed
		}
	}

	// Search hotels
	result, err := h.travelService.SearchHotels(destination, checkIn, checkOut, adults)
	if err != nil {
		h.logger.Error("Failed to search hotels", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to search hotels")
		return
	}

	// Return results
	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Hotels retrieved successfully",
		Data:    result,
	})
}

// GetPopularHotels handles requests for popular hotels
// @Summary Get popular hotels
// @Description Get popular hotels in a city
// @Tags Hotels
// @Produce json
// @Param cityId path string true "City ID"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /hotels/popular/{cityId} [get]
func (h *HotelHandler) GetPopularHotels(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityID := vars["cityId"]

	if cityID == "" {
		respondWithError(w, http.StatusBadRequest, "City ID is required")
		return
	}

	result, err := h.travelService.GetPopularHotels(cityID)
	if err != nil {
		h.logger.Error("Failed to get popular hotels", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to get popular hotels")
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Popular hotels retrieved successfully",
		Data:    result,
	})
}

// GetDestinations handles destination search requests
// @Summary Search destinations
// @Description Search for destinations
// @Tags Hotels
// @Produce json
// @Param query query string true "Search query"
// @Success 200 {object} models.APIResponse
// @Router /hotels/destinations [get]
func (h *HotelHandler) GetDestinations(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	if query == "" {
		respondWithError(w, http.StatusBadRequest, "Search query is required")
		return
	}

	result, err := h.travelService.GetDestinations(query)
	if err != nil {
		h.logger.Error("Failed to search destinations", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to search destinations")
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Destinations retrieved successfully",
		Data:    result,
	})
}

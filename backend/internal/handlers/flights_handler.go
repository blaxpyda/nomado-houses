package handlers

import (
	"encoding/json"
	"net/http"
	"nomado-houses/internal/logger"
	"nomado-houses/internal/models"
	"nomado-houses/internal/service"
)

// FlightHandler handles flight-related HTTP requests
type FlightHandler struct {
	travelPayoutsService service.TravelPayoutsService
	logger               *logger.Logger
}

// NewFlightHandler creates a new flight handler
func NewFlightHandler(travelPayoutsService service.TravelPayoutsService, logger *logger.Logger) *FlightHandler {
	return &FlightHandler{
		travelPayoutsService: travelPayoutsService,
		logger:               logger,
	}
}

// FlightSearchRequest represents the flight search request
type FlightSearchRequest struct {
	Origin      string `json:"origin" validate:"required"`
	Destination string `json:"destination" validate:"required"`
	DepartDate  string `json:"depart_date" validate:"required"`
	ReturnDate  string `json:"return_date,omitempty"` // Optional for one-way flights
}

// SearchFlights searches for flights
// @Summary Search flights
// @Description Search for flights between two destinations
// @Tags flights
// @Accept json
// @Produce json
// @Param origin query string true "Origin airport/city code"
// @Param destination query string true "Destination airport/city code"
// @Param depart_date query string true "Departure date (YYYY-MM-DD)"
// @Param return_date query string false "Return date (YYYY-MM-DD) - optional for one-way"
// @Success 200 {object} service.FlightSearchResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/flights/search [get]
func (h *FlightHandler) SearchFlights(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	origin := r.URL.Query().Get("origin")
	destination := r.URL.Query().Get("destination")
	departDate := r.URL.Query().Get("depart_date")
	returnDate := r.URL.Query().Get("return_date")

	// Validate required parameters
	if origin == "" || destination == "" || departDate == "" {
		h.errorResponse(w, "Missing required parameters: origin, destination, and depart_date", http.StatusBadRequest)
		return
	}

	// Search for flights using the TravelPayouts service
	flights, err := h.travelPayoutsService.SearchFlights(origin, destination, departDate, returnDate)
	if err != nil {
		h.logger.Error("Failed to search flights: %v", err)
		h.errorResponse(w, "Failed to search flights: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return successful response with proper API format
	response := models.APIResponse{
		Success: true,
		Message: "Flights retrieved successfully",
		Data:    flights,
	}
	h.successResponse(w, response)
}

// GetPopularRoutes returns popular flight routes
// @Summary Get popular flight routes
// @Description Get popular flight routes and destinations
// @Tags flights
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/flights/popular-routes [get]
func (h *FlightHandler) GetPopularRoutes(w http.ResponseWriter, r *http.Request) {
	// For now, return some mock popular routes
	// In a real implementation, you might want to track popular searches or have a predefined list
	popularRoutes := []map[string]interface{}{
		{
			"route":       "NYC-LAX",
			"origin":      "New York",
			"destination": "Los Angeles",
			"origin_code": "NYC",
			"dest_code":   "LAX",
			"avg_price":   250.00,
			"popularity":  95,
		},
		{
			"route":       "LHR-JFK",
			"origin":      "London",
			"destination": "New York",
			"origin_code": "LHR",
			"dest_code":   "JFK",
			"avg_price":   450.00,
			"popularity":  90,
		},
		{
			"route":       "LAX-NRT",
			"origin":      "Los Angeles",
			"destination": "Tokyo",
			"origin_code": "LAX",
			"dest_code":   "NRT",
			"avg_price":   650.00,
			"popularity":  85,
		},
		{
			"route":       "DXB-LHR",
			"origin":      "Dubai",
			"destination": "London",
			"origin_code": "DXB",
			"dest_code":   "LHR",
			"avg_price":   380.00,
			"popularity":  80,
		},
		{
			"route":       "SIN-SYD",
			"origin":      "Singapore",
			"destination": "Sydney",
			"origin_code": "SIN",
			"dest_code":   "SYD",
			"avg_price":   320.00,
			"popularity":  75,
		},
	}

	response := models.APIResponse{
		Success: true,
		Message: "Popular routes retrieved successfully",
		Data:    popularRoutes,
	}

	h.successResponse(w, response)
}

// GetAirports returns a list of airports for autocomplete
// @Summary Get airports
// @Description Get list of airports for search autocomplete
// @Tags flights
// @Produce json
// @Param query query string false "Search query for airport name or code"
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/flights/airports [get]
func (h *FlightHandler) GetAirports(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	// Mock airport data - in a real implementation, you'd search a database or external API
	allAirports := []map[string]interface{}{
		{"code": "JFK", "name": "John F. Kennedy International Airport", "city": "New York", "country": "USA"},
		{"code": "LAX", "name": "Los Angeles International Airport", "city": "Los Angeles", "country": "USA"},
		{"code": "LHR", "name": "Heathrow Airport", "city": "London", "country": "UK"},
		{"code": "CDG", "name": "Charles de Gaulle Airport", "city": "Paris", "country": "France"},
		{"code": "NRT", "name": "Narita International Airport", "city": "Tokyo", "country": "Japan"},
		{"code": "DXB", "name": "Dubai International Airport", "city": "Dubai", "country": "UAE"},
		{"code": "SIN", "name": "Singapore Changi Airport", "city": "Singapore", "country": "Singapore"},
		{"code": "SYD", "name": "Sydney Kingsford Smith Airport", "city": "Sydney", "country": "Australia"},
		{"code": "LOS", "name": "Murtala Muhammed International Airport", "city": "Lagos", "country": "Nigeria"},
		{"code": "CPT", "name": "Cape Town International Airport", "city": "Cape Town", "country": "South Africa"},
		{"code": "CAI", "name": "Cairo International Airport", "city": "Cairo", "country": "Egypt"},
		{"code": "NBO", "name": "Jomo Kenyatta International Airport", "city": "Nairobi", "country": "Kenya"},
	}

	var filteredAirports []map[string]interface{}

	// If query is provided, filter airports
	if query != "" {
		for _, airport := range allAirports {
			code := airport["code"].(string)
			name := airport["name"].(string)
			city := airport["city"].(string)

			// Simple case-insensitive search
			if containsIgnoreCase(code, query) ||
				containsIgnoreCase(name, query) ||
				containsIgnoreCase(city, query) {
				filteredAirports = append(filteredAirports, airport)
			}
		}
	} else {
		filteredAirports = allAirports
	}

	response := models.APIResponse{
		Success: true,
		Message: "Airports retrieved successfully",
		Data:    filteredAirports,
	}

	h.successResponse(w, response)
}

// Helper functions

func (h *FlightHandler) successResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *FlightHandler) errorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Success: false,
		Message: message,
	})
}

// containsIgnoreCase checks if a string contains a substring (case-insensitive)
func containsIgnoreCase(s, substr string) bool {
	s = toLower(s)
	substr = toLower(substr)
	return contains(s, substr)
}

func toLower(s string) string {
	var result []rune
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			result = append(result, r+32)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (len(substr) == 0 || indexOf(s, substr) >= 0)
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

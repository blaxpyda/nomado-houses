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

// ServiceHandler handles service requests
type ServiceHandler struct {
	serviceService service.ServiceService
	logger         *logger.Logger
}

// NewServiceHandler creates a new service handler
func NewServiceHandler(serviceService service.ServiceService, logger *logger.Logger) *ServiceHandler {
	return &ServiceHandler{serviceService: serviceService, logger: logger}
}

// GetAllServices handles GET /api/services
// @Summary Get all services
// @Description Get all available services, optionally filtered by category
// @Tags Services
// @Produce json
// @Param category query string false "Service category"
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /services [get]
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

// CreateService handles POST /api/services
func (h *ServiceHandler) CreateService(w http.ResponseWriter, r *http.Request) {
	var req models.CreateServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	service := &models.Service{
		ProviderID:    req.ProviderID,
		ServiceTypeID: req.ServiceTypeID,
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		Availability:  req.Availability,
	}

	if err := h.serviceService.CreateService(service); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Service created successfully",
		Data:    service,
	})
}

// UpdateService handles PUT /api/services/{id}
func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid service ID")
		return
	}

	var req models.UpdateServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	service := &models.Service{
		ID:            id,
		ProviderID:    req.ProviderID,
		ServiceTypeID: req.ServiceTypeID,
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		Availability:  req.Availability,
	}

	if err := h.serviceService.UpdateService(service); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Service updated successfully",
		Data:    service,
	})
}

// DeleteService handles DELETE /api/services/{id}
func (h *ServiceHandler) DeleteService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid service ID")
		return
	}

	if err := h.serviceService.DeleteService(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Service deleted successfully",
	})
}

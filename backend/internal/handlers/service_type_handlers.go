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

// ServiceTypeHandler handles service type-related requests
type ServiceTypeHandler struct {
	serviceTypeService service.ServiceTypeService
	logger             *logger.Logger
}

// NewServiceTypeHandler creates a new service type handler
func NewServiceTypeHandler(serviceTypeService service.ServiceTypeService, logger *logger.Logger) *ServiceTypeHandler {
	return &ServiceTypeHandler{
		serviceTypeService: serviceTypeService,
		logger:             logger,
	}
}

// respondWithError sends an error response
func (h *ServiceTypeHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJSON(w, code, models.ErrorResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// respondWithJSON sends a JSON response
func (h *ServiceTypeHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		h.logger.Error("Failed to encode JSON response", err)
	}
}

// GetAllServiceTypes handles GET /api/service-types
// @Summary Get all service types
// @Description Get all available service types from the database
// @Tags ServiceTypes
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /service-types [get]
func (h *ServiceTypeHandler) GetAllServiceTypes(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Fetching all service types from database")

	serviceTypes, err := h.serviceTypeService.GetAllServiceTypes()
	if err != nil {
		h.logger.Error("Failed to get service types", err)
		h.respondWithError(w, http.StatusInternalServerError, "Failed to retrieve service types")
		return
	}

	h.logger.Info("Successfully retrieved service types from database")
	h.respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Service types retrieved successfully",
		Data:    serviceTypes,
	})
}

// GetServiceTypeByID handles GET /api/service-types/{id}
// @Summary Get service type by ID
// @Description Get a specific service type from the database by its ID
// @Tags ServiceTypes
// @Produce json
// @Param id path int true "Service Type ID"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /service-types/{id} [get]
func (h *ServiceTypeHandler) GetServiceTypeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid service type ID", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid service type ID")
		return
	}

	h.logger.Info("Fetching service type with ID: " + vars["id"])

	serviceType, err := h.serviceTypeService.GetServiceTypeByID(id)
	if err != nil {
		h.logger.Error("Failed to get service type", err)
		h.respondWithError(w, http.StatusNotFound, "Service type not found")
		return
	}

	h.logger.Info("Successfully retrieved service type from database")
	h.respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Service type retrieved successfully",
		Data:    serviceType,
	})
}

// CreateServiceType handles POST /api/service-types
// @Summary Create service type
// @Description Create a new service type in the database (Admin only)
// @Tags ServiceTypes
// @Accept json
// @Produce json
// @Param request body models.CreateServiceTypeRequest true "Create service type request"
// @Security Bearer
// @Success 201 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /service-types [post]
func (h *ServiceTypeHandler) CreateServiceType(w http.ResponseWriter, r *http.Request) {
	var req models.CreateServiceTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request payload", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	h.logger.Info("Creating new service type: " + req.Name)

	serviceType := &models.ServiceType{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.serviceTypeService.CreateServiceType(serviceType); err != nil {
		h.logger.Error("Failed to create service type", err)
		h.respondWithError(w, http.StatusInternalServerError, "Failed to create service type")
		return
	}

	h.logger.Info("Successfully created service type in database")
	h.respondWithJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Service type created successfully",
		Data:    serviceType,
	})
}

// UpdateServiceType handles PUT /api/service-types/{id}
// @Summary Update service type
// @Description Update an existing service type in the database (Admin only)
// @Tags ServiceTypes
// @Accept json
// @Produce json
// @Param id path int true "Service Type ID"
// @Param request body models.UpdateServiceTypeRequest true "Update service type request"
// @Security Bearer
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /service-types/{id} [put]
func (h *ServiceTypeHandler) UpdateServiceType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid service type ID", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid service type ID")
		return
	}

	var req models.UpdateServiceTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request payload", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	h.logger.Info("Updating service type with ID: " + vars["id"])

	serviceType := &models.ServiceType{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.serviceTypeService.UpdateServiceType(serviceType); err != nil {
		h.logger.Error("Failed to update service type", err)
		h.respondWithError(w, http.StatusInternalServerError, "Failed to update service type")
		return
	}

	h.logger.Info("Successfully updated service type in database")
	h.respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Service type updated successfully",
		Data:    serviceType,
	})
}

// DeleteServiceType handles DELETE /api/service-types/{id}
// @Summary Delete service type
// @Description Delete a service type from the database (Admin only)
// @Tags ServiceTypes
// @Produce json
// @Param id path int true "Service Type ID"
// @Security Bearer
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /service-types/{id} [delete]
func (h *ServiceTypeHandler) DeleteServiceType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid service type ID", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid service type ID")
		return
	}

	h.logger.Info("Deleting service type with ID: " + vars["id"])

	if err := h.serviceTypeService.DeleteServiceType(id); err != nil {
		h.logger.Error("Failed to delete service type", err)
		h.respondWithError(w, http.StatusInternalServerError, "Failed to delete service type")
		return
	}

	h.logger.Info("Successfully deleted service type from database")
	h.respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Service type deleted successfully",
	})
}

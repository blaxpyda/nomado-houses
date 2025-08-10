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

// DestinationHandler handles destination requests
type DestinationHandler struct {
	destinationService service.DestinationService
	logger             *logger.Logger
}

// NewDestinationHandler creates a new destination handler
func NewDestinationHandler(destinationService service.DestinationService, logger *logger.Logger) *DestinationHandler {
	return &DestinationHandler{destinationService: destinationService, logger: logger}
}

// GetAllDestinations handles GET /api/destinations
// @Summary Get all destinations
// @Description Get all available destinations
// @Tags Destinations
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /destinations [get]
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
// @Summary Get destination by ID
// @Description Get destination by ID
// @Tags Destinations
// @Produce json
// @Param id path int true "Destination ID"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /destinations/{id} [get]
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

// CreateDestination handles POST /api/destinations
// @Summary Create destination
// @Description Create a new destination (Admin only)
// @Tags Destinations
// @Accept json
// @Produce json
// @Param request body models.CreateDestinationRequest true "Create destination request"
// @Security Bearer
// @Success 201 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /destinations [post]
func (h *DestinationHandler) CreateDestination(w http.ResponseWriter, r *http.Request) {
	var req models.CreateDestinationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	destination := &models.Destination{
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
	}

	if err := h.destinationService.CreateDestination(destination); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Destination created successfully",
		Data:    destination,
	})
}

// UpdateDestination handles PUT /api/destinations/{id}
// @Summary Update destination
// @Description Update destination (Admin only)
// @Tags Destinations
// @Accept json
// @Produce json
// @Param id path int true "Destination ID"
// @Param request body models.UpdateDestinationRequest true "Update destination request"
// @Security Bearer
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /destinations/{id} [put]
func (h *DestinationHandler) UpdateDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid destination ID")
		return
	}

	var req models.UpdateDestinationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	destination := &models.Destination{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
	}

	if err := h.destinationService.UpdateDestination(destination); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Destination updated successfully",
		Data:    destination,
	})
}

// DeleteDestination handles DELETE /api/destinations/{id}
// @Summary Delete destination
// @Description Delete destination (Admin only)
// @Tags Destinations
// @Produce json
// @Param id path int true "Destination ID"
// @Security Bearer
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /destinations/{id} [delete]
func (h *DestinationHandler) DeleteDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid destination ID")
		return
	}

	if err := h.destinationService.DeleteDestination(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Destination deleted successfully",
	})
}

// // ServiceHandler handles service requests
// type ServiceHandler struct {
// 	serviceService service.ServiceService
// 	logger         *logger.Logger
// }

// // NewServiceHandler creates a new service handler
// func NewServiceHandler(serviceService service.ServiceService, logger *logger.Logger) *ServiceHandler {
// 	return &ServiceHandler{serviceService: serviceService, logger: logger}
// }

// // GetAllServices handles GET /api/services
// // @Summary Get all services
// // @Description Get all available services, optionally filtered by category
// // @Tags Services
// // @Produce json
// // @Param category query string false "Service category"
// // @Success 200 {object} models.APIResponse
// // @Failure 500 {object} models.ErrorResponse
// // @Router /services [get]
// func (h *ServiceHandler) GetAllServices(w http.ResponseWriter, r *http.Request) {
// 	category := r.URL.Query().Get("category")

// 	var services []models.Service
// 	var err error

// 	if category != "" {
// 		services, err = h.serviceService.GetServicesByCategory(category)
// 	} else {
// 		services, err = h.serviceService.GetAllServices()
// 	}

// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Services retrieved successfully",
// 		Data:    services,
// 	})
// }

// // GetServiceByID handles GET /api/services/{id}
// func (h *ServiceHandler) GetServiceByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid service ID")
// 		return
// 	}

// 	service, err := h.serviceService.GetServiceByID(id)
// 	if err != nil {
// 		respondWithError(w, http.StatusNotFound, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Service retrieved successfully",
// 		Data:    service,
// 	})
// }

// // CreateService handles POST /api/services
// func (h *ServiceHandler) CreateService(w http.ResponseWriter, r *http.Request) {
// 	var req models.CreateServiceRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	service := &models.Service{
// 		ProviderID:    req.ProviderID,
// 		ServiceTypeID: req.ServiceTypeID,
// 		Name:          req.Name,
// 		Description:   req.Description,
// 		Price:         req.Price,
// 		Availability:  req.Availability,
// 	}

// 	if err := h.serviceService.CreateService(service); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, models.APIResponse{
// 		Success: true,
// 		Message: "Service created successfully",
// 		Data:    service,
// 	})
// }

// // UpdateService handles PUT /api/services/{id}
// func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid service ID")
// 		return
// 	}

// 	var req models.UpdateServiceRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	service := &models.Service{
// 		ID:            id,
// 		ProviderID:    req.ProviderID,
// 		ServiceTypeID: req.ServiceTypeID,
// 		Name:          req.Name,
// 		Description:   req.Description,
// 		Price:         req.Price,
// 		Availability:  req.Availability,
// 	}

// 	if err := h.serviceService.UpdateService(service); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Service updated successfully",
// 		Data:    service,
// 	})
// }

// // DeleteService handles DELETE /api/services/{id}
// func (h *ServiceHandler) DeleteService(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid service ID")
// 		return
// 	}

// 	if err := h.serviceService.DeleteService(id); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Service deleted successfully",
// 	})
// }

// // BookingHandler handles booking requests
// type BookingHandler struct {
// 	bookingService service.BookingService
// 	logger         *logger.Logger
// }

// // NewBookingHandler creates a new booking handler
// func NewBookingHandler(bookingService service.BookingService, logger *logger.Logger) *BookingHandler {
// 	return &BookingHandler{bookingService: bookingService, logger: logger}
// }

// // CreateBooking handles POST /api/bookings
// // @Summary Create booking
// // @Description Create a new booking for a service
// // @Tags Bookings
// // @Accept json
// // @Produce json
// // @Param request body models.CreateBookingRequest true "Create booking request"
// // @Security Bearer
// // @Success 201 {object} models.APIResponse
// // @Failure 400 {object} models.ErrorResponse
// // @Failure 401 {object} models.ErrorResponse
// // @Router /bookings [post]
// func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
// 	var req models.CreateBookingRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	// Get user ID from header (set by auth middleware)
// 	userIDStr := r.Header.Get("X-User-ID")
// 	userID, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
// 		return
// 	}

// 	booking := &models.Booking{
// 		UserID:           userID,
// 		ServiceID:        req.ServiceID,
// 		ProviderID:       req.ProviderID,
// 		BookingDateStart: req.BookingDateStart,
// 		BookingDateEnd:   req.BookingDateEnd,
// 		TotalPrice:       req.TotalPrice,
// 		Status:           "pending",
// 	}

// 	if err := h.bookingService.CreateBooking(booking); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, models.APIResponse{
// 		Success: true,
// 		Message: "Booking created successfully",
// 		Data:    booking,
// 	})
// }

// // GetUserBookings handles GET /api/bookings
// // @Summary Get user bookings
// // @Description Get all bookings for the authenticated user
// // @Tags Bookings
// // @Produce json
// // @Security Bearer
// // @Success 200 {object} models.APIResponse
// // @Failure 401 {object} models.ErrorResponse
// // @Router /bookings [get]
// func (h *BookingHandler) GetUserBookings(w http.ResponseWriter, r *http.Request) {
// 	// Get user ID from header (set by auth middleware)
// 	userIDStr := r.Header.Get("X-User-ID")
// 	userID, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
// 		return
// 	}

// 	bookings, err := h.bookingService.GetBookingsByUserID(userID)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Bookings retrieved successfully",
// 		Data:    bookings,
// 	})
// }

// // GetBookingByID handles GET /api/bookings/{id}
// func (h *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid booking ID")
// 		return
// 	}

// 	booking, err := h.bookingService.GetBookingByID(id)
// 	if err != nil {
// 		respondWithError(w, http.StatusNotFound, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Booking retrieved successfully",
// 		Data:    booking,
// 	})
// }

// // UpdateBookingStatus handles PUT /api/bookings/{id}/status
// func (h *BookingHandler) UpdateBookingStatus(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid booking ID")
// 		return
// 	}

// 	var req models.UpdateBookingStatusRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	if err := h.bookingService.UpdateBookingStatus(id, req.Status); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Booking status updated successfully",
// 	})
// }

// // ServiceTypeHandler handles service type requests
// type ServiceTypeHandler struct {
// 	serviceTypeService service.ServiceTypeService
// 	logger             *logger.Logger
// }

// // NewServiceTypeHandler creates a new service type handler
// func NewServiceTypeHandler(serviceTypeService service.ServiceTypeService, logger *logger.Logger) *ServiceTypeHandler {
// 	return &ServiceTypeHandler{serviceTypeService: serviceTypeService, logger: logger}
// }

// // GetAllServiceTypes handles GET /api/service-types
// func (h *ServiceTypeHandler) GetAllServiceTypes(w http.ResponseWriter, r *http.Request) {
// 	serviceTypes, err := h.serviceTypeService.GetAllServiceTypes()
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Service types retrieved successfully",
// 		Data:    serviceTypes,
// 	})
// }

// // GetServiceTypeByID handles GET /api/service-types/{id}
// func (h *ServiceTypeHandler) GetServiceTypeByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid service type ID")
// 		return
// 	}

// 	serviceType, err := h.serviceTypeService.GetServiceTypeByID(id)
// 	if err != nil {
// 		respondWithError(w, http.StatusNotFound, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Service type retrieved successfully",
// 		Data:    serviceType,
// 	})
// }

// // CreateServiceType handles POST /api/service-types
// func (h *ServiceTypeHandler) CreateServiceType(w http.ResponseWriter, r *http.Request) {
// 	var req models.CreateServiceTypeRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	serviceType := &models.ServiceType{
// 		Name:        req.Name,
// 		Description: req.Description,
// 	}

// 	if err := h.serviceTypeService.CreateServiceType(serviceType); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, models.APIResponse{
// 		Success: true,
// 		Message: "Service type created successfully",
// 		Data:    serviceType,
// 	})
// }

// // UpdateServiceType handles PUT /api/service-types/{id}
// func (h *ServiceTypeHandler) UpdateServiceType(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid service type ID")
// 		return
// 	}

// 	var req models.UpdateServiceTypeRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	serviceType := &models.ServiceType{
// 		ID:          id,
// 		Name:        req.Name,
// 		Description: req.Description,
// 	}

// 	if err := h.serviceTypeService.UpdateServiceType(serviceType); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Service type updated successfully",
// 		Data:    serviceType,
// 	})
// }

// // DeleteServiceType handles DELETE /api/service-types/{id}
// func (h *ServiceTypeHandler) DeleteServiceType(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid service type ID")
// 		return
// 	}

// 	if err := h.serviceTypeService.DeleteServiceType(id); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Message: "Service type deleted successfully",
// 	})
// }

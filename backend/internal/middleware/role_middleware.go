package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"nomado-houses/internal/models"
	"nomado-houses/internal/service"
	"strings"
)

// RoleMiddleware checks if user has required role
type RoleMiddleware struct {
	authService service.AuthService
	userService service.UserService
}

// NewRoleMiddleware creates a new role middleware
func NewRoleMiddleware(authService service.AuthService, userService service.UserService) *RoleMiddleware {
	return &RoleMiddleware{
		authService: authService,
		userService: userService,
	}
}

// RequireRole middleware that checks if user has the required role
func (rm *RoleMiddleware) RequireRole(allowedRoles ...models.UserRole) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				rm.unauthorizedResponse(w, "Authorization header required")
				return
			}

			// Extract token from "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				rm.unauthorizedResponse(w, "Invalid authorization header format")
				return
			}

			token := parts[1]

			// Validate token and get user ID
			userID, err := rm.authService.ValidateToken(token)
			if err != nil {
				rm.unauthorizedResponse(w, "Invalid token")
				return
			}

			// Get user details to check role
			user, err := rm.userService.GetUserByID(userID)
			if err != nil {
				rm.unauthorizedResponse(w, "User not found")
				return
			}

			// Check if user has any of the allowed roles
			hasPermission := false
			for _, allowedRole := range allowedRoles {
				if user.HasRole(allowedRole) {
					hasPermission = true
					break
				}
			}

			if !hasPermission {
				rm.forbiddenResponse(w, "Insufficient permissions")
				return
			}

			// Add user to request context
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequireAdmin middleware that checks if user is admin
func (rm *RoleMiddleware) RequireAdmin() func(http.Handler) http.Handler {
	return rm.RequireRole(models.RoleAdmin)
}

// RequireProvider middleware that checks if user is provider
func (rm *RoleMiddleware) RequireProvider() func(http.Handler) http.Handler {
	return rm.RequireRole(models.RoleProvider)
}

// RequireAdminOrProvider middleware that checks if user is admin or provider
func (rm *RoleMiddleware) RequireAdminOrProvider() func(http.Handler) http.Handler {
	return rm.RequireRole(models.RoleAdmin, models.RoleProvider)
}

// RequireAnyRole middleware that allows any authenticated user
func (rm *RoleMiddleware) RequireAnyRole() func(http.Handler) http.Handler {
	return rm.RequireRole(models.RoleUser, models.RoleAdmin, models.RoleProvider)
}

func (rm *RoleMiddleware) unauthorizedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Success: false,
		Message: message,
	})
}

func (rm *RoleMiddleware) forbiddenResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Success: false,
		Message: message,
	})
}

// GetUserFromContext extracts user from request context
func GetUserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value("user").(*models.User)
	return user, ok
}

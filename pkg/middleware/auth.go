package middleware

import (
    "net/http"
    "jukeBE/internal/service"
)

// CreateBasicAuthMiddleware creates a middleware that validates basic auth credentials
func CreateBasicAuthMiddleware(authService service.AuthService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            username, password, ok := r.BasicAuth()
            if !ok {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            // Validate credentials using your auth service
            if !authService.ValidateCredentials(username, password) {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}
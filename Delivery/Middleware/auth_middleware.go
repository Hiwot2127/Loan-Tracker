package Middleware

import (
    "context"
    "net/http"
    "strings"
    "Loan-Tracker/Infrastructure"
    "github.com/gin-gonic/gin"
)

type contextKey string

const userContextKey contextKey = "user"

// AuthMiddleware validates the JWT token and extracts the user email from it
func AuthMiddleware(tokenService *Infrastructure.TokenService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if (authHeader == "") {
                http.Error(w, "Authorization header is required", http.StatusUnauthorized)
                return
            }

            tokenString := strings.TrimPrefix(authHeader, "Bearer ")
            claims, err := tokenService.ValidateToken(tokenString)
            if err != nil {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            ctx := context.WithValue(r.Context(), userContextKey, claims.Email)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// GinAuthMiddleware is a wrapper to make AuthMiddleware compatible with gin.HandlerFunc
func GinAuthMiddleware(tokenService *Infrastructure.TokenService) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := tokenService.ValidateToken(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Set(string(userContextKey), claims.Email)
        c.Next()
    }
}

// GetUserFromContext extracts the user email from the context
func GetUserFromContext(ctx context.Context) string {
    if ctx == nil {
        return ""
    }
    email, _ := ctx.Value(userContextKey).(string)
    return email
}
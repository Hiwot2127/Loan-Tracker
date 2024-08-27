package Middleware

import (
    "net/http"
    "Loan-Tracker/Usecases"
    "github.com/gin-gonic/gin"
)

// AdminMiddleware checks if the user is authenticated and has admin privileges
func AdminMiddleware(userUsecase *Usecases.UserUsecase) gin.HandlerFunc {
    return func(c *gin.Context) {
        email := GetUserFromContext(c.Request.Context())
        if email == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        user, err := userUsecase.GetUserByEmail(email)
        if err != nil || !user.IsAdmin {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            return
        }

        c.Next()
    }
}
package Routers

import (
    "Loan-Tracker/Delivery/Controllers"
    "Loan-Tracker/Delivery/Middleware"
    "Loan-Tracker/Infrastructure"
    "github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.Engine, userController *Controllers.UserController, tokenService *Infrastructure.TokenService) {
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/register", userController.RegisterUser)
        userRoutes.POST("/login", userController.LoginUser) 
    }
  // Apply middleware only to routes that require authentication
    authRoutes := userRoutes.Group("/auth")
    authRoutes.Use(Middleware.GinAuthMiddleware(tokenService))
    {
        authRoutes.GET("/verify-email/:email", userController.VerifyEmail)
        authRoutes.POST("/refresh-token", userController.RefreshToken)
        authRoutes.GET("/profile/:id", userController.GetUserProfile)
        authRoutes.POST("/password-reset", userController.PasswordResetRequest)
        authRoutes.POST("/password-update", userController.PasswordUpdateAfterReset)
    }
    }
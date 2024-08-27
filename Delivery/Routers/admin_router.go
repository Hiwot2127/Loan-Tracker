package Routers

import (
    "Loan-Tracker/Delivery/Controllers"
    "Loan-Tracker/Delivery/Middleware"
    "Loan-Tracker/Infrastructure"
    "github.com/gin-gonic/gin"
)

func SetAdminRoutes(router *gin.RouterGroup, adminController *Controllers.AdminController, adminMiddleware gin.HandlerFunc, tokenService *Infrastructure.TokenService) {
    adminRoutes := router.Group("/admin")
    adminRoutes.Use(Middleware.GinAuthMiddleware(tokenService))
    adminRoutes.Use(adminMiddleware)
    {
        adminRoutes.GET("/users", adminController.GetAllUsers)
        adminRoutes.DELETE("/users/:id", adminController.DeleteUser)
    }
}
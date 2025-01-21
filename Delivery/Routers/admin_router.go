package Routers

import (
    "Loan-Tracker/Delivery/Controllers"
    "Loan-Tracker/Delivery/Middleware"
    "Loan-Tracker/Infrastructure"
    "github.com/gin-gonic/gin"
)

func SetAdminRoutes(router *gin.RouterGroup, adminController *Controllers.AdminController, loanController *Controllers.LoanController, adminMiddleware gin.HandlerFunc, tokenService *Infrastructure.TokenService) {
    adminRoutes := router.Group("/admin")
    adminRoutes.Use(Middleware.GinAuthMiddleware(tokenService))
    adminRoutes.Use(adminMiddleware)
    {
        adminRoutes.GET("/users", adminController.GetAllUsers)
        adminRoutes.DELETE("/users/:id", adminController.DeleteUser)

        // Loan routes
        adminRoutes.GET("/loans", loanController.GetAllLoans)
        adminRoutes.PATCH("/loans/:id/status", loanController.UpdateLoanStatus)
        adminRoutes.DELETE("/loans/:id", loanController.DeleteLoan)
    }
}
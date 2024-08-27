package Controllers

import (
    "net/http"
    "Loan-Tracker/Usecases"
    "github.com/gin-gonic/gin"
)

// AdminController handles admin-related HTTP requests
type AdminController struct {
    adminUsecase *Usecases.AdminUsecase
}

// NewAdminController creates a new instance of AdminController
func NewAdminController(adminUsecase *Usecases.AdminUsecase) *AdminController {
    return &AdminController{
        adminUsecase: adminUsecase,
    }
}

// GetAllUsers handles the HTTP request to retrieve all users
func (c *AdminController) GetAllUsers(ctx *gin.Context) {
    users, err := c.adminUsecase.GetAllUsers()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, users)
}

// DeleteUser handles the HTTP request to delete a user
func (c *AdminController) DeleteUser(ctx *gin.Context) {
    id := ctx.Param("id")

    err := c.adminUsecase.DeleteUser(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusNoContent)
}
package Controllers

import (
    "net/http"
    "Loan-Tracker/Domain"
    "Loan-Tracker/Usecases"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanController struct {
    loanUsecase *Usecases.LoanUsecase
}

func NewLoanController(loanUsecase *Usecases.LoanUsecase) *LoanController {
    return &LoanController{
        loanUsecase: loanUsecase,
    }
}

func (c *LoanController) ApplyForLoan(ctx *gin.Context) {
    var loan Domain.Loan
    if err := ctx.ShouldBindJSON(&loan); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    loan.ID = primitive.NewObjectID()
    loan.Status = "pending"
    if err := c.loanUsecase.ApplyForLoan(ctx, &loan); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"status": "application submitted"})
}

func (c *LoanController) GetLoanStatus(ctx *gin.Context) {
    id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
        return
    }
    loan, err := c.loanUsecase.GetLoanStatus(ctx, id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) GetAllLoans(ctx *gin.Context) {
    filter := make(map[string]interface{})
    if status := ctx.Query("status"); status != "" {
        filter["status"] = status
    }
    loans, err := c.loanUsecase.GetAllLoans(ctx, filter)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, loans)
}

func (c *LoanController) UpdateLoanStatus(ctx *gin.Context) {
    id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
        return
    }
    var status struct {
        Status string `json:"status"`
    }
    if err := ctx.ShouldBindJSON(&status); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.loanUsecase.UpdateLoanStatus(ctx, id, status.Status); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"status": "loan status updated"})
}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {
    id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
        return
    }
    if err := c.loanUsecase.DeleteLoan(ctx, id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"status": "loan deleted"})
}
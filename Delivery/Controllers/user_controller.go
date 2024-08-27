package Controllers

import (
    "net/http"
    "Loan-Tracker/Domain"
    "Loan-Tracker/Usecases"
    "Loan-Tracker/Infrastructure"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// UserController handles user-related HTTP requests.
type UserController struct {
    userUsecase    *Usecases.UserUsecase
    emailService   *Infrastructure.EmailService
    tokenService   *Infrastructure.TokenService
}

// NewUserController creates a new UserController instance.
func NewUserController(userUsecase *Usecases.UserUsecase, emailService *Infrastructure.EmailService, tokenService *Infrastructure.TokenService) *UserController {
    return &UserController{
        userUsecase:  userUsecase,
        emailService: emailService,
        tokenService: tokenService,
    }
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
    var user Domain.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.ID = primitive.NewObjectID()

    if err := c.userUsecase.RegisterUser(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := c.emailService.SendVerificationEmail(user.Email); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
        return
    }

    ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) VerifyEmail(ctx *gin.Context) {
    email := ctx.Param("email")

    if err := c.userUsecase.VerifyEmail(email); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := ctx.ShouldBindJSON(&credentials); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := c.userUsecase.GetUserByEmail(credentials.Email)
    if err != nil || user.Password != credentials.Password {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if !user.Verified {
        ctx.JSON(http.StatusForbidden, gin.H{"error": "Email not verified"})
        return
    }

    accessToken, refreshToken, err := c.tokenService.GenerateTokens(user.ID.Hex())
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "access_token":  accessToken,
        "refresh_token": refreshToken,
    })
}

func (c *UserController) RefreshToken(ctx *gin.Context) {
    var request struct {
        RefreshToken string `json:"refresh_token"`
    }
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    claims, err := c.tokenService.ValidateToken(request.RefreshToken)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
        return
    }

    userID, err := primitive.ObjectIDFromHex(claims.UserID)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    newAccessToken, err := c.userUsecase.RefreshToken(userID)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to refresh token"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "access_token": newAccessToken,
    })
}

func (c *UserController) GetUserProfile(ctx *gin.Context) {
    userID := ctx.Param("id")

    // Convert userID from string to primitive.ObjectID
    objectID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := c.userUsecase.GetUserProfile(objectID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (c *UserController) PasswordResetRequest(ctx *gin.Context) {
    var request struct {
        Email string `json:"email"`
    }
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.userUsecase.PasswordResetRequest(request.Email); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := c.emailService.SendPasswordResetEmail(request.Email); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send password reset email"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Password reset link sent"})
}

func (c *UserController) PasswordUpdateAfterReset(ctx *gin.Context) {
    var request struct {
        Token       string `json:"token"`
        NewPassword string `json:"new_password"`
    }
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, err := primitive.ObjectIDFromHex(request.Token)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
        return
    }

    if err := c.userUsecase.PasswordUpdateAfterReset(userID, request.NewPassword); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
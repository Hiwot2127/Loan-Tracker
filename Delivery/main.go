package main

import (
    "log"
    "os"
    "Loan-Tracker/Infrastructure"
    "Loan-Tracker/Delivery/Controllers"
    "Loan-Tracker/Delivery/Middleware"
    "Loan-Tracker/Delivery/Routers"
    "Loan-Tracker/Repositories"
    "Loan-Tracker/Usecases"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
)

func main() {
     // Load environment variables from .env file
     err := godotenv.Load()
     if err != nil {
         log.Fatalf("Error loading .env file")
     }
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPort := os.Getenv("SMTP_PORT")
    smtpUser := os.Getenv("SMTP_USER")
    smtpPass := os.Getenv("SMTP_PASS")

    // Initialize the database connection
    dbURI := os.Getenv("DB_URI")
    dbName := os.Getenv("DB_NAME")
    // Validate environment variables
    if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" || dbURI == "" || dbName == "" {
        log.Fatalf("Missing required environment variables")
    }

    db, err := Infrastructure.NewDatabase(dbURI, dbName)
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    defer db.Close()

    userRepository := Repositories.NewUserRepository(db.Database)
    loanRepository := Repositories.NewLoanRepository(db.Database)
    emailService := Infrastructure.NewEmailService(smtpHost, smtpPort, smtpUser, smtpPass)
    tokenService := Infrastructure.NewTokenService()

    userUsecase := Usecases.NewUserUsecase(userRepository, tokenService, emailService)
    loanUsecase := Usecases.NewLoanUsecase(loanRepository)
    userController := Controllers.NewUserController(userUsecase, emailService, tokenService)
    loanController := Controllers.NewLoanController(loanUsecase)

    adminUsecase := Usecases.NewAdminUsecase(userRepository)
    adminController := Controllers.NewAdminController(adminUsecase)
    adminMiddleware := Middleware.AdminMiddleware(userUsecase)
    router := gin.Default()
    // Set up user routes
    Routers.SetUserRoutes(router, userController, loanController, tokenService)

    // Set up admin routes with middleware
    authRoutes := router.Group("/auth")
    authRoutes.Use(Middleware.GinAuthMiddleware(tokenService))
    Routers.SetAdminRoutes(authRoutes, adminController, loanController, adminMiddleware, tokenService)

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
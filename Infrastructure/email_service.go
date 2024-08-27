package Infrastructure

import (
    "fmt"
    "net/smtp"
)

// EmailService struct holds the SMTP server configuration
type EmailService struct {
    smtpHost string
    smtpPort string
    auth     smtp.Auth
}

// NewEmailService initializes a new EmailService with the given SMTP configuration
func NewEmailService(smtpHost, smtpPort, smtpUser, smtpPass string) *EmailService {
    // Create SMTP authentication
    auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
    return &EmailService{
        smtpHost: smtpHost,
        smtpPort: smtpPort,
        auth:     auth,
    }
}

// SendEmail sends an email with the given recipient, subject, and body
func (e *EmailService) SendEmail(to, subject, body string) error {
    from := "no-reply@loan-tracker.com"
    msg := []byte("To: " + to + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" + body + "\r\n")

    addr := fmt.Sprintf("%s:%s", e.smtpHost, e.smtpPort)
    err := smtp.SendMail(addr, e.auth, from, []string{to}, msg)
    if err != nil {
        fmt.Printf("SMTP error: %v\n", err)
    }
    return err
}

// SendPasswordResetEmail sends a password reset email to the user
func (e *EmailService) SendPasswordResetEmail(to string) error {
    subject := "Password Reset Request" // Email subject
    body := "Please click the link below to reset your password:\n\n" +
        "http://example.com/reset-password?email=" + to // Email body with reset link
    return e.SendEmail(to, subject, body) // Send the password reset email
}
// SendVerificationEmail sends a verification email to the user
func (e *EmailService) SendVerificationEmail(to string) error {
    subject := "Email Verification"
    body := "Please click the link below to verify your email address:\n\n" +
        "http://example.com/verify-email?email=" + to
    err := e.SendEmail(to, subject, body)
    if err != nil {
        fmt.Printf("Failed to send verification email: %v\n", err)
    }
    return err
}
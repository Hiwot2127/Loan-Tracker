package Usecases

import (
	"Loan-Tracker/Domain"
	"Loan-Tracker/Infrastructure"
	"Loan-Tracker/Repositories"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
    userRepository *Repositories.UserRepository
    emailService   *Infrastructure.EmailService
    tokenService   *Infrastructure.TokenService
}

// NewUserUsecase creates a new UserUsecase
func NewUserUsecase(userRepository *Repositories.UserRepository,tokenService *Infrastructure.TokenService, emailService *Infrastructure.EmailService) *UserUsecase {
    return &UserUsecase{
        userRepository: userRepository,
        emailService:   emailService,
        tokenService:   tokenService,
    }
}

// RegisterUser registers a new user in the system
func (u *UserUsecase) RegisterUser(user *Domain.User) error {
    return u.userRepository.Create(user)
}

// VerifyEmail verifies the user's email by their email address
func (u *UserUsecase) VerifyEmail(email string) error {
    user, err := u.userRepository.FindByEmail(email)
    if err != nil {
        return err
    }
    if user == nil {
        return errors.New("user not found")
    }
    user.Verified = true
    user.UpdatedAt = time.Now()
    return u.userRepository.Update(user)
}

// LoginUser logs in a user by their email and password
func (u *UserUsecase) LoginUser(email, password string) (*Domain.User, error) {
    user, err := u.userRepository.FindByEmail(email)
    if err != nil {
        return nil, err
    }
    if user == nil || user.Password != password {
        return nil, errors.New("invalid email or password")
    }
    return user, nil
}

// RefreshToken refreshes the user's authentication token
func (u *UserUsecase) RefreshToken(userID primitive.ObjectID) (string, error) {
    user, err := u.userRepository.FindByID(userID)
    if err != nil {
        return "", err
    }
    if user == nil {
        return "", errors.New("user not found")
    }
    token, err := u.tokenService.GenerateToken(user.ID.Hex())
    if err != nil {
        return "", err
    }
    return token, nil
}

// GetUserProfile retrieves the user's profile by their ID
func (u *UserUsecase) GetUserProfile(userID primitive.ObjectID) (*Domain.User, error) {
    user, err := u.userRepository.FindByID(userID)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, errors.New("user not found")
    }
    return user, nil
}
// PasswordResetRequest initiates a password reset request for the user
func (u *UserUsecase) PasswordResetRequest(email string) error {
    user, err := u.userRepository.FindByEmail(email)
    if err != nil {
        return err
    }
    if user == nil {
        return errors.New("user not found")
    }
    return u.emailService.SendPasswordResetEmail(user.Email)
}

// PasswordUpdateAfterReset updates the user's password after a reset
func (u *UserUsecase) PasswordUpdateAfterReset(userID primitive.ObjectID, newPassword string) error {
    user, err := u.userRepository.FindByID(userID)
    if err != nil {
        return err
    }
    if user == nil {
        return errors.New("user not found")
    }
    user.Password = newPassword
    user.UpdatedAt = time.Now()
    return u.userRepository.Update(user)
}
// GetUserByEmail retrieves a user by their email
func (u *UserUsecase) GetUserByEmail(email string) (*Domain.User, error) {
    return u.userRepository.FindByEmail(email)
}
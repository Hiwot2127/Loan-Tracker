package Usecases

import (
    "Loan-Tracker/Domain"
	"Loan-Tracker/Repositories"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminUsecase struct {
    userRepository *Repositories.UserRepository
}

// NewAdminUsecase creates a new AdminUsecase
func NewAdminUsecase(userRepo *Repositories.UserRepository) *AdminUsecase {
    return &AdminUsecase{
        userRepository: userRepo,
    }
}

// GetAllUsers retrieves all users from the repository
func (u *AdminUsecase) GetAllUsers() ([]Domain.User, error) {
    return u.userRepository.GetAllUsers()
}

// DeleteUser deletes a user by their ID
func (u *AdminUsecase) DeleteUser(id string) error {
    userID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    return u.userRepository.Delete(userID)
}
package Domain

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system
type User struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`    
    Email     string             `json:"email" bson:"email"`         
    Password  string             `json:"password" bson:"password"`   
    FirstName string             `json:"first_name" bson:"first_name"` 
    LastName  string             `json:"last_name" bson:"last_name"`   
    Verified  bool               `json:"verified" bson:"verified"`   
    IsAdmin   bool               `json:"is_admin" bson:"is_admin"`
    CreatedAt time.Time          `json:"created_at" bson:"created_at"` 
    UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"` 
}

// UserRepository defines the methods that any
// data storage provider needs to implement to get
// and store users.
type UserRepository interface {
    Create(user *User) error
    FindByID(userID primitive.ObjectID) (*User, error)
    FindByEmail(email string) (*User, error)
    Update(user *User) error
    Delete(userID primitive.ObjectID) error
    GetAllUsers() ([]User, error)
}
// UserUsecase defines the methods that any
// user use case needs to implement.
type UserUsecase interface {
    RegisterUser(user *User) error
    VerifyEmail(userID primitive.ObjectID) error
    LoginUser(email, password string) (*User, error)
    RefreshToken(userID primitive.ObjectID) (string, error)
    GetUserProfile(userID primitive.ObjectID) (*User, error)
    PasswordResetRequest(email string) error
    PasswordUpdateAfterReset(userID primitive.ObjectID, newPassword string) error
}
// AdminUsecase defines the methods that any
// admin use case needs to implement.
type AdminUsecase interface {
    GetAllUsers() ([]User, error)
    DeleteUser(id string) error
}
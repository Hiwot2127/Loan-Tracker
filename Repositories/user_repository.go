package Repositories

import (
    "Loan-Tracker/Domain"
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
)

type UserRepository struct {
    collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
    return &UserRepository{
        collection: db.Collection("users"),
    }
}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *Domain.User) error {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    _, err := r.collection.InsertOne(context.Background(), user)
    if err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }
    return nil
}

// FindByID retrieves a user by their ID
func (r *UserRepository) FindByID(userID primitive.ObjectID) (*Domain.User, error) {
    var user Domain.User
    err := r.collection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to find user by ID: %w", err)
    }
    return &user, nil
}

// FindByEmail retrieves a user by their email
func (r *UserRepository) FindByEmail(email string) (*Domain.User, error) {
    var user Domain.User
    err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to find user by email: %w", err)
    }
    return &user, nil
}

// Update modifies an existing user in the database
func (r *UserRepository) Update(user *Domain.User) error {
    user.UpdatedAt = time.Now()
    _, err := r.collection.UpdateOne(
        context.Background(),
        bson.M{"_id": user.ID},
        bson.M{"$set": user},
        options.Update().SetUpsert(true),
    )
    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }
    return nil
}

// Delete removes a user from the database by their ID
func (r *UserRepository) Delete(userID primitive.ObjectID) error {
    _, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": userID})
    if err != nil {
        return fmt.Errorf("failed to delete user: %w", err)
    }
    return nil
}
// GetAllUsers retrieves all users from the database
func (r *UserRepository) GetAllUsers() ([]Domain.User, error) {
    var users []Domain.User
    cursor, err := r.collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, fmt.Errorf("failed to get all users: %w", err)
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var user Domain.User
        if err := cursor.Decode(&user); err != nil {
            return nil, fmt.Errorf("failed to decode user: %w", err)
        }
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        return nil, fmt.Errorf("cursor error: %w", err)
    }

    return users, nil
}
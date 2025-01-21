package Repositories

import (
    "context"
    "Loan-Tracker/Domain"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanRepository struct {
    collection *mongo.Collection
}

func NewLoanRepository(db *mongo.Database) *LoanRepository {
    return &LoanRepository{
        collection: db.Collection("loans"),
    }
}

func (r *LoanRepository) CreateLoan(ctx context.Context, loan *Domain.Loan) error {
    _, err := r.collection.InsertOne(ctx, loan)
    return err
}

func (r *LoanRepository) GetLoanByID(ctx context.Context, id primitive.ObjectID) (*Domain.Loan, error) {
    var loan Domain.Loan
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&loan)
    return &loan, err
}

func (r *LoanRepository) GetAllLoans(ctx context.Context, filter bson.M) ([]Domain.Loan, error) {
    var loans []Domain.Loan
    cursor, err := r.collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    err = cursor.All(ctx, &loans)
    return loans, err
}

func (r *LoanRepository) UpdateLoanStatus(ctx context.Context, id primitive.ObjectID, status string) error {
    _, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status}})
    return err
}

func (r *LoanRepository) DeleteLoan(ctx context.Context, id primitive.ObjectID) error {
    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}
package Domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Loan struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"`
    UserID     primitive.ObjectID `bson:"user_id"`
    Amount     float64            `bson:"amount"`
    Status     string             `bson:"status"`
    CreatedAt  int64              `bson:"created_at"`
    UpdatedAt  int64              `bson:"updated_at"`
}
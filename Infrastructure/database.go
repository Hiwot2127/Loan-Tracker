package Infrastructure

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
)

// Database holds the MongoDB client and database instance
type Database struct {
    Client   *mongo.Client
    Database *mongo.Database
}

// NewDatabase initializes a new database connection
func NewDatabase(uri, dbName string) (*Database, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
    }

    db := client.Database(dbName)
    return &Database{
        Client:   client,
        Database: db,
    }, nil
}

// Close closes the database connection
func (d *Database) Close() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    return d.Client.Disconnect(ctx)
}
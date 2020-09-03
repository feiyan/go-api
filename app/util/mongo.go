package util

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        fmt.Println("MongoDB Connect Failure")
    }

    err = mongoClient.Ping(context.TODO(), nil)
    if err != nil {
        fmt.Println("MongoDB Connect Failure")
    }
    fmt.Println("Connected to MongoDB!")
}

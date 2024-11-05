package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// InitMongoDB initializes the MongoDB connection
func InitMongoDB() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: no .env file found")
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	collection = client.Database("golang_db").Collection("todos")
	return nil
}
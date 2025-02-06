package db

import (
	"context"
	"errors"
	"karim/http_server/logger"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func GetDB() *mongo.Client {
	if db == nil {
		logger.WarnLogger.Println("DB is not initialized. Did you call Setup()?")
	}
	return db
}

func Setup() (*mongo.Client, error) {
	if db != nil {
		return nil, errors.New("DB already set up")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, errors.New("Failed to get environment variable: MONGODB_URI")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db = client

	logger.InfoLogger.Println("Connected to MongoDB")
	return db, nil
}

func CloseDB() {
	if db != nil {
		db.Disconnect(context.Background())
		logger.InfoLogger.Println("Disconnected from MongoDB")
		db = nil
	}
}

package db_test

import (
	"context"
	"fmt"
	"karim/http_server/db"
	"os"
	"testing"
)

func TestSetup(t *testing.T) {
	fmt.Println("Calling db.Setup()")
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
	client, err := db.Setup()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Fatalf("Expected a mongo.Client, got nil")
	}

	fmt.Println("Successfully connected to MongoDB")

	// Clean up
	fmt.Println("Disconnecting from MongoDB")
	err = client.Disconnect(context.Background())
	if err != nil {
		t.Fatalf("Expected no error on disconnect, got %v", err)
	}

	fmt.Println("Successfully disconnected from MongoDB")
}

func TestSetupMissingEnvVar(t *testing.T) {
	// Unset the environment variable for testing
	fmt.Println("Unsetting MONGODB_URI environment variable")
	os.Unsetenv("MONGODB_URI")

	fmt.Println("Calling db.Setup()")
	client, err := db.Setup()
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}

	if client != nil {
		t.Fatalf("Expected nil client, got %v", client)
	}

	fmt.Println("Received expected error due to missing MONGODB_URI")
}

func TestGetDB(t *testing.T) {
	fmt.Println("Calling db.GetDB() without setup")
	client := db.GetDB()
	if client != nil {
		t.Fatalf("Expected nil client, got %v", client)
	}

	fmt.Println("Calling db.Setup()")
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
	_, err := db.Setup()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	fmt.Println("Calling db.GetDB() after setup")
	client = db.GetDB()
	if client == nil {
		t.Fatalf("Expected a mongo.Client, got nil")
	}

	fmt.Println("Successfully retrieved MongoDB client")

	// Clean up
	fmt.Println("Disconnecting from MongoDB")
	err = client.Disconnect(context.Background())
	if err != nil {
		t.Fatalf("Expected no error on disconnect, got %v", err)
	}

	fmt.Println("Successfully disconnected from MongoDB")
}

func TestCloseDB(t *testing.T) {
	fmt.Println("Calling db.Setup()")
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
	client, err := db.Setup()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Fatalf("Expected a mongo.Client, got nil")
	}

	fmt.Println("Calling db.CloseDB()")
	db.CloseDB()

	fmt.Println("Calling db.GetDB() after close")
	client = db.GetDB()
	if client != nil {
		t.Fatalf("Expected nil client, got %v", client)
	}

	fmt.Println("Successfully closed MongoDB connection")
}

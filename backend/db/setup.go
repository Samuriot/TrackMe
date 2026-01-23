package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
	"os"
)

var database *mongo.Database
var dbName string

func Init() (*mongo.Database, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error: env not loaded properly")
	}

	mongo_url := os.Getenv("MONGO_URL")
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "trackme"
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongo_url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect error: %w", err)
	}

	// Ping the database
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("mongo ping error: %w", err)
	}

	// Store the database globally
	database = client.Database(dbName)
	return database, nil
}

// Close disconnects the MongoDB client
func Close() error {
	if database == nil {
		return nil
	}
	return database.Client().Disconnect(context.Background())
}

// GetDB returns the database object
func GetDB() *mongo.Database {
	return database
}

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

var client *mongo.Client
var dbName string

func Init() error {
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
		return err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return nil
}

// Close gracefully disconnects the MongoDB client
func Close() error {
	if client == nil {
		return nil
	}
	return client.Disconnect(context.Background())
}

func SetupQuery() (*mongo.Client, string) {
	return client, dbName
}

package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var client *mongo.Client

func Init() error {
  err := godotenv.Load();

  if err != nil {
    log.Fatal("Error: env not loaded properly")
  }

  mongo_url := os.Getenv("MONGO_URL")

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

// GetClient returns the MongoDB client
func GetClient() *mongo.Client {
  return client
}

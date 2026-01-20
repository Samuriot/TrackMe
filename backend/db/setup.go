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

func Init() {
  err := godotenv.Load();

  if err != nil {
    log.Fatal("Error: env not loaded properly")
  }

  mongo_url := os.Getenv("MONGO_URL")
  // Use the SetServerAPIOptions() method to set the version of the Stable API on the client
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI(mongo_url).SetServerAPIOptions(serverAPI)

  // Create a new client and connect to the server
  client, err := mongo.Connect(opts)
  if err != nil {
    panic(err)
  }

  defer func() {
    if err = client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()

  // Send a ping to confirm a successful connection
  if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
    panic(err)
  }

  fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

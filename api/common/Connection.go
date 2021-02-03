package common

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func Connection() (*mongo.Database, error) {
	Host := os.Getenv("MONGO_HOST")

	clientOpt := options.Client().ApplyURI(Host)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOpt)

	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client.Database("authentication"), nil
}
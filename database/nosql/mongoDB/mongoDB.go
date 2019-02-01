package mongoDB

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"context"
	"time"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
)

var (
	client *mongo.Client
)



type MongoClient struct {
	*mongo.Client
}

// Init Mongo.
func initMongo() {
	// Get a new driver.
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("mongo dirver init failure.")
		panic(err)
	}

	// Check connect mongoDB
	pingCtx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		log.Fatal("ping mongo failure.")
		panic(err)
	}

	collection := client.Database("").Collection("")
	fmt.Println("Connected to MongoDB! DB Name: ", collection.Name())
}

// Disconnect mongoDB.
func (c *MongoClient) CloseMongo() {
	err := c.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println("Connection to MongoDB closed")

}
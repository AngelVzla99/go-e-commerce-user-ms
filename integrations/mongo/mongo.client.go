package mongo_client

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()
var client *mongo.Client
var database = "user-ms"

type MongoClient struct {
	client *mongo.Client
}

func createClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	// Connect to the MongoDB cluster
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	// Ping the primary
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}
	return client
}

// start: singleton logic

var instance *MongoClient
var once sync.Once

func NewMongoClient() *MongoClient {
	once.Do(func() {
		instance = &MongoClient{
			client: createClient(),
		}
	})
	return instance
}

// end: singleton logic

func (c *MongoClient) GetCollection(collectionName string) *mongo.Collection {
	return c.client.Database(database).Collection(collectionName)
}

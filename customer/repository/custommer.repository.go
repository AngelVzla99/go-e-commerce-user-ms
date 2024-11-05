package repository

import (
	"context"

	mongo_client "github.com/AngelVzla99/e-commerce-user-ms/integrations/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCustomerRepository() *CustomerRepository {
	mongoClient := mongo_client.NewMongoClient()
	return &CustomerRepository{
		collection: mongoClient.GetCollection("customers"),
		ctx:        context.TODO(),
	}
}

func (repository *CustomerRepository) CreateCustomers(dto CreateCustomerDbe) (string, error) {
	dbResponse, dbError := repository.collection.InsertOne(repository.ctx, dto)

	if dbError != nil {
		return "", dbError
	}

	return dbResponse.InsertedID.(primitive.ObjectID).Hex(), nil
}

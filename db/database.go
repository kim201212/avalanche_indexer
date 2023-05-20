package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TransactionCollection *mongo.Collection

func Connect(str string) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	TransactionCollection = client.Database("avalanche").Collection(str)
	CreateUniqueIndex(TransactionCollection, "hash")
}

func CreateUniqueIndex(collection *mongo.Collection, field string) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: field, Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
}

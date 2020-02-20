package basic

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

// Pelajar struct
type Pelajar struct {
	Name  string `bson:"Name"`
	Grade int    `bson:"Grade"`
}

func ConnectMongo() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

func Insert() {
	db, err := ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, Pelajar{"Kautsar", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, Pelajar{"ilham", 25})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("insert success")
}

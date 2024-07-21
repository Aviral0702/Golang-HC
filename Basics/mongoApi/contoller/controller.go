package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoUrl = "mongodb+srv://aviralasthana0704:Aviral123@cluster0.0iuvkbz.mongodb.net/"
const dbName = "Netflix"
const collectionName = "watchList"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(mongoUrl)

	//connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection successful")
	collection = client.Database(dbName).Collection(collectionName)

	//collection instance
	fmt.Println("Collection instance created")

}

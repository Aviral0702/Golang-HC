package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Aviral0702/mongoApi/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func main() {

}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie to the database with id: ", inserted.InsertedID)
}

package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Aviral0702/mongoApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie to the database with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated one movie to the database with count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got delete with delete account: ", deleteCount)
}

func deleteAllMovies() int64 {
	const deletedResult,err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all movies with count: ", deletedResult.DeletedCount,nil)
	return deletedResult.DeletedCount
}

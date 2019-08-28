package db

import (
	"context"
	"time"

	"../model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDb is our mongo db struct
type MongoDb struct {
	MongoCollection *mongo.Collection
	MongoClient     *mongo.Client
}

// ConnectToMongoDatabase get connectoin to a mongo database TODO abstract hard coded values for db connx
func ConnectToMongoDatabase() UserDb {
	mdb := MongoDb{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("http://localhost:27017")
	mdb.MongoClient, _ = mongo.Connect(ctx, clientOptions)
	mdb.MongoCollection = mdb.MongoClient.Database("MPG").Collection("Users")
}

func CreateUser(user model.User) {

}

func ReadUser(username string) {

}

func UpdateUser(id string) {

}

func DeleteUser(id string) {

}

package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"../model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDb is our mongo db struct
type MongoDb struct {
	MongoCollection *mongo.Collection
	MongoClient     *mongo.Client
}

const ShortTimeout = 5 * time.Second
const MediumTimeout = 10 * time.Second
const LargeTimeout = 30 * time.Second

// ConnectToMongoDatabase get connectoin to a mongo database TODO abstract hard coded values for db connx
func ConnectToMongoDatabase() UserDb {
	mdb := MongoDb{}

	ctx, cancel := context.WithTimeout(context.Background(), MediumTimeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI("http://localhost:27017")
	mdb.MongoClient, _ = mongo.Connect(ctx, clientOptions)
	mdb.MongoCollection = mdb.MongoClient.Database("MPG").Collection("Users")
}

func (mdb MongoDb) InsertOne(user model.User) {
	ctx, cancel := context.WithTimeout(context.Background(), ShortTimeout)
	defer cancel()

	document := bson.M{
		"username":  user.Username,
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
		"password":  user.Password,
	}

	insertResult, err := mdb.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted new user: ", insertResult.InsertedID)
}

func (mdb MongoDb) FindOne(username string) (result model.User) {
	ctx, cancel := context.WithTimeout(context.Background(), ShortTimeout)
	defer cancel()

	filter := bson.D{{Key: "username", Value: username}}
	err := mdb.MongoCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Fatal("whoops, couldn't find user")
	}
	return result
}

func UpdateUser(id string) {

}

func DeleteUser(id string) {

}

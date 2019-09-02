package db

import "../model"

// ConnectToDatabase connects to database and returns db to be used with Controller
func ConnectToDatabase() UserDb {
	return ConnectToMongoDatabase
}

// UserDb is the user database interface
type UserDb interface {
	InsertOne(user model.User) error
	FindOne(username string) (model.User, error)
	UpdateUser(id string) error
	DeleteUser(id string) error
}

package db

import "../model"

// UserDb is the user database interface
type UserDb interface {
	CreateUser(user model.User) error
	ReadUser(username string) error
	UpdateUser(id string) error
	DeleteUser(id string) error
}

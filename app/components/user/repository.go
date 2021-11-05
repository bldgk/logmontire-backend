package user

import (
	"logmontire/common/errors"
	"logmontire/storage"
	"logmontire/storage/mongo"
)

var (
	collection = "users"
)

func New(args map[string]interface{}) *User {
	user := &User{}
	user.New(args["username"].(string))
	return user
}

func Save(user *User) ([]User, *errors.Error) {
	var users []User

	if err := storage.Save(collection, user, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func Find(selector *mongo.Selector) ([]User, *errors.Error) {
	var users []User

	if err := storage.Find(collection, selector, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func FindById(id string) (*User, *errors.Error) {
	var users []User

	if err := storage.FindById(collection, id, &users); err != nil {
		return nil, err
	}

	return &users[0], nil
}

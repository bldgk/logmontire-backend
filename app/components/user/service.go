package user

import (
	"logmontire/common/errors"
	"logmontire/storage/mongo"
)

func registerUser(args map[string]interface{}) (interface{}, *errors.Error) {
	return Save(New(args))
}

func findUser(args map[string]interface{}) (interface{}, *errors.Error) {
	return Find(mongo.NewSelector(args))
}

func allUsers(args map[string]interface{}) (interface{}, *errors.Error) {
	return Find(mongo.NewEmptySelector())
}

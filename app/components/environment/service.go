package environment

import (
	"logmontire/common/errors"
	"logmontire/storage/mongo"
)

var (
	NewEnvironment = make(chan *Environment)
)

func createEnvironment(args map[string]interface{}) (interface{}, *errors.Error) {
	env := New(args)
	envs, err := Save(env)
	if err != nil {
		NewEnvironment <- env
	}
	return envs, err
}

func findEnvironment(args map[string]interface{}) (interface{}, *errors.Error) {
	return Find(mongo.NewSelector(args))
}

func allEnvironments(_ map[string]interface{}) (interface{}, *errors.Error) {
	return Find(mongo.NewEmptySelector())
}

package environment

import (
	"logmontire/common/converting"
	"logmontire/common/errors"
	"logmontire/storage"
	"logmontire/storage/mongo"
)

var (
	collection = "environments"
)

func New(args map[string]interface{}) *Environment {
	e := &Environment{}
	e.New(args["name"].(string), converting.Itosa(args["users"]))
	return e
}

func Save(environment *Environment) ([]Environment, *errors.Error) {
	var environments []Environment

	if err := storage.Save(collection, environment, &environments); err != nil {
		return nil, err
	}

	return environments, nil
}

func Find(selector *mongo.Selector) ([]Environment, *errors.Error) {
	var environments []Environment

	if err := storage.Find(collection, selector, &environments); err != nil {
		return nil, err
	}

	return environments, nil
}

func FindById(id string) (*Environment, *errors.Error) {
	var environments []Environment

	if err := storage.FindById(collection, id, &environments); err != nil {
		return nil, err
	}

	return &environments[0], nil
}

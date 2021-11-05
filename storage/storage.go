package storage

import (
	"logmontire/common/errors"
	"logmontire/settings"
	"logmontire/storage/mongo"
)

func Init(db *settings.MongoDBConfiguration) {
	mongo.Init(db)
}

func Save(collectionName string, able Savable, result interface{}) (err *errors.Error) {
	session, collection := mongo.GetCollection(collectionName)
	defer session.Close()

	if err := collection.Insert(able.GetInstance()); err != nil {
		return mongo.MongoDBError(err)
	}

	if err := collection.FindId(able.GetId()).All(result); err != nil {
		return mongo.MongoDBError(err)
	}

	return nil
}

func Find(collectionName string, selector *mongo.Selector, objects interface{}) (err *errors.Error) {
	session, collection := mongo.GetCollection(collectionName)
	defer session.Close()

	if err := collection.Find(selector.GetSelect()).All(objects); err != nil {
		return mongo.MongoDBError(err)
	}

	return nil
}

func FindById(collectionName string, id string, objects interface{}) *errors.Error {
	session, collection := mongo.GetCollection(collectionName)
	defer session.Close()

	if err := collection.FindId(mongo.NewSelectorById(id).GetId()).All(objects); err != nil {
		return mongo.MongoDBError(err)
	}

	return nil
}

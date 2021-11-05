package log

import (
	"logmontire/common/errors"
	"logmontire/storage"
	"logmontire/storage/mongo"
)

var (
//collection = "channels"
)

func New(args map[string]interface{}) *Log {
	log := &Log{}
	log.New(
		args["title"].(string),
		args["username"].(string),
		args["body"].(string),
		args["channel"].(string))
	//args["channel_id"].(string))
	return log
}

func Save(log *Log) ([]Log, *errors.Error) {
	var logs []Log

	if err := storage.Save(log.Channel, log, &logs); err != nil {
		return nil, err
	}

	return logs, nil
}

func Find(selector *mongo.Selector) ([]Log, *errors.Error) {
	var logs []Log

	if err := storage.Find(selector.GetSelect()["channel"].(string), selector, &logs); err != nil {
		return nil, err
	}

	return logs, nil
}

//
//func FindById(id string) (*Channel, *errors.Error) {
//	var channels []Channel
//
//	if err := storage.FindById(collection, id, &channels); err != nil {
//		return nil, err
//	}
//
//	return &channels[0], nil
//}

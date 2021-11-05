package log

import (
	"logmontire/common/errors"
	"logmontire/storage/mongo"
)

var (
	NewLog = make(chan *Log)
)

func storeLog(args map[string]interface{}) (interface{}, *errors.Error) {
	log := New(args)
	NewLog <- log
	return Save(log)
}

func FindLog(args map[string]interface{}) ([]Log, *errors.Error) {
	return Find(mongo.NewSelector(args))
}

//
//func allLogs(_ map[string]interface{}) (interface{}, *errors.Error) {
//	return Find(mongo.NewEmptySelector())
//}

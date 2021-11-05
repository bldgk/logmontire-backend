package logging

import (
	"logmontire/common/errors"
	"logmontire/crypto/encoding/json"
	"logmontire/log"
)

type ExternalLogger struct {
	level log.LogLevel
}

func (el *ExternalLogger) Log(l log.Loggable) *errors.Error {
	req := make(map[string]interface{}, 4)
	req["channel"] = "logmontire"
	req["username"] = "api"
	req["title"] = ""
	req["body"] = string(json.FromInterface(l.GetLogData()))
	//http.Post("http://95.216.169.56:3001/save", nil, req)
	return nil
}

func (el *ExternalLogger) Success(l log.Loggable) *errors.Error {
	req := make(map[string]interface{}, 4)
	req["channel"] = "logmontire"
	req["username"] = "api"
	req["title"] = ""
	req["status"] = "success"
	req["body"] = string(json.FromInterface(l.GetLogData()))
	//http.Post("http://95.216.169.56:3001/save", nil, req)
	return nil
}

func (el *ExternalLogger) Fatal(l log.Loggable) *errors.Error {
	req := make(map[string]interface{}, 4)
	req["channel"] = "logmontire"
	req["username"] = "api"
	req["title"] = ""
	req["status"] = "error"
	req["body"] = string(json.FromInterface(l.GetLogData()))
	//http.Post("http://95.216.169.56:3001/save", nil, req)
	return nil
}

func (el *ExternalLogger) GetLogLevel() log.LogLevel {
	return el.level
}

func RegisterExternalLogger() {
	log.Register(&ExternalLogger{log.ExternalLog})
}

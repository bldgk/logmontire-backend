package log

import "logmontire/common/errors"

func Register(logger Logger) {
	loggers[logger.GetLogLevel()] = append(loggers[logger.GetLogLevel()], logger)
}

type Logger interface {
	GetLogLevel() LogLevel
	Log(Loggable) *errors.Error
	Success(Loggable) *errors.Error
	Fatal(Loggable) *errors.Error
}

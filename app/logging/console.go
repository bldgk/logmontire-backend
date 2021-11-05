package logging

import (
	"fmt"
	"logmontire/common/errors"
	"logmontire/log"
	"time"
)

type ConsoleLogger struct {
	level log.LogLevel
}

func (cl *ConsoleLogger) Log(l log.Loggable) *errors.Error {
	fmt.Printf("%s\t"+l.GetLogFormat()+"\n", time.Now().Format("2006-01-02 15:04:05.000000"), l.GetLogData())
	return nil
}

func (cl *ConsoleLogger) Success(l log.Loggable) *errors.Error {
	fmt.Printf("%s\t"+l.GetLogFormat()+"\n", time.Now().Format("2006-01-02 15:04:05.000000"), l.GetLogData())
	return nil
}

func (cl *ConsoleLogger) Fatal(l log.Loggable) *errors.Error {
	fmt.Printf("%s\t"+l.GetLogFormat()+"\n", time.Now().Format("2006-01-02 15:04:05.000000"), l.GetLogData())
	return nil
}

func (cl *ConsoleLogger) GetLogLevel() log.LogLevel {
	return cl.level
}

func RegisterConsoleLogger() {
	log.Register(&ConsoleLogger{log.ConsoleLog})
}

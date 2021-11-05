package log

import (
	"fmt"
)

type LogLevel int

const (
	ConsoleLog  LogLevel = 0
	ExternalLog LogLevel = 1
	AllLogs     LogLevel = 2
)

var (
	loggers = make(map[LogLevel][]Logger, 0)
)

type Loggable interface {
	GetLogData() string
	GetLogFormat() string
}

type LogString struct {
	content string
	format  string
}

func (lg *LogString) New(data interface{}) {
	lg.content = fmt.Sprint(data)
	lg.format = "%s"
}

func (lg *LogString) GetLogData() string {
	return lg.content
}

func (lg *LogString) GetLogFormat() string {
	return lg.format
}

/**
Create Log string
*/
func LStr(format string, a ...interface{}) *LogString {
	l := &LogString{}
	l.New(fmt.Sprintf(format, a))
	return l
}

func Log(l Loggable, level LogLevel) {
	if level == AllLogs {
		for _, leveledLogs := range loggers {
			for _, logger := range leveledLogs {
				if err := logger.Log(l); err != nil {
					fmt.Println(err)
				}
			}
		}
	} else {
		for _, logger := range loggers[level] {
			//go func() {
			if err := logger.Log(l); err != nil {
				fmt.Println(err)
			}
			//}()
		}
	}
}

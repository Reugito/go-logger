package logger

import (
	"os"
	"sync"
)

type Logger struct {
	file       *os.File
	mutex      sync.Mutex
	logLevel   LogLevel
	timeFormat string
}

var loggerInstance *Logger
var initMutex sync.Mutex

type DataStore interface {
	writeLog(level LogLevel, message string, data interface{})

	Debug(message string, data interface{})
	Info(message string, data interface{})
	Warning(message string, data interface{})
	Error(message string, error error)

	SetLogLevel(level LogLevel)
	SetTimeFormat(format string)
}

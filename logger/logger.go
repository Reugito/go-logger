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
	writeLog(level LogLevel, message string)

	Debug(message string)
	Info(message string)
	Warning(message string)
	Error(message string)

	SetLogLevel(level LogLevel)
	SetTimeFormat(format string)
}

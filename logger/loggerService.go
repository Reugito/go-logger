package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func NewLogger() *Logger {
	initMutex.Lock()
	defer initMutex.Unlock()

	if loggerInstance == nil {
		loggerInstance = &Logger{
			mutex:      sync.Mutex{},
			logLevel:   INFO,
			timeFormat: "2006-01-02 15:04:05",
		}
	}

	return loggerInstance
}

func (l *Logger) writeLog(level LogLevel, message string, data interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if level < l.logLevel {
		return
	}

	logTime := time.Now().Format(l.timeFormat)
	logEntry := fmt.Sprintf("[%s] %s %s %v\n", logLevelToString(level), logTime, message, data)

	log.Print(logEntry)

	if l.file != nil {
		l.file.WriteString(logEntry) // Write log entry to file
	}
}

func (l *Logger) Debug(message string, data interface{}) {
	l.writeLog(DEBUG, message, data)
}

func (l *Logger) Info(message string, data interface{}) {
	l.writeLog(INFO, message, data)
}

func (l *Logger) Warning(message string, data interface{}) {
	l.writeLog(WARNING, message, data)
}

func (l *Logger) Error(message string, err error) {
	l.writeLog(ERROR, message, err)
}

func (l *Logger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l *Logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

func logLevelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func SetLogFile(filePath string) error {
	initMutex.Lock()
	defer initMutex.Unlock()

	if loggerInstance != nil {
		return fmt.Errorf("logger: log file already set")
	}

	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
			return err
		}

		loggerInstance = &Logger{
			file:       file,
			mutex:      sync.Mutex{},
			logLevel:   INFO,
			timeFormat: time.RFC3339,
		}
	} else {
		loggerInstance = &Logger{
			mutex:      sync.Mutex{},
			logLevel:   INFO,
			timeFormat: time.RFC3339,
		}
	}

	return nil
}

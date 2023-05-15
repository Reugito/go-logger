package logger

import (
	"log"
	"os"
	"time"
)

// Logger represents a logger object
type Logger struct{}

// NewLogger creates a new logger object
func NewLogger() *Logger {
	return &Logger{}
}

// LogInfo logs an info message
func (l *Logger) LogInfo(message string, data interface{}) {
	log.Printf("[INFO] %s: %s %v\n", time.Now().Format(time.RFC3339), message, data)
}

// LogWarning logs a warning message
func (l *Logger) LogWarning(message string, data interface{}) {
	log.Printf("[WARNING] %s: %s %v\n", time.Now().Format(time.RFC3339), message, data)
}

// LogError logs an error message
func (l *Logger) LogError(message string, error error) {
	log.Printf("[ERROR] %s: %s %v\n", time.Now().Format(time.RFC3339), message, error)
}

// SetLogFile sets the log file path
func SetLogFile(logFilePath string) {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	log.SetOutput(logFile)
}

//// SetLogLevel sets the log level
//func SetLogLevel(level LogLevel) {
//	switch level {
//	case InfoLevel:
//		log.SetFlags(log.Ldate | log.Ltime)
//	case DebugLevel, ErrorLevel, WarningLevel:
//		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
//	}
//}

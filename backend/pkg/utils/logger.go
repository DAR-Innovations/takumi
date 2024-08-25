package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

type CustomLogger struct {
	logger *log.Logger
}

func NewLogger() *CustomLogger {
	return &CustomLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *CustomLogger) LogError(message string, err error, details map[string]interface{}) {
	logMsg := fmt.Sprintf("%s | ERROR: %s | DETAILS: %v | %s", time.Now().Format(time.RFC3339), message, details, err)
	l.logger.Println(logMsg)
}

func (l *CustomLogger) LogInfo(message string, details map[string]interface{}) {
	logMsg := fmt.Sprintf("%s | INFO: %s | DETAILS: %v", time.Now().Format(time.RFC3339), message, details)
	l.logger.Println(logMsg)
}

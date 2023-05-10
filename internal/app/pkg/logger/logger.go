package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the global logger instance
var Logger *logrus.Logger

func init() {
	// set the global logger instance
	Logger = logrus.New()

	// set logger formatter
	Logger.Formatter = &logrus.JSONFormatter{}

	// set logger output to stdout
	Logger.Out = os.Stdout

	// set logger level
	level, err := logrus.ParseLevel(getLogLevel())
	if err != nil {
		// fallback to default level (info) on error
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)
}

// getLogLevel returns the log level from environment variable or default to "info"
func getLogLevel() string {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	return logLevel
}

// Info logs a message with info level
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Warn logs a message with warning level
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Error logs a message with error level
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Fatal logs a message with fatal level
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// WithFields adds fields to the logger instance
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Logger.WithFields(fields)
}

// SetLogLevel sets the log level
func SetLogLevel(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("invalid log level %s", logLevel)
	}
	Logger.SetLevel(level)
	return nil
}

//logrus sebagai library logging. Log level dapat diatur melalui
//environment variable dengan variabel LOG_LEVEL. Log level yang
//tersedia adalah debug, info, warning, error, dan fatal. Default log level yang digunakan adalah info.

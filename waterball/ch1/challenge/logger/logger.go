package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	ERROR
)

type Logger struct {
	level  LogLevel
	writer io.Writer
}

// NewLogger creates a new Logger with the specified log level
func NewLogger(level string) *Logger {
	logger := &Logger{
		writer: os.Stdout, // Default output to stdout
	}

	switch strings.ToUpper(level) {
	case "DEBUG":
		logger.level = DEBUG
	case "INFO":
		logger.level = INFO
	case "ERROR":
		logger.level = ERROR
	default:
		logger.level = INFO
	}

	return logger
}

// SetOutput sets the log output destination
func (l *Logger) SetOutput(w io.Writer) {
	l.writer = w
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level <= DEBUG {
		l.log("DEBUG", format, args...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	if l.level <= INFO {
		l.log("INFO", format, args...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	if l.level <= ERROR {
		l.log("ERROR", format, args...)
	}
}

// log formats and writes the log message with file and line number
func (l *Logger) log(level, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006/01/02 15:04:05")

	// Get caller file and line number, skipping runtime-internal files
	file, line := getCaller()
	fmt.Fprintf(l.writer, "%s [%s] %s:%d %s\n", timestamp, level, file, line, message)
}

// getCaller retrieves the file name and line number of the caller, skipping runtime-internal files
func getCaller() (string, int) {
	for i := 3; i < 10; i++ { // Start at 3 and try up to a reasonable depth
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			return "unknown", 0
		}
		// Skip runtime-internal files (e.g., runtime/proc.go)
		if !strings.Contains(file, "runtime/") {
			return filepath.Base(file), line
		}
	}
	return "unknown", 0
}

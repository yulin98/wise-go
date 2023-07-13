package log

import (
	"fmt"
	"log"
	"os"
)

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if ok {
		SetLevel(lvl)
	}
}

// Level specifies the severity of a given log message
type Level int

// Log levels
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelPlain
)

var (
	logger   = log.New(os.Stdout, "[Gateway] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	logLevel = LevelDebug
)

// String returns the string form for a given LogLevel
func (lvl Level) String() string {
	switch lvl {
	case LevelInfo:
		return "[INFO] "
	case LevelWarning:
		return "[WARN] "
	case LevelError:
		return "[ERROR] "
	case LevelFatal:
		return "[FATAL] "
	case LevelDebug:
		return "[DEBUG] "
	}
	return ""
}

// IsEnabled 判断当前日志是否支持指定级别输出
func IsEnabled(lvl Level) bool {
	return lvl >= logLevel
}

func l(lvl Level, line string, args ...interface{}) {
	if IsEnabled(lvl) {
		message := fmt.Sprintf(line, args...)
		_ = logger.Output(3, fmt.Sprintf("%s%s\n", lvl, message))
	}
}

// Plain Log plain
func Plain(line string, args ...interface{}) {
	l(LevelPlain, line, args...)
}

// Debug Log debug
func Debug(line string, args ...interface{}) {
	l(LevelDebug, line, args...)
}

// Info Log info
func Info(line string, args ...interface{}) {
	l(LevelInfo, line, args...)
}

// Warn Log warn
func Warn(line string, args ...interface{}) {
	l(LevelWarning, line, args...)
}

// Error Log error
func Error(line string, args ...interface{}) {
	l(LevelError, line, args...)
}

// Fatal Log fatal
func Fatal(line string, args ...interface{}) {
	l(LevelFatal, line, args...)
	os.Exit(1)
}

// SetLevel 设置日志等级
func SetLevel(level string) {
	switch level {
	case "fatal":
		logLevel = LevelFatal
	case "error":
		logLevel = LevelError
	case "warning":
		logLevel = LevelWarning
	case "info":
		logLevel = LevelInfo
	case "debug":
		logLevel = LevelDebug
	case "plain":
		logLevel = LevelPlain
	default:
		logLevel = LevelInfo
	}
}

// GetLogger 获取 logger，以便其他组件直接使用
func GetLogger() *log.Logger {
	return logger
}

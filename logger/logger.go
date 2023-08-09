package logger

import (
    "fmt"
    "log"
    "os"
)

const (
    Reset  = "\033[0m"
    Black  = "\033[30m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
    White  = "\033[37m"

    Bold       = "\033[1m"
    Underline  = "\033[4m"
    Background = "\033[7m"
)

type LogLevel int

const (
    LogError LogLevel = iota
    LogInfo
    LogWarning
    LogDebug
)

var LogLevelNames = []string{"ERROR", "INFO", "WARNING", "DEBUG"}

var logLevel LogLevel

func init() {
    setLogLevelFromEnv()
}

func setLogLevel(level LogLevel) {
    logLevel = level
}

func setLogLevelFromEnv() {
    logLevelStr := os.Getenv("LOG_LEVEL")
    if logLevelStr == "" {
        setLogLevel(LogInfo)
        return
    }
    switch logLevelStr {
    case "ERROR":
        setLogLevel(LogError)
    case "INFO":
        setLogLevel(LogInfo)
    case "WARNING":
        setLogLevel(LogWarning)
    case "DEBUG":
        setLogLevel(LogDebug)
    default:
        setLogLevel(LogInfo)
    }
}

func logf(color string, level LogLevel, format string, v ...interface{}) {
    if level <= logLevel {
        message := fmt.Sprintf(format, v...)
        log.Printf("[%s%s%s] %s%s%s\n", color, LogLevelNames[level], Reset, color, message, Reset)
    }
}

// Error 打印错误日志
func Error(v ...interface{}) {
    logf(Red, LogError, "%v", v...)
}

// Errorf 格式化打印错误日志
func Errorf(format string, v ...interface{}) {
    logf(Red, LogError, format, v...)
}

// Info 打印信息日志
func Info(v ...interface{}) {
    logf(Green, LogInfo, "%v", v...)
}

// Infof 格式化打印信息日志
func Infof(format string, v ...interface{}) {
    logf(Green, LogInfo, format, v...)
}

// Warning 打印警告日志
func Warning(v ...interface{}) {
    logf(Yellow, LogWarning, "%v", v...)
}

// Warningf 格式化打印警告日志
func Warningf(format string, v ...interface{}) {
    logf(Yellow, LogWarning, format, v...)
}

// Debug 打印调试日志
func Debug(v ...interface{}) {
    logf(Cyan, LogDebug, "%v", v...)
}

// Debugf 格式化打印调试日志
func Debugf(format string, v ...interface{}) {
    logf(Cyan, LogDebug, format, v...)
}

package logs

import (
	"fmt"
	"os"
	"path"
)

// 提供了Logs的默认实现

const (
	maxLogSize = 1000 * 1000 * 1000
)

var (
	appLogger *Logger
)

// LogSegment ....
func LogSegment(logInterval string) SegDuration {
	if logInterval == "hour" {
		return HourDur
	} else if logInterval == "day" {
		return DayDur
	} else {
		return NoDur
	}
}

// LogLevelByName ...
func LogLevelByName(level string) int {
	m := map[string]int{
		"trace":  LevelTrace,
		"debug":  LevelDebug,
		"info":   LevelInfo,
		"notice": LevelNotice,
		"warn":   LevelWarn,
		"error":  LevelError,
		"fatal":  LevelFatal,
	}
	return m[level]
}

// InitAppLog 初始化app log
func InitAppLog(level, interval, logPath string) {

	logLevel := LogLevelByName(level)
	appLogger = NewLogger(10)
	appLogger.SetLevel(logLevel)
	appLogger.SetCallDepth(3)

	dir := path.Dir(logPath)
	os.MkdirAll(dir, os.ModeDir)

	fileProvider := NewFileProvider(logPath, LogSegment(interval), maxLogSize)
	fileProvider.SetLevel(logLevel)
	if err := appLogger.AddProvider(fileProvider); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to add fileProvider: %s\n", err)
	}

	appLogger.StartLogger()

	InitLogger(appLogger)
}

// CloseAppLog 停止applog
func CloseAppLog() {
	if appLogger != nil {
		appLogger.Flush()
		appLogger.Stop()
	}
}

package logging

import (
	"path/filepath"
	"runtime"
	"strings"
	
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	l *zap.Logger
}

type Level = zapcore.Level

const (
	DebugLevel  Level = zap.DebugLevel  // -1, default level
	InfoLevel   Level = zap.InfoLevel   // 0,
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log  PanicLevel logs a message, then panics
	PanicLevel  Level = zap.PanicLevel  // 4
	FatalLevel  Level = zap.FatalLevel  // 5  FatalLevel logs a message, then calls os.Exit(1).
)

func GetRootLogger() *Logger {
	return rootLogger.l
}

func GetLogger(name ...string) *Logger {
	var loggerName string
	// Return name as 'package/file'
	if len(name) == 0 {
		_, f, _, _ := runtime.Caller(1)
		s := strings.LastIndex(filepath.Dir(f), "/")
		loggerName = f[s+1:]
	}
	return rootLogger.getLogger(loggerName)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}
func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

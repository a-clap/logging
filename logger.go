package logging

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
	conf   *config
}

type Level = zapcore.Level
type HandlerName = string

const (
	DebugLevel  Level = zap.DebugLevel  // -1, default level
	InfoLevel   Level = zap.InfoLevel   // 0,
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log  PanicLevel logs a message, then panics
	PanicLevel  Level = zap.PanicLevel  // 4
	FatalLevel  Level = zap.FatalLevel  // 5  FatalLevel logs a message, then calls os.Exit(1).
)

const (
	Console    HandlerName = "_console"
	File       HandlerName = "_file"
	RotateFile HandlerName = "_rotate_file"
)

func GetRootLogger() *Logger {
	return rootLogger.l
}

func GetLogger(name ...string) *Logger {
	var loggerName string
	// Return name as 'package/file'
	if len(name) == 0 || name[0] == "" {
		_, f, _, _ := runtime.Caller(1)
		f = strings.TrimSuffix(f, ".go")
		s := strings.LastIndex(filepath.Dir(f), "/")
		loggerName = f[s+1:]
	} else {
		loggerName = name[0]
	}
	return rootLogger.getLogger(loggerName)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...Field) {
	l.logger.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.logger.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}

func (l *Logger) SetLevel(level Level) {
	for _, c := range l.conf.cores {
		c.level.SetLevel(level)
	}
}

func (l *Logger) SetHandlerLevel(h HandlerName, level Level) error {
	core, ok := l.conf.cores[h]
	if !ok {
		return fmt.Errorf("no such handler: %v", h)
	}
	core.level.SetLevel(level)
	return nil
}

func (l *Logger) AddConsoleHandler(e EncoderConfig, level Level) {
	l.conf.addConsoleHandler(e, level)
	l.logger = l.conf.build()

	for _, child := range rootLogger.getChilds(l.conf.name) {
		child.conf.addConsoleHandler(e, level)
		child.logger = child.conf.build()
	}
}

func (l *Logger) AddFileHandler(e EncoderConfig, level Level, w io.Writer) {
	l.conf.addFileHandler(w, e, level)
	l.logger = l.conf.build()

	for _, child := range rootLogger.getChilds(l.conf.name) {
		child.conf.addFileHandler(w, e, level)
		child.logger = child.conf.build()
	}
}

package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type root struct {
	l       *Logger
	loggers map[string]*Logger
}

const (
	initLevel = DebugLevel
)

var (
	rootLogger *root
)

func newDefaultLogger() *Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(initLevel)
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	log, _ := cfg.Build(zap.AddCallerSkip(1))
	
	return &Logger{l: log}
}

func init() {
	rootLogger = &root{
		l:       newDefaultLogger(),
		loggers: make(map[string]*Logger, 0),
	}
}

func (r *root) getLogger(name string) *Logger {
	if _, ok := r.loggers[name]; !ok {
		r.loggers[name] = newDefaultLogger()
	}
	return r.loggers[name]
}

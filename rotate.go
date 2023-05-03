package logging

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type RotateOptions struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
}

func (l *Logger) AddRotateFileHandler(e EncoderConfig, level Level, filename string, opts RotateOptions) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    opts.MaxSize,
		MaxAge:     opts.MaxAge,
		MaxBackups: opts.MaxBackups,
		LocalTime:  opts.LocalTime,
		Compress:   opts.Compress,
	})

	l.conf.addRotateFileHandler(w, e, level)
	l.logger = l.conf.build()

	for _, child := range rootLogger.getChilds(l.conf.name) {
		child.conf.addRotateFileHandler(w, e, level)
		child.logger = child.conf.build()
	}
}

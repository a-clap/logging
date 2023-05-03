package logging

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type core struct {
	enc    zapcore.Encoder
	writer io.Writer
	level  zap.AtomicLevel
}

type config struct {
	name  string
	cores map[string]core
}

func newConfig(name string) *config {
	c := &config{
		name:  name,
		cores: make(map[string]core, 0),
	}
	return c
}

func inheritConfig(name string, conf *config) *config {
	cores := make(map[string]core, len(conf.cores))
	for k, v := range conf.cores {
		cores[k] = v
	}

	c := &config{
		name:  name,
		cores: cores,
	}
	return c
}

func (c *config) addHandler(name string, encoder zapcore.Encoder, w io.Writer, level Level) {
	c.cores[name] = core{enc: encoder, writer: w, level: zap.NewAtomicLevel()}
	c.cores[name].level.SetLevel(level)
}

func (c *config) addConsoleHandler(e EncoderConfig, level Level) {
	c.addHandler("console", zapcore.NewConsoleEncoder(e), os.Stdout, level)
}

func (c *config) addFileHandler(w io.Writer, e EncoderConfig, level Level) {
	c.addHandler("file", zapcore.NewJSONEncoder(e), w, level)
}

func (c *config) build() *zap.Logger {
	cores := make([]zapcore.Core, 0, len(c.cores))
	for _, c := range c.cores {
		cores = append(cores, zapcore.NewCore(c.enc, zapcore.AddSync(c.writer), c.level))
	}
	tee := zapcore.NewTee(cores...)
	return zap.New(tee).WithOptions(zap.AddCallerSkip(1), zap.AddCaller())

}

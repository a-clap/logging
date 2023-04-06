package logging

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type root struct {
	l       *Logger
	loggers map[string]*Logger
}

var (
	rootLogger *root
)

func init() {
	rootLogger = &root{
		l:       newDefaultLogger(""),
		loggers: make(map[string]*Logger, 0),
	}

}

func newDefaultLogger(name string) *Logger {
	e := zap.NewDevelopmentEncoderConfig()
	e.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	e.EncodeLevel = zapcore.CapitalColorLevelEncoder

	c := newConfig(name)
	c.addConsoleHandler(e, DebugLevel)

	return &Logger{logger: c.build(), conf: c}
}

// createLogger creates logger and his parents
func (r *root) createLogger(name string) {
	// Split name for parent name and child name
	var p *Logger
	idx := strings.LastIndex(name, "/")
	// Do we have parent?
	if idx == -1 {
		// No, we don't, so just inherit from rootlogger
		p = GetRootLogger()
	} else {
		// Get parent logger
		parent := name[0:idx]
		// It recursively creates parent and his parents
		p = r.getLogger(parent)
	}

	// copy config from parent
	conf := inheritConfig(name, p.conf)
	ch := &Logger{
		logger: conf.build(),
		conf:   conf,
	}
	r.loggers[name] = ch
}

func (r *root) getLogger(name string) *Logger {
	if l, ok := r.loggers[name]; ok {
		// Found logger, just return
		return l
	}
	// We need to create logger
	r.createLogger(name)
	return r.getLogger(name)
}

func (r *root) getChilds(name string) []*Logger {
	var childs []*Logger
	for k, v := range r.loggers {
		if strings.HasPrefix(k, name) {
			childs = append(childs, v)
		}
	}
	return childs
}

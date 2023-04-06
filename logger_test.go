package logging

import (
	"testing"
	
	"github.com/stretchr/testify/suite"
)

type LoggingSuite struct {
	suite.Suite
}

func TestLogging(t *testing.T) {
	suite.Run(t, new(LoggingSuite))
}

func (l *LoggingSuite) TestGetChildLoggers() {
	t := l.Require()
	
	args := []struct {
		createLoggers  []string
		child          string
		expectedChilds []string
	}{
		{
			createLoggers:  []string{"1/2"},
			child:          "1",
			expectedChilds: []string{"1", "1/2"},
		},
		{
			createLoggers:  []string{"1/2", "2/3"},
			child:          "1",
			expectedChilds: []string{"1", "1/2"},
		},
	}
	for _, arg := range args {
		// Clear loggers
		rootLogger.loggers = make(map[string]*Logger, 0)
		// Create new
		for _, c := range arg.createLoggers {
			GetLogger(c)
		}
		
		// GetChilds
		childs := rootLogger.getChilds(arg.child)
		childNames := make([]string, len(childs))
		for i, v := range childs {
			childNames[i] = v.conf.name
		}
		l.Assert().Len(childs, len(arg.expectedChilds))
		t.ElementsMatch(childNames, arg.expectedChilds)
		
	}
	
}
func (l *LoggingSuite) TestCreateProperLoggers() {
	t := l.Require()
	
	args := []struct {
		loggerName      string
		expectedLoggers []string
	}{
		{
			loggerName:      "logger",
			expectedLoggers: []string{"logger"},
		},
		{
			loggerName:      "logger/test",
			expectedLoggers: []string{"logger", "logger/test"},
		},
		{
			loggerName:      "",
			expectedLoggers: []string{"logging", "logging/logger_test"},
		},
		{
			loggerName:      "1/2/3/4/5.go",
			expectedLoggers: []string{"1", "1/2", "1/2/3", "1/2/3/4", "1/2/3/4/5.go"},
		},
	}
	for _, arg := range args {
		// Prepare map clean
		rootLogger.loggers = make(map[string]*Logger, 0)
		GetLogger(arg.loggerName)
		
		var got []string
		var confName []string
		for k, v := range rootLogger.loggers {
			got = append(got, k)
			confName = append(confName, v.conf.name)
		}
		t.ElementsMatch(arg.expectedLoggers, got, arg.loggerName)
		t.ElementsMatch(arg.expectedLoggers, confName, arg.loggerName)
	}
}

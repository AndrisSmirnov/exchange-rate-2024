package scheduler_service

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type customLogger struct {
	logger *logrus.Logger
}

var jobNames = map[cron.EntryID]string{}
var ENTRY_ID = 3

func (c *customLogger) Info(msg string, keysAndValues ...interface{}) {
	for i, key := range keysAndValues {
		if i == ENTRY_ID {
			c.logger.Infof("%s job: %v", msg, jobNames[key.(cron.EntryID)])
		}
	}
}

// Cron log format
// now
// 2023-02-02 20:31:43.000472385 +0000 UTC
// entry
// 5
// next
// 2023-02-02 20:32:03 +0000 UTC

func (c *customLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	c.logger.Info(msg)
	c.logger.Info(keysAndValues...)
}

func newLogger() *customLogger {
	return &customLogger{
		logger: logrus.StandardLogger(),
	}
}

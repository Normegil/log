package log

import "github.com/Sirupsen/logrus"

type StructuredLog struct {
	Logger logrus.FieldLogger
}

func (l StructuredLog) Log(lvl Level, str Structure, v ...interface{}) {
	if nil != l.Logger {
		logger := l.StructureMessage(str)
		switch lvl {
		case PANIC:
			logger.Panic(v...)
		case INFO:
			logger.Info(v...)
		case DEBUG:
			logger.Debug(v...)
		}
	}
}

func (l StructuredLog) StructureMessage(str Structure) logrus.FieldLogger {
	logger := l.Logger
	for key, value := range str {
		logger = logger.WithField(key, value)
	}
	return logger
}

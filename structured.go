package log

import "github.com/Sirupsen/logrus"

type StructuredLog struct {
	Logger logrus.FieldLogger
}

func (l StructuredLog) Log(lvl Level, str Structure, v ...interface{}) {
	if nil != l.Logger {
		fields := logrus.Fields{}
		for key, value := range str {
			fields[key] = value
		}

		logger := l.Logger.WithFields(fields)
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

func (l StructuredLog) With(str Structure) AgnosticLogger {
	fields := logrus.Fields{}
	for key, value := range str {
		fields[key] = value
	}

	l.Logger = l.Logger.WithFields(fields)
	return l
}

package log

import "github.com/Sirupsen/logrus"

type StructuredLog struct {
	Logger logrus.FieldLogger
}

func (l StructuredLog) Panic(str Structure, v ...interface{}) {
	if nil != l.Logger {
		l.StructureMessage(str).Panic(v...)
	}
}

func (l StructuredLog) Info(str Structure, v ...interface{}) {
	if nil != l.Logger {
		l.StructureMessage(str).Info(v...)
	}
}

func (l StructuredLog) Debug(str Structure, v ...interface{}) {
	if nil != l.Logger {
		l.StructureMessage(str).Debug(v...)
	}
}

func (l StructuredLog) StructureMessage(str Structure) logrus.FieldLogger {
	logger := l.Logger
	for key, value := range str {
		logger = logger.WithField(key, value)
	}
	return logger
}

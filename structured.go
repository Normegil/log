package log

import "github.com/Sirupsen/logrus"

// StructuredLog support decorate Logrus to implement AgnosticLogger
type StructuredLog struct {
	Logger logrus.FieldLogger
}

// Log log a message to the output defined in logrus.
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
		case FATAL:
			logger.Fatal(v...)
		case ERROR:
			logger.Error(v...)
		case WARN:
			logger.Warn(v...)
		case INFO:
			logger.Info(v...)
		case DEBUG:
			logger.Debug(v...)
		case TRACE:
			logger.WithField("level", "trace").Print(v...)
		}
	}
}

// With send back a logger containing the fields in the given structure
func (l StructuredLog) With(str Structure) AgnosticLogger {
	fields := logrus.Fields{}
	for key, value := range str {
		fields[key] = value
	}

	l.Logger = l.Logger.WithFields(fields)
	return l
}

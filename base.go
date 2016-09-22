package log

import (
	"log"
	"strings"
)

// BasicLog decorate the go logger.
type BasicLog struct {
	Logger    *log.Logger
	Level     Level
	structure Structure
}

// Log log your message on the specified level, with a structure holding the fields you want to log and ending with the message
func (l BasicLog) Log(lvl Level, str Structure, v ...interface{}) {
	if nil != l.Logger && lvl >= l.Level {
		l.structure = l.structure.With(str)
		switch lvl {
		case PANIC:
			l.Logger.Panic(l.toString(l.structure, lvl, v...)...)
		case FATAL:
			l.Logger.Fatal(l.toString(l.structure, lvl, v...)...)
		case ERROR, WARN, INFO, DEBUG, TRACE:
			l.Logger.Print(l.toString(l.structure, lvl, v...)...)
		}
	}
}

// With add some fields to a new logger created from the source and return it
func (l BasicLog) With(str Structure) AgnosticLogger {
	l.structure = l.structure.With(str)
	return l
}

func (l BasicLog) toString(str Structure, lvl Level, v ...interface{}) []interface{} {
	toLog := []interface{}{strings.Join([]string{"[", strings.ToUpper(lvl.String()), "]"}, "")}
	toLog = append(toLog, v...)
	if 0 != len(str) {
		toLog = append(toLog, " "+str.String())
	}
	return toLog
}

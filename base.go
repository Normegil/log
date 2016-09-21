package log

import (
	"log"
)

type BasicLog struct {
	Logger     *log.Logger
	PrintDebug bool
}

func (l BasicLog) Panic(str Structure, v ...interface{}) {
	if nil != l.Logger {
		l.Logger.Panic(l.toString(str, "[PANIC]", v...)...)
	}
}

func (l BasicLog) Info(str Structure, v ...interface{}) {
	if nil != l.Logger {
		l.Logger.Print(l.toString(str, "[INFO]", v...)...)
	}
}

func (l BasicLog) Debug(str Structure, v ...interface{}) {
	if nil != l.Logger && l.PrintDebug {
		l.Logger.Print(l.toString(str, "[DEBUG]", v...)...)
	}
}

func (l BasicLog) toString(str Structure, prefix string, v ...interface{}) []interface{} {
	toLog := []interface{}{
		prefix,
	}
	toLog = append(toLog, v...)
	if 0 != len(str) {
		toLog = append(toLog, " "+str.String())
	}
	return toLog
}

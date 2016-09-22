package log

import (
	"bytes"
	"log"
)

type BasicLog struct {
	Logger     *log.Logger
	PrintDebug bool
}

func (l BasicLog) Log(lvl Level, str Structure, v ...interface{}) {
	if nil != l.Logger {
		switch lvl {
		case PANIC:
			l.Logger.Panic(l.toString(str, lvl, v...)...)
		case INFO:
			l.Logger.Print(l.toString(str, lvl, v...)...)
		case DEBUG:
			if l.PrintDebug {
				l.Logger.Print(l.toString(str, lvl, v...)...)
			}
		}
	}
}

func (l BasicLog) toString(str Structure, lvl Level, v ...interface{}) []interface{} {
	lvlToLog := bytes.Buffer{}
	lvlToLog.WriteString("[")
	lvlToLog.WriteString(string(lvl))
	lvlToLog.WriteString("]")
	toLog := []interface{}{lvlToLog.String()}
	toLog = append(toLog, v...)
	if 0 != len(str) {
		toLog = append(toLog, " "+str.String())
	}
	return toLog
}

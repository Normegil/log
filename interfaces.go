package log

type AgnosticLogger interface {
	Log(lvl Level, str Structure, v ...interface{})
}

package log

type AgnosticLogger interface {
	Panic(str Structure, v ...interface{})
	Info(str Structure, v ...interface{})
	Debug(str Structure, v ...interface{})
}

type Logger interface {
	Panic(v ...interface{})
	Info(v ...interface{})
	Debug(v ...interface{})
}

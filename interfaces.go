package log

// AgnosticLogger defines the interface to use if you want to support all other loggers in your library. It define a logger that can log on different levels and support structured logging
type AgnosticLogger interface {
	Log(lvl Level, str Structure, v ...interface{})
	With(Structure) AgnosticLogger
}

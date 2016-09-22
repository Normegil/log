package log

// Level represent the level of logging
type Level int

// Levels of logging
const (
	PANIC Level = 10
	FATAL Level = 8
	ERROR Level = 7
	WARN  Level = 5
	INFO  Level = 0
	DEBUG Level = -5
	TRACE Level = -10
)

// String sends the string representation of the current level.
func (l Level) String() string {
	switch l {
	case PANIC:
		return "Panic"
	case FATAL:
		return "Fatal"
	case ERROR:
		return "Error"
	case WARN:
		return "Warn"
	case INFO:
		return "Info"
	case DEBUG:
		return "Debug"
	case TRACE:
		return "Trace"
	}
	return ""
}

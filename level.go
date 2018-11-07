package logmatic

// LogLevel controls what log statements are
// activated for a given logger
type LogLevel uint8

// Log levels
const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

package logmatic

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var printf = fmt.Printf

// logFunc represents a log function
type logFunc func(a ...interface{}) string

// Logger maintains a set of logging functions
// and has a log level that can be modified dynamically
type Logger struct {
	level LogLevel
	trace logFunc
	debug logFunc
	info  logFunc
	warn  logFunc
	error logFunc
}

func (l *Logger) now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (l *Logger) log(level string, format string, a ...interface{}) {
	printf("%s %14s %s %s\n",
		color.MagentaString(l.now()),
		level,
		color.MagentaString("=>"),
		fmt.Sprintf(format, a...))
}

// SetLevel updates the logging level for future logs
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Trace logs a trace statement
// TRACE only
func (l *Logger) Trace(format string, a ...interface{}) {
	if l.level == TRACE {
		l.log(l.trace("TRACE"), format, a...)
	}
}

// Debug logs a debug statement
// DEBUG or lower
func (l *Logger) Debug(format string, a ...interface{}) {
	if l.level <= DEBUG {
		l.log(l.debug("DEBUG"), format, a...)
	}
}

// Info logs an info statement
// INFO or lower
func (l *Logger) Info(format string, a ...interface{}) {
	if l.level <= INFO {
		l.log(l.info("INFO"), format, a...)
	}
}

// Warn logs a warn statement
// WARN or lower
func (l *Logger) Warn(format string, a ...interface{}) {
	if l.level <= WARN {
		l.log(l.warn("WARN"), format, a...)
	}
}

// Error logs an error statement
// ERROR or lower (any level)
func (l *Logger) Error(format string, a ...interface{}) {
	l.log(l.error("ERROR"), format, a...)
}

// NewLogger creates a new logger
// Default level is INFO
func NewLogger() *Logger {
	return &Logger{
		level: INFO,
		trace: color.New(color.FgBlue).SprintFunc(),
		debug: color.New(color.FgGreen).SprintFunc(),
		info:  color.New(color.FgCyan).SprintFunc(),
		warn:  color.New(color.FgYellow).SprintFunc(),
		error: color.New(color.FgRed).SprintFunc(),
	}
}

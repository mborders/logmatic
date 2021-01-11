package logmatic

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var printf = fmt.Printf
var exit = os.Exit

// default colors for various log levels
const (
	defaultTraceColor = color.FgBlue
	defaultDebugColor = color.FgGreen
	defaultInfoColor  = color.FgCyan
	defaultWarnColor  = color.FgYellow
	defaultErrorColor = color.FgRed
)

// logFunc represents a log function
type logFunc func(a ...interface{}) string

// Logger maintains a set of logging functions
// and has a log level that can be modified dynamically
type Logger struct {
	level             LogLevel
	trace             logFunc
	debug             logFunc
	info              logFunc
	warn              logFunc
	error             logFunc
	fatal             logFunc
	ExitOnFatal       bool   // If true, logger will run os.Exit(1) when calling Fatal().
	ShowTimestamp     bool   // If true, logger will show the current timestamp in its output. Default: true
	UseUnixTimestamp  bool   // If true and ShowTimestamp is true, logger will print the current timestamp as a unix timestamp instead of a UTC string. Default: false
	UnixTimestampNano bool   // If true and UseUnixTimestamp is true, logger will print the current timestamp as nanoseconds instead of seconds. Default: false
	Separator         string // Separator string to use between log level output and log message output. Default: "=>"
	ShowSeparator     bool   // Controls whether Separator is shown when logging. Default: true
	ColorizeMessages  bool   // Controls whether log messages are shown with the same color (unbolded) as their log level. Default: false
}

func (l *Logger) now() string {
	now := time.Now()

	if l.UseUnixTimestamp {
		var timestamp int64

		switch l.UnixTimestampNano {
		case true:
			timestamp = now.UnixNano()
		default:
			timestamp = now.Unix()
		}
		return fmt.Sprintf("%[1]d", timestamp)
	}

	return now.Format(defaultTimeformat)
}

// builds a format string using the logger's parameter fields
func (l *Logger) buildFormatString() string {
	var sb strings.Builder

	// add format placeholder for the current timestamp, if applicable
	if l.ShowTimestamp {
		switch l.UseUnixTimestamp {
		case true:
			sb.WriteString("%d")
		default:
			sb.WriteString("%s")
		}
	}

	// add format placeholder for log level
	sb.WriteString("%s")

	// add format placeholder for separator, if applicable
	if l.ShowSeparator {
		sb.WriteString("%s")
	}

	// add format placeholder for the actual data
	sb.WriteString("%s")

	sb.WriteString("\n")

	return sb.String()
}

func (l *Logger) buildLogArgs(level, format string, a ...interface{}) []interface{} {
	var logArgs []interface{}

	if l.ShowTimestamp {
		logArgs = append(logArgs, color.MagentaString(l.now()))
	}

	logArgs = append(logArgs, level)

	if l.ShowSeparator {
		logArgs = append(logArgs, color.MagentaString(l.Separator))
	}

	msgStr := fmt.Sprintf(format, a...)
	if l.ColorizeMessages {
		var sprintFunc logFunc

		switch level {
		case "TRACE":
			sprintFunc = l.trace
		case "DEBUG":
			sprintFunc = l.debug
		case "INFO":
			sprintFunc = l.info
		case "WARN":
			sprintFunc = l.warn
		case "ERROR":
			sprintFunc = l.error
		case "FATAL":
			sprintFunc = l.fatal
		default:
			sprintFunc = l.debug
		}

		msgStr = sprintFunc(msgStr)
	}

	logArgs = append(logArgs, msgStr)

	return logArgs
}

func (l *Logger) log(level string, format string, a ...interface{}) {
	formatString := l.buildFormatString()
	logArgs := l.buildLogArgs(level, format, a...)

	printf(formatString, logArgs...)
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
	if l.level <= ERROR {
		l.log(l.error("ERROR"), format, a...)
	}
}

// Fatal logs an error statement and exits the application
// FATAL or lower (any level)
func (l *Logger) Fatal(format string, a ...interface{}) {
	l.log(l.fatal("FATAL"), format, a...)

	if l.ExitOnFatal {
		exit(1)
	}
}

// NewLogger creates a new logger using provided params (or default values).
// Default log level is INFO
func NewLogger(params *LoggerParams) *Logger {
	if params == nil {
		params = NewLoggerParams()
	}

	return &Logger{
		level:             params.LogLevel(),
		trace:             color.New(defaultTraceColor).SprintFunc(),
		debug:             color.New(defaultDebugColor).SprintFunc(),
		info:              color.New(defaultInfoColor).SprintFunc(),
		warn:              color.New(defaultWarnColor).SprintFunc(),
		error:             color.New(defaultErrorColor).SprintFunc(),
		fatal:             color.New(defaultErrorColor, color.Bold).SprintFunc(),
		ExitOnFatal:       params.ExitOnFatal(),
		ShowTimestamp:     params.ShowTimestamp(),
		UseUnixTimestamp:  params.UseUnixTimestamp(),
		UnixTimestampNano: params.UnixTimestampNano(),
		Separator:         params.Separator(),
		ShowSeparator:     params.ShowSeparator(),
		ColorizeMessages:  params.ColorizeMessages(),
	}
}

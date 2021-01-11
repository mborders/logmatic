package logmatic

const (
	defaultSeparator         = "=>"
	defaultShowSeparator     = true
	defaultExitOnFatal       = true
	defaultShowTimestamp     = true
	defaultUseUnixTimestamp  = false
	defaultUnixTimestampNano = false
	defaultColorizeMessages  = false
	defaultLogLevel          = INFO
	defaultTimeformat        = "2006-01-02 15:04:05"
)

type LoggerParams struct {
	exitOnFatal       *bool
	showTimestamp     *bool
	useUnixTimestamp  *bool
	unixTimestampNano *bool
	separator         *string
	showSeparator     *bool
	colorizeMessages  *bool
	logLevel          *LogLevel
}

type loggerColorParams struct {
}

func NewLoggerParams() *LoggerParams {
	return &LoggerParams{}
}

func (p *LoggerParams) LogLevel() LogLevel {
	if p.logLevel != nil {
		return *p.logLevel
	}

	return defaultLogLevel
}

func (p *LoggerParams) SetLogLevel(logLevel LogLevel) *LoggerParams {
	p.logLevel = &logLevel

	return p
}

func (p *LoggerParams) ExitOnFatal() bool {
	if p.exitOnFatal != nil {
		return *p.exitOnFatal
	}

	return defaultExitOnFatal
}

func (p *LoggerParams) SetExitOnFatal(exitOnFatal bool) *LoggerParams {
	p.exitOnFatal = &exitOnFatal

	return p
}

func (p *LoggerParams) ShowTimestamp() bool {
	if p.showTimestamp != nil {
		return *p.showTimestamp
	}

	return defaultShowTimestamp
}

func (p *LoggerParams) SetShowTimestamp(showTimestamp bool) *LoggerParams {
	p.showTimestamp = &showTimestamp

	return p
}

func (p *LoggerParams) Separator() string {
	if p.separator != nil {
		return *p.separator
	}

	return defaultSeparator
}

func (p *LoggerParams) SetSeparator(separator string) *LoggerParams {
	p.separator = &separator

	return p
}

func (p *LoggerParams) ShowSeparator() bool {
	if p.showSeparator != nil {
		return *p.showSeparator
	}

	return defaultShowSeparator
}

func (p *LoggerParams) SetShowSeparator(showSeparator bool) *LoggerParams {
	p.showSeparator = &showSeparator

	return p
}

func (p *LoggerParams) UseUnixTimestamp() bool {
	if p.useUnixTimestamp != nil {
		return *p.useUnixTimestamp
	}

	return defaultUseUnixTimestamp
}

func (p *LoggerParams) SetUseUnixTimestamp(useUnixTimestamp bool) *LoggerParams {
	p.useUnixTimestamp = &useUnixTimestamp

	return p
}

func (p *LoggerParams) UnixTimestampNano() bool {
	if p.unixTimestampNano != nil {
		return *p.unixTimestampNano
	}

	return defaultUnixTimestampNano
}

func (p *LoggerParams) SetUnixTimestampNano(unixTimestampNano bool) *LoggerParams {
	p.unixTimestampNano = &unixTimestampNano

	return p
}

func (p *LoggerParams) ColorizeMessages() bool {
	if p.colorizeMessages != nil {
		return *p.colorizeMessages
	}

	return defaultColorizeMessages
}

func (p *LoggerParams) SetColorizeMessages(colorizeMessages bool) *LoggerParams {
	p.colorizeMessages = &colorizeMessages

	return p
}

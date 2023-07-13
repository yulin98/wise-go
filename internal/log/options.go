package log

type LoggerOptions struct {
	DebugMode bool `mapstructure:"debug,omitempty"`
}

func NewLoggerOptions() *LoggerOptions {
	return &LoggerOptions{DebugMode: true}
}

func NewLog() LoggerOptions {
	return LoggerOptions{DebugMode: true}
}

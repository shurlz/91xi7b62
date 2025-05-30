package main

type Logger struct {
	config *LoggerConfig
}

// Further sort by Group IDs?

func (l *Logger) Error() {}

func (l *Logger) Info() {}

func (l *Logger) Message() {}

func (l *Logger) Warning() {}

func (l *Logger) Fatal() {}

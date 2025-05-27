package main

import (
	"fmt"
)

type Config struct{}

type Logger struct {
	config *Config
}

func (l *Logger) Error() {
	// Logs error, takes in string or in-built error type
}

func (l *Logger) Info() {
	// Logs String
}

func (l *Logger) Message() {
	// Logs String
}

func (l *Logger) Warning() {
	// Logs String
}

func (l *Logger) Fatal() {
	// Trips runtime
}

func main() {
	fmt.Println("Logger Base...")
}

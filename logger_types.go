package main

import (
	"errors"
)

type LogLevelType string

const (
	ErrorType   LogLevelType = "ERROR"
	InfoType    LogLevelType = "INFO"
	MessageType LogLevelType = "MESSAGE"
	FatalType   LogLevelType = "FATAL"
	WarningType LogLevelType = "WARNING"
)

func LogLevelIsValid(message interface{}) (bool, error) {
	switch mssg := message.(type) {
	case LogLevelType:
		switch mssg {
		case ErrorType, InfoType, MessageType, FatalType, WarningType:
			return true, nil
		default:
			return false, errors.New("invalid log level string")
		}
	default:
		return false, errors.New("invalid log level type")
	}
}

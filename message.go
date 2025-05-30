package main

import (
	"time"
)

type LogEventMessage struct {
	LogLevel LogLevelType `json:"log_level"`
	Message  *string      `json:"message"`
	Time     time.Time    `json:"time"`
	Data     interface{}  `json:"data"`
}

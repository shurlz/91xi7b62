package main

type FormatTypes string

const (
	Json     FormatTypes = "json"
	Regular  FormatTypes = "regular"
	Vertical FormatTypes = "vertical"
)

var validFormats = map[FormatTypes]struct{}{
	Json:     {},
	Regular:  {},
	Vertical: {},
}

func (f FormatTypes) IsValid() bool {
	_, ok := validFormats[f]
	return ok
}

type LoggerConfig struct {
	Format       FormatTypes
	Hex          string
	Stream       *StreamConfig
	CustomTopics bool
}

type StreamConfig struct {
	KafkaUrl string
}

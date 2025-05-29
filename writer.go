package main

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter(topic, brokerUrl string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerUrl},
		Topic:   topic,
	})
}

func PublishMessage(message LogEventMessage, writer *kafka.Writer) error {
	var msgByte []byte
	msgByte, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Key:     []byte{}, // Partition Key
		Value:   msgByte,
		Headers: []kafka.Header{},
	})
	if err != nil {
		return err
	}

	return nil
}

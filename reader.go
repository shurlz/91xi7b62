package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(topic, brokerUrl string, groupId *string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerUrl},
		GroupID:  *groupId,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}

// for { ConsumeMessage() } , constant stream
func ConsumeMessage(reader *kafka.Reader) {
	m, err := reader.ReadMessage(context.Background())
	if err != nil {
		// Panicki??? Maybe not?
		fmt.Println(err.Error())
	}
	fmt.Printf("Received message: %s = %s\n", string(m.Key), string(m.Value))
}

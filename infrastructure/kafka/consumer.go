package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, GroupID string, Topic string) *Consumer {
	
	return &Consumer{ reader: kafka.NewReader(kafka.ReaderConfig{ Brokers: brokers, GroupID: GroupID, Topic: Topic }) }
}

func(c *Consumer) Start(ctx context.Context) {
	msg, err := c.reader.ReadMessage(ctx)
	for {
		if err != nil {
		return
	}

	value := msg.Value
	topic := msg.Topic
	fmt.Printf("Message from the consumer %v is: %v", topic, value)
	// add process
	}


}

func(c *Consumer) Close() {
	c.reader.Close()
}


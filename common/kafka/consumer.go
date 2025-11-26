package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type consumer struct {
	reader *kafka.Reader
}

func NewConsumer(cfg *Config, groupID string, topic string) *consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     cfg.Brokers,
		GroupID:     groupID,
		Topic:       topic,
		StartOffset: kafka.FirstOffset,
	})
	return &consumer{reader: reader}
}

func (c *consumer) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return c.reader.ReadMessage(ctx)
}

func (c *consumer) Close() error {
	return c.reader.Close()
}

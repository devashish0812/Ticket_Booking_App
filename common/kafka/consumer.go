package kafka

import "github.com/segmentio/kafka-go"

type consumer struct {
	reader *kafka.Reader
}

func NewConsumer(cfg *Config, groupID string) *consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     cfg.Brokers,
		GroupID:     groupID,
		StartOffset: kafka.FirstOffset,
	})
	return &consumer{reader: reader}
}

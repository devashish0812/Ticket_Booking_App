package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(cfg *Config) *Producer {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       cfg.TLS,
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      cfg.Brokers,
		Dialer:       dialer,
		BatchTimeout: 100 * time.Millisecond,
		RequiredAcks: int(kafka.RequireAll),
		Balancer:     &kafka.LeastBytes{},
	})

	return &Producer{writer: writer}
}

func (p *Producer) Publish(topic string, event interface{}) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	msg := kafka.Message{
		Topic: topic,
		Value: payload,
	}

	retries := 3
	backoff := time.Second

	for i := 0; i < retries; i++ {
		err := p.writer.WriteMessages(context.Background(), msg)
		if err == nil {
			log.Printf("Event published to topic '%s'\n", topic)
			return nil
		}
		log.Printf("Publish failed (try %d/%d): %v", i+1, retries, err)
		time.Sleep(backoff)
		backoff *= 2
	}

	return fmt.Errorf("failed after retries: %w", err)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

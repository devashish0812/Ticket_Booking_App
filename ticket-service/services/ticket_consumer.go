package services

import (
	"context"
	"fmt"
	"log"

	customkafka "github.com/devashish0812/Ticketing_App/common/kafka"
	kafka "github.com/segmentio/kafka-go"
)

type Worker struct {
	topic   string
	groupID string
}

func NewWorker(topic string, groupID string) *Worker {
	return &Worker{
		topic:   topic,
		groupID: groupID,
	}
}

func (w *Worker) Run(ctx context.Context) {

	cfg := customkafka.LoadConfig()

	kConsumer := customkafka.NewConsumer(cfg, w.groupID, w.topic)
	defer func() {
		if err := kConsumer.Close(); err != nil {
			log.Printf("Error closing consumer: %v", err)
		}
	}()

	fmt.Printf("Worker started for Topic: %s | Group: %s\n", w.topic, w.groupID)

	msgChan := make(chan kafka.Message)

	// 4. Start the Pump (Go Routine)
	// Reads from Kafka -> Writes to Channel
	go func() {
		defer close(msgChan)
		for {
			if ctx.Err() != nil {
				return
			}

			msg, err := kConsumer.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Printf("Error reading message: %v\n", err)
				continue
			}

			select {
			case msgChan <- msg:
			case <-ctx.Done():
				return
			}
		}
	}()

	for msg := range msgChan {
		// Your processing logic goes here
		fmt.Printf("Received Key: %s | Value: %s\n", string(msg.Key), string(msg.Value))
	}
}

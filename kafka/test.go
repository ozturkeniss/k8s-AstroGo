package kafka

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TestKafkaConnection() error {
	brokers := []string{"localhost:9092"}
	topic := "test-topic"

	// Create producer
	producer, err := NewProducer(brokers)
	if err != nil {
		return fmt.Errorf("failed to create producer: %v", err)
	}
	defer producer.Close()

	// Create consumer
	messageReceived := make(chan bool)
	handler := func(message []byte) error {
		log.Printf("Received message: %s\n", string(message))
		messageReceived <- true
		return nil
	}

	consumer, err := NewConsumer(brokers, "test-group", []string{topic}, handler)
	if err != nil {
		return fmt.Errorf("failed to create consumer: %v", err)
	}
	defer consumer.Close()

	// Start consumer in a goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := consumer.Start(ctx); err != nil {
			log.Printf("Consumer error: %v\n", err)
		}
	}()

	// Wait a bit to ensure consumer is ready
	time.Sleep(1 * time.Second)

	// Send test message
	testMessage := "Hello Kafka!"
	if err := producer.SendMessage(topic, testMessage); err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	// Wait for message to be received
	select {
	case <-messageReceived:
		// Message received successfully
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("timeout waiting for message")
	}
}

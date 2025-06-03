package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer sarama.ConsumerGroup
	topics   []string
	handler  MessageHandler
}

type MessageHandler func(message []byte) error

func NewConsumer(brokers []string, groupID string, topics []string, handler MessageHandler) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %v", err)
	}

	return &Consumer{
		consumer: consumer,
		topics:   topics,
		handler:  handler,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	consumerGroup := &ConsumerGroup{
		handler: c.handler,
	}

	for {
		err := c.consumer.Consume(ctx, c.topics, consumerGroup)
		if err != nil {
			return fmt.Errorf("error from consumer: %v", err)
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}

type ConsumerGroup struct {
	handler MessageHandler
}

func (g *ConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (g *ConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (g *ConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		err := g.handler(message.Value)
		if err != nil {
			log.Printf("Error processing message: %v\n", err)
			continue
		}
		session.MarkMessage(message, "")
	}
	return nil
}

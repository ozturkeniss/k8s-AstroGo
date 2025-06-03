package main

import (
	"log"

	"astrogo/kafka"
)

func main() {
	if err := kafka.TestKafkaConnection(); err != nil {
		log.Fatalf("Kafka test failed: %v", err)
	}
	log.Println("Kafka test completed successfully!")
}

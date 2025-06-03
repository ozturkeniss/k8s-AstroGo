package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	energy "astrogo/grpc/energy/grpc/protos"
	"astrogo/internal/energy/handler"
	"astrogo/internal/energy/repository"
	"astrogo/internal/energy/service"
	"astrogo/kafka"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Energy Service starting...")

	// Initialize repository
	log.Println("Initializing repository...")
	repo, err := repository.NewEnergyRepository()
	if err != nil {
		log.Printf("Failed to initialize repository: %v", err)
		return
	}
	log.Println("Repository initialized successfully")

	// Initialize Kafka producer
	log.Println("Initializing Kafka producer...")
	brokers := []string{os.Getenv("KAFKA_BROKERS")}
	producer, err := kafka.NewProducer(brokers)
	if err != nil {
		log.Printf("Failed to create Kafka producer: %v", err)
		// Uygulamanın çökmesini engellemek için producer'ı nil olarak bırakıyoruz
		producer = nil
	} else {
		defer producer.Close()
		log.Println("Kafka producer initialized successfully")
	}

	// Initialize service
	log.Println("Initializing service...")
	energyService := service.NewEnergyService(repo, producer)
	log.Println("Service initialized successfully")

	// Initialize HTTP handler
	log.Println("Initializing HTTP handler...")
	energyHandler := handler.NewEnergyHandler(energyService)
	log.Println("HTTP handler initialized successfully")

	// Start Kafka consumer
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Println("Starting Kafka consumer...")
		if err := energyService.StartKafkaConsumer(ctx); err != nil {
			log.Printf("Failed to start Kafka consumer: %v", err)
		}
	}()

	// Initialize gRPC server
	log.Println("Initializing gRPC server...")
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Printf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	energy.RegisterEnergyServiceServer(s, energyService)
	log.Println("gRPC server initialized successfully")

	// HTTP routes
	http.HandleFunc("/energy/stock", energyHandler.GetCurrentStock)
	http.HandleFunc("/energy/add", energyHandler.AddEnergy)
	http.HandleFunc("/energy/consume", energyHandler.ConsumeEnergy)
	http.HandleFunc("/energy/history", energyHandler.GetTransactionHistory)
	http.HandleFunc("/energy/check", energyHandler.CheckAvailability)
	log.Println("HTTP routes registered")

	// Start HTTP server
	go func() {
		log.Println("Starting HTTP server on :8085")
		if err := http.ListenAndServe(":8085", nil); err != nil {
			log.Printf("Failed to start HTTP server: %v", err)
		}
	}()

	// Start gRPC server
	go func() {
		log.Println("Starting gRPC server on :8083")
		if err := s.Serve(lis); err != nil {
			log.Printf("Failed to serve: %v", err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	s.GracefulStop()
}

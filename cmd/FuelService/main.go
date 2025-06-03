package main

import (
	"log"
	"net"
	"net/http"

	fuel "astrogo/grpc/fuel/grpc/protos"
	"astrogo/internal/fuel/handler"
	"astrogo/internal/fuel/repository"
	"astrogo/internal/fuel/service"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Fuel Service starting...")

	// Initialize repository
	log.Println("Initializing repository...")
	repo, err := repository.NewFuelRepository()
	if err != nil {
		log.Printf("Failed to initialize repository: %v", err)
		return
	}
	log.Println("Repository initialized successfully")

	// Initialize service
	log.Println("Initializing service...")
	fuelService := service.NewFuelService(repo)
	log.Println("Service initialized successfully")

	// Initialize HTTP handler
	log.Println("Initializing HTTP handler...")
	fuelHandler := handler.NewFuelHandler(fuelService)
	log.Println("HTTP handler initialized successfully")

	// Create gRPC server
	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Printf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	fuel.RegisterFuelServiceServer(s, fuelService)

	// HTTP routes
	http.HandleFunc("/fuel/stock", fuelHandler.GetCurrentStock)
	http.HandleFunc("/fuel/add", fuelHandler.AddFuel)
	http.HandleFunc("/fuel/consume", fuelHandler.ConsumeFuel)
	http.HandleFunc("/fuel/history", fuelHandler.GetTransactionHistory)
	http.HandleFunc("/fuel/check", fuelHandler.CheckAvailability)
	log.Println("HTTP routes registered")

	// Start HTTP server
	go func() {
		log.Println("Starting HTTP server on :8088")
		if err := http.ListenAndServe(":8088", nil); err != nil {
			log.Printf("Failed to start HTTP server: %v", err)
		}
	}()

	// Start gRPC server
	log.Println("Starting gRPC server on :8085")
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v", err)
	}
}

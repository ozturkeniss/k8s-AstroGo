package main

import (
	"astrogo/internal/astronaut/handler"
	"astrogo/internal/astronaut/repository"
	"astrogo/internal/astronaut/service"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	// Initialize repository
	repo, err := repository.NewAstronautRepository()
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Initialize service
	svc := service.NewAstronautService(repo)

	// Initialize handler
	h := handler.NewAstronautHandler(svc)

	// Define routes
	http.HandleFunc("/astronaut/create", h.CreateAstronaut)
	http.HandleFunc("/astronaut/get", h.GetAstronaut)
	http.HandleFunc("/astronaut/available", h.ListAvailableAstronauts)
	http.HandleFunc("/astronaut/assign", h.AssignToMission)
	http.HandleFunc("/astronaut/complete", h.CompleteMission)

	// Start HTTP server
	go func() {
		log.Println("Starting HTTP server on :8082")
		if err := http.ListenAndServe(":8082", nil); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Start gRPC server
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// TODO: Register gRPC service
	// protos.RegisterAstronautServiceServer(s, svc)

	log.Println("Starting gRPC server on :8081")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

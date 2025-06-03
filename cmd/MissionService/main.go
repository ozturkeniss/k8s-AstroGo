package main

import (
	"log"
	"net"
	"net/http"

	mission "astrogo/grpc/mission/grpc/protos"
	"astrogo/internal/mission/handler"
	"astrogo/internal/mission/repository"
	"astrogo/internal/mission/service"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Mission Service starting...")

	// Initialize repository
	log.Println("Initializing repository...")
	missionRepo, err := repository.NewMissionRepository()
	if err != nil {
		log.Printf("Failed to initialize repository: %v", err)
		return
	}
	log.Println("Repository initialized successfully")

	// Initialize service
	log.Println("Initializing service...")
	missionService := service.NewMissionService(missionRepo)
	log.Println("Service initialized successfully")

	// Initialize handler
	log.Println("Initializing HTTP handler...")
	missionHandler := handler.NewMissionHandler(missionService)
	log.Println("HTTP handler initialized successfully")

	// Create gRPC server
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Printf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	mission.RegisterMissionServiceServer(s, missionService)

	// HTTP routes
	http.HandleFunc("/missions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			missionHandler.CreateMission(w, r)
		case http.MethodGet:
			if r.URL.Query().Get("id") != "" {
				missionHandler.GetMission(w, r)
			} else {
				missionHandler.ListMissions(w, r)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/missions/start", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		missionHandler.StartMission(w, r)
	})

	http.HandleFunc("/missions/complete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		missionHandler.CompleteMission(w, r)
	})

	http.HandleFunc("/missions/fail", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		missionHandler.FailMission(w, r)
	})

	http.HandleFunc("/missions/assign", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		missionHandler.AssignAstronaut(w, r)
	})
	log.Println("HTTP routes registered")

	// Start HTTP server
	go func() {
		log.Println("Starting HTTP server on :8086")
		if err := http.ListenAndServe(":8086", nil); err != nil {
			log.Printf("Failed to start HTTP server: %v", err)
		}
	}()

	// Start gRPC server
	log.Println("Starting gRPC server on :8084")
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v", err)
	}
}

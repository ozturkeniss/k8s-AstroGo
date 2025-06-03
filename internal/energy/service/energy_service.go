package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"astrogo/grpc/energy/grpc/protos"
	"astrogo/internal/energy/model"
	"astrogo/internal/energy/repository"
	"astrogo/kafka"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EnergyService struct {
	protos.UnimplementedEnergyServiceServer
	repo     *repository.EnergyRepository
	producer *kafka.Producer
}

func NewEnergyService(repo *repository.EnergyRepository, producer *kafka.Producer) *EnergyService {
	return &EnergyService{
		repo:     repo,
		producer: producer,
	}
}

func (s *EnergyService) StartKafkaConsumer(ctx context.Context) error {
	consumer, err := kafka.NewConsumer(
		[]string{os.Getenv("KAFKA_BROKERS")},
		"energy-service",
		[]string{kafka.MissionCreatedEvent, kafka.MissionStartedEvent},
		s.handleMessage,
	)
	if err != nil {
		return fmt.Errorf("failed to create consumer: %v", err)
	}
	defer consumer.Close()

	return consumer.Start(ctx)
}

func (s *EnergyService) handleMessage(message []byte) error {
	// Parse the message to determine the event type
	var event map[string]interface{}
	if err := json.Unmarshal(message, &event); err != nil {
		return fmt.Errorf("failed to unmarshal message: %v", err)
	}

	// Handle mission created event
	if event["mission_id"] != nil {
		missionID := uint64(event["mission_id"].(float64))
		energyReq := event["energy_req"].(float64)

		// Check if we have enough energy
		stock, err := s.repo.GetStock()
		if err != nil {
			return fmt.Errorf("failed to get energy stock: %v", err)
		}

		if stock.Amount < energyReq {
			// Notify that we don't have enough energy
			log.Printf("Not enough energy for mission %d. Required: %f, Available: %f", missionID, energyReq, stock.Amount)
			return nil
		}

		// Consume energy
		if err := s.repo.ConsumeEnergy(energyReq); err != nil {
			return fmt.Errorf("failed to consume energy: %v", err)
		}

		// Record transaction
		transaction := &model.EnergyTransaction{
			Amount:      energyReq,
			Type:        model.TransactionTypeConsumption,
			Description: fmt.Sprintf("Energy consumed for mission %d", missionID),
			CreatedAt:   time.Now().Format(time.RFC3339),
		}
		if err := s.repo.CreateTransaction(transaction); err != nil {
			return fmt.Errorf("failed to record transaction: %v", err)
		}

		// Publish energy consumed event
		event := kafka.EnergyConsumed{
			MissionID:  missionID,
			Amount:     energyReq,
			ConsumedAt: transaction.CreatedAt,
		}
		if err := s.producer.SendMessage(kafka.EnergyConsumedEvent, event); err != nil {
			log.Printf("Failed to publish energy consumed event: %v", err)
		}
	}

	return nil
}

func (s *EnergyService) GetCurrentStock(ctx context.Context, req *protos.GetCurrentStockRequest) (*protos.GetCurrentStockResponse, error) {
	stock, err := s.repo.GetCurrentStock()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &protos.GetCurrentStockResponse{Amount: stock.Amount}, nil
}

func (s *EnergyService) AddEnergy(ctx context.Context, req *protos.AddEnergyRequest) (*protos.AddEnergyResponse, error) {
	err := s.repo.AddEnergy(req.Amount, req.Description)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &protos.AddEnergyResponse{Success: true}, nil
}

func (s *EnergyService) ConsumeEnergy(ctx context.Context, req *protos.ConsumeEnergyRequest) (*protos.ConsumeEnergyResponse, error) {
	err := s.repo.ConsumeEnergy(req.Amount)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &protos.ConsumeEnergyResponse{Success: true}, nil
}

func (s *EnergyService) GetTransactionHistory(ctx context.Context, req *protos.GetTransactionHistoryRequest) (*protos.GetTransactionHistoryResponse, error) {
	transactions, err := s.repo.GetTransactionHistory()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var protoTransactions []*protos.EnergyTransaction
	for _, t := range transactions {
		protoTransactions = append(protoTransactions, &protos.EnergyTransaction{
			Id:          t.ID,
			Amount:      t.Amount,
			Type:        protos.TransactionType(int32(t.Type)),
			Description: t.Description,
			CreatedAt:   t.CreatedAt,
		})
	}

	return &protos.GetTransactionHistoryResponse{Transactions: protoTransactions}, nil
}

func (s *EnergyService) CheckAvailability(ctx context.Context, req *protos.CheckAvailabilityRequest) (*protos.CheckAvailabilityResponse, error) {
	available, err := s.repo.CheckAvailability(req.Amount)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &protos.CheckAvailabilityResponse{Available: available}, nil
}

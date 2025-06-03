package service

import (
	"context"
	"fmt"
	"time"

	"astrogo/grpc/fuel/grpc/protos"
	"astrogo/internal/fuel/model"
	"astrogo/internal/fuel/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FuelService struct {
	protos.UnimplementedFuelServiceServer
	repo *repository.FuelRepository
}

func NewFuelService(repo *repository.FuelRepository) *FuelService {
	return &FuelService{
		repo: repo,
	}
}

func (s *FuelService) GetCurrentStock(ctx context.Context, req *protos.GetCurrentStockRequest) (*protos.GetCurrentStockResponse, error) {
	stock, err := s.repo.GetCurrentStock()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get current stock: %v", err))
	}

	return &protos.GetCurrentStockResponse{
		Stock: stock,
	}, nil
}

func (s *FuelService) AddFuel(ctx context.Context, req *protos.AddFuelRequest) (*protos.AddFuelResponse, error) {
	transaction := &model.FuelTransaction{
		Type:      model.TransactionTypeAdd,
		Amount:    req.Amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if err := s.repo.AddTransaction(transaction); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to add fuel: %v", err))
	}

	return &protos.AddFuelResponse{
		Success: true,
		Message: fmt.Sprintf("Added %.2f units of fuel", req.Amount),
	}, nil
}

func (s *FuelService) ConsumeFuel(ctx context.Context, req *protos.ConsumeFuelRequest) (*protos.ConsumeFuelResponse, error) {
	stock, err := s.repo.GetCurrentStock()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get current stock: %v", err))
	}

	if stock < req.Amount {
		return nil, status.Error(codes.FailedPrecondition, "insufficient fuel stock")
	}

	transaction := &model.FuelTransaction{
		Type:      model.TransactionTypeConsume,
		Amount:    req.Amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if err := s.repo.AddTransaction(transaction); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to consume fuel: %v", err))
	}

	return &protos.ConsumeFuelResponse{
		Success: true,
		Message: fmt.Sprintf("Consumed %.2f units of fuel", req.Amount),
	}, nil
}

func (s *FuelService) GetTransactionHistory(ctx context.Context, req *protos.GetTransactionHistoryRequest) (*protos.GetTransactionHistoryResponse, error) {
	transactions, err := s.repo.GetTransactionHistory()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get transaction history: %v", err))
	}

	var protoTransactions []*protos.FuelTransaction
	for _, t := range transactions {
		protoTransactions = append(protoTransactions, &protos.FuelTransaction{
			Id:        t.ID,
			Type:      protos.TransactionType(t.Type),
			Amount:    t.Amount,
			Timestamp: t.Timestamp,
		})
	}

	return &protos.GetTransactionHistoryResponse{
		Transactions: protoTransactions,
	}, nil
}

func (s *FuelService) CheckAvailability(ctx context.Context, req *protos.CheckAvailabilityRequest) (*protos.CheckAvailabilityResponse, error) {
	stock, err := s.repo.GetCurrentStock()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get current stock: %v", err))
	}

	return &protos.CheckAvailabilityResponse{
		Available: stock >= req.Amount,
		Stock:     stock,
	}, nil
}

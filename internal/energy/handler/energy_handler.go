package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"astrogo/grpc/energy/grpc/protos"
	"astrogo/internal/energy/service"
)

type EnergyHandler struct {
	service *service.EnergyService
}

func NewEnergyHandler(svc *service.EnergyService) *EnergyHandler {
	return &EnergyHandler{
		service: svc,
	}
}

func (h *EnergyHandler) GetCurrentStock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := h.service.GetCurrentStock(r.Context(), &protos.GetCurrentStockRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{
		"stock": resp.Amount,
	})
}

func (h *EnergyHandler) AddEnergy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.AddEnergy(r.Context(), &protos.AddEnergyRequest{
		Amount: req.Amount,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{
		"success": resp.Success,
	})
}

func (h *EnergyHandler) ConsumeEnergy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.ConsumeEnergy(r.Context(), &protos.ConsumeEnergyRequest{
		Amount: req.Amount,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{
		"success": resp.Success,
	})
}

func (h *EnergyHandler) GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := int32(10) // default limit
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	resp, err := h.service.GetTransactionHistory(r.Context(), &protos.GetTransactionHistoryRequest{
		Limit: limit,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp.Transactions)
}

func (h *EnergyHandler) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	amountStr := r.URL.Query().Get("amount")
	if amountStr == "" {
		http.Error(w, "amount parameter is required", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.service.CheckAvailability(r.Context(), &protos.CheckAvailabilityRequest{
		Amount: amount,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{
		"available": resp.Available,
	})
}

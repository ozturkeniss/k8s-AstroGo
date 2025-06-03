package handler

import (
	"astrogo/internal/astronaut/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type AstronautHandler struct {
	service *service.AstronautService
}

func NewAstronautHandler(service *service.AstronautService) *AstronautHandler {
	return &AstronautHandler{service: service}
}

func (h *AstronautHandler) CreateAstronaut(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	astronaut, err := h.service.CreateAstronaut(request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(astronaut)
}

func (h *AstronautHandler) GetAstronaut(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid astronaut ID", http.StatusBadRequest)
		return
	}

	astronaut, err := h.service.GetAstronaut(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(astronaut)
}

func (h *AstronautHandler) ListAvailableAstronauts(w http.ResponseWriter, r *http.Request) {
	astronauts, err := h.service.ListAvailableAstronauts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(astronauts)
}

func (h *AstronautHandler) AssignToMission(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AstronautID uint64 `json:"astronaut_id"`
		MissionID   uint64 `json:"mission_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.AssignToMission(request.AstronautID, request.MissionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AstronautHandler) CompleteMission(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AstronautID uint64 `json:"astronaut_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.CompleteMission(request.AstronautID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

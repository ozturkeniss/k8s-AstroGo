package handler

import (
	"astrogo/grpc/mission/grpc/protos"
	"astrogo/internal/mission/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type MissionHandler struct {
	service *service.MissionService
}

func NewMissionHandler(service *service.MissionService) *MissionHandler {
	return &MissionHandler{service: service}
}

func (h *MissionHandler) CreateMission(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Type        int32   `json:"type"`
		EnergyReq   float64 `json:"energy_req"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req := &protos.CreateMissionRequest{
		Name:        request.Name,
		Description: request.Description,
		Type:        protos.MissionType(request.Type),
		EnergyReq:   request.EnergyReq,
	}

	resp, err := h.service.CreateMission(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Mission)
}

func (h *MissionHandler) GetMission(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid mission ID", http.StatusBadRequest)
		return
	}

	req := &protos.GetMissionRequest{
		MissionId: id,
	}

	resp, err := h.service.GetMission(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Mission)
}

func (h *MissionHandler) ListMissions(w http.ResponseWriter, r *http.Request) {
	statusStr := r.URL.Query().Get("status")
	typeStr := r.URL.Query().Get("type")

	req := &protos.ListMissionsRequest{}

	if statusStr != "" {
		status, err := strconv.ParseInt(statusStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid status", http.StatusBadRequest)
			return
		}
		req.Status = protos.MissionStatus(status)
	}

	if typeStr != "" {
		missionType, err := strconv.ParseInt(typeStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid type", http.StatusBadRequest)
			return
		}
		req.Type = protos.MissionType(missionType)
	}

	resp, err := h.service.ListMissions(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Missions)
}

func (h *MissionHandler) StartMission(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ID uint64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req := &protos.StartMissionRequest{
		MissionId: request.ID,
	}

	resp, err := h.service.StartMission(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Mission)
}

func (h *MissionHandler) CompleteMission(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ID uint64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req := &protos.CompleteMissionRequest{
		MissionId: request.ID,
	}

	resp, err := h.service.CompleteMission(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Mission)
}

func (h *MissionHandler) FailMission(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ID uint64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req := &protos.FailMissionRequest{
		MissionId: request.ID,
	}

	resp, err := h.service.FailMission(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Mission)
}

func (h *MissionHandler) AssignAstronaut(w http.ResponseWriter, r *http.Request) {
	var request struct {
		MissionID   uint64 `json:"mission_id"`
		AstronautID uint64 `json:"astronaut_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req := &protos.AssignAstronautRequest{
		MissionId:   request.MissionID,
		AstronautId: request.AstronautID,
	}

	resp, err := h.service.AssignAstronaut(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Mission)
}

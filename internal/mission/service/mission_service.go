package service

import (
	"context"
	"fmt"
	"time"

	"astrogo/grpc/mission/grpc/protos"
	"astrogo/internal/mission/model"
	"astrogo/internal/mission/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MissionService struct {
	protos.UnimplementedMissionServiceServer
	repo *repository.MissionRepository
}

func NewMissionService(repo *repository.MissionRepository) *MissionService {
	return &MissionService{
		repo: repo,
	}
}

func (s *MissionService) CreateMission(ctx context.Context, req *protos.CreateMissionRequest) (*protos.CreateMissionResponse, error) {
	mission := &model.Mission{
		Name:        req.Name,
		Description: req.Description,
		Type:        model.MissionType(req.Type),
		Status:      model.MissionStatusPending,
		EnergyReq:   req.EnergyReq,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	if err := s.repo.Create(mission); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create mission: %v", err))
	}

	return &protos.CreateMissionResponse{
		Mission: &protos.Mission{
			Id:          mission.ID,
			Name:        mission.Name,
			Description: mission.Description,
			Type:        protos.MissionType(mission.Type),
			Status:      protos.MissionStatus(mission.Status),
			EnergyReq:   mission.EnergyReq,
			CreatedAt:   mission.CreatedAt,
			UpdatedAt:   mission.UpdatedAt,
		},
	}, nil
}

func (s *MissionService) GetMission(ctx context.Context, req *protos.GetMissionRequest) (*protos.GetMissionResponse, error) {
	mission, err := s.repo.GetByID(req.MissionId)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("mission not found: %v", err))
	}

	return &protos.GetMissionResponse{
		Mission: &protos.Mission{
			Id:          mission.ID,
			Name:        mission.Name,
			Description: mission.Description,
			Type:        protos.MissionType(mission.Type),
			Status:      protos.MissionStatus(mission.Status),
			EnergyReq:   mission.EnergyReq,
			CreatedAt:   mission.CreatedAt,
			UpdatedAt:   mission.UpdatedAt,
		},
	}, nil
}

func (s *MissionService) ListMissions(ctx context.Context, req *protos.ListMissionsRequest) (*protos.ListMissionsResponse, error) {
	missions, err := s.repo.ListAll()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to list missions: %v", err))
	}

	var protoMissions []*protos.Mission
	for _, m := range missions {
		protoMissions = append(protoMissions, &protos.Mission{
			Id:          m.ID,
			Name:        m.Name,
			Description: m.Description,
			Type:        protos.MissionType(m.Type),
			Status:      protos.MissionStatus(m.Status),
			EnergyReq:   m.EnergyReq,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
		})
	}

	return &protos.ListMissionsResponse{
		Missions: protoMissions,
	}, nil
}

func (s *MissionService) StartMission(ctx context.Context, req *protos.StartMissionRequest) (*protos.StartMissionResponse, error) {
	mission, err := s.repo.GetByID(req.MissionId)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("mission not found: %v", err))
	}

	if mission.Status != model.MissionStatusPending {
		return nil, status.Error(codes.FailedPrecondition, "mission is not in pending status")
	}

	mission.Status = model.MissionStatusInProgress
	mission.StartedAt = time.Now().Format(time.RFC3339)
	mission.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := s.repo.Update(mission); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update mission: %v", err))
	}

	return &protos.StartMissionResponse{
		Mission: &protos.Mission{
			Id:          mission.ID,
			Name:        mission.Name,
			Description: mission.Description,
			Type:        protos.MissionType(mission.Type),
			Status:      protos.MissionStatus(mission.Status),
			EnergyReq:   mission.EnergyReq,
			CreatedAt:   mission.CreatedAt,
			UpdatedAt:   mission.UpdatedAt,
			StartedAt:   mission.StartedAt,
		},
	}, nil
}

func (s *MissionService) CompleteMission(ctx context.Context, req *protos.CompleteMissionRequest) (*protos.CompleteMissionResponse, error) {
	mission, err := s.repo.GetByID(req.MissionId)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("mission not found: %v", err))
	}

	if mission.Status != model.MissionStatusInProgress {
		return nil, status.Error(codes.FailedPrecondition, "mission is not in progress")
	}

	mission.Status = model.MissionStatusCompleted
	mission.CompletedAt = time.Now().Format(time.RFC3339)
	mission.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := s.repo.Update(mission); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update mission: %v", err))
	}

	return &protos.CompleteMissionResponse{
		Mission: &protos.Mission{
			Id:          mission.ID,
			Name:        mission.Name,
			Description: mission.Description,
			Type:        protos.MissionType(mission.Type),
			Status:      protos.MissionStatus(mission.Status),
			EnergyReq:   mission.EnergyReq,
			CreatedAt:   mission.CreatedAt,
			UpdatedAt:   mission.UpdatedAt,
			StartedAt:   mission.StartedAt,
			CompletedAt: mission.CompletedAt,
		},
	}, nil
}

func (s *MissionService) FailMission(ctx context.Context, req *protos.FailMissionRequest) (*protos.FailMissionResponse, error) {
	mission, err := s.repo.GetByID(req.MissionId)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("mission not found: %v", err))
	}

	if mission.Status != model.MissionStatusInProgress {
		return nil, status.Error(codes.FailedPrecondition, "mission is not in progress")
	}

	mission.Status = model.MissionStatusFailed
	mission.FailedAt = time.Now().Format(time.RFC3339)
	mission.UpdatedAt = time.Now().Format(time.RFC3339)
	mission.FailureReason = req.Reason

	if err := s.repo.Update(mission); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update mission: %v", err))
	}

	return &protos.FailMissionResponse{
		Mission: &protos.Mission{
			Id:            mission.ID,
			Name:          mission.Name,
			Description:   mission.Description,
			Type:          protos.MissionType(mission.Type),
			Status:        protos.MissionStatus(mission.Status),
			EnergyReq:     mission.EnergyReq,
			CreatedAt:     mission.CreatedAt,
			UpdatedAt:     mission.UpdatedAt,
			StartedAt:     mission.StartedAt,
			FailedAt:      mission.FailedAt,
			FailureReason: mission.FailureReason,
		},
	}, nil
}

func (s *MissionService) AssignAstronaut(ctx context.Context, req *protos.AssignAstronautRequest) (*protos.AssignAstronautResponse, error) {
	mission, err := s.repo.GetByID(req.MissionId)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("mission not found: %v", err))
	}

	if mission.Status != model.MissionStatusPending {
		return nil, status.Error(codes.FailedPrecondition, "mission is not in pending status")
	}

	mission.AstronautID = req.AstronautId
	mission.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := s.repo.Update(mission); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update mission: %v", err))
	}

	return &protos.AssignAstronautResponse{
		Mission: &protos.Mission{
			Id:          mission.ID,
			Name:        mission.Name,
			Description: mission.Description,
			Type:        protos.MissionType(mission.Type),
			Status:      protos.MissionStatus(mission.Status),
			EnergyReq:   mission.EnergyReq,
			CreatedAt:   mission.CreatedAt,
			UpdatedAt:   mission.UpdatedAt,
			AstronautId: mission.AstronautID,
		},
	}, nil
}

package service

import (
	"astrogo/internal/astronaut/model"
	"astrogo/internal/astronaut/repository"
	"errors"
)

type AstronautService struct {
	repo *repository.AstronautRepository
}

func NewAstronautService(repo *repository.AstronautRepository) *AstronautService {
	return &AstronautService{repo: repo}
}

func (s *AstronautService) CreateAstronaut(name string) (*model.Astronaut, error) {
	astronaut := &model.Astronaut{
		Name:   name,
		Status: model.Available,
	}

	err := s.repo.Create(astronaut)
	if err != nil {
		return nil, err
	}

	return astronaut, nil
}

func (s *AstronautService) GetAstronaut(id uint64) (*model.Astronaut, error) {
	return s.repo.GetByID(id)
}

func (s *AstronautService) ListAvailableAstronauts() ([]model.Astronaut, error) {
	return s.repo.ListAvailable()
}

func (s *AstronautService) AssignToMission(astronautID uint64, missionID uint64) error {
	astronaut, err := s.repo.GetByID(astronautID)
	if err != nil {
		return err
	}

	if astronaut.Status != model.Available {
		return errors.New("astronaut is not available for mission assignment")
	}

	return s.repo.AssignMission(astronautID, missionID)
}

func (s *AstronautService) CompleteMission(astronautID uint64) error {
	astronaut, err := s.repo.GetByID(astronautID)
	if err != nil {
		return err
	}

	if astronaut.Status != model.Unavailable {
		return errors.New("astronaut is not on a mission")
	}

	astronaut.Status = model.Available
	astronaut.CurrentMissionID = nil

	return s.repo.Update(astronaut)
}

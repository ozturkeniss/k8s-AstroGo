package repository

import (
	"astrogo/internal/mission/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MissionRepository struct {
	db *gorm.DB
}

func NewMissionRepository() (*MissionRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&model.Mission{})
	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
		return nil, err
	}

	return &MissionRepository{db: db}, nil
}

func (r *MissionRepository) Create(mission *model.Mission) error {
	return r.db.Create(mission).Error
}

func (r *MissionRepository) GetByID(id uint64) (*model.Mission, error) {
	var mission model.Mission
	if err := r.db.First(&mission, id).Error; err != nil {
		return nil, err
	}
	return &mission, nil
}

func (r *MissionRepository) Update(mission *model.Mission) error {
	return r.db.Save(mission).Error
}

func (r *MissionRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Mission{}, id).Error
}

func (r *MissionRepository) ListAll() ([]model.Mission, error) {
	var missions []model.Mission
	if err := r.db.Find(&missions).Error; err != nil {
		return nil, err
	}
	return missions, nil
}

func (r *MissionRepository) ListByStatus(status model.MissionStatus) ([]model.Mission, error) {
	var missions []model.Mission
	if err := r.db.Where("status = ?", status).Find(&missions).Error; err != nil {
		return nil, err
	}
	return missions, nil
}

func (r *MissionRepository) ListByType(missionType model.MissionType) ([]model.Mission, error) {
	var missions []model.Mission
	if err := r.db.Where("type = ?", missionType).Find(&missions).Error; err != nil {
		return nil, err
	}
	return missions, nil
}

func (r *MissionRepository) UpdateStatus(id uint64, status model.MissionStatus) error {
	return r.db.Model(&model.Mission{}).Where("id = ?", id).Update("status", status).Error
}

func (r *MissionRepository) AssignAstronaut(id uint64, astronautID uint64) error {
	return r.db.Model(&model.Mission{}).Where("id = ?", id).Update("astronaut_id", astronautID).Error
}

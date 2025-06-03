package repository

import (
	"astrogo/internal/astronaut/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AstronautRepository struct {
	db *gorm.DB
}

func NewAstronautRepository() (*AstronautRepository, error) {
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
	err = db.AutoMigrate(&model.Astronaut{})
	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
		return nil, err
	}

	return &AstronautRepository{db: db}, nil
}

func (r *AstronautRepository) Create(astronaut *model.Astronaut) error {
	return r.db.Create(astronaut).Error
}

func (r *AstronautRepository) GetByID(id uint64) (*model.Astronaut, error) {
	var astronaut model.Astronaut
	err := r.db.First(&astronaut, id).Error
	if err != nil {
		return nil, err
	}
	return &astronaut, nil
}

func (r *AstronautRepository) Update(astronaut *model.Astronaut) error {
	return r.db.Save(astronaut).Error
}

func (r *AstronautRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Astronaut{}, id).Error
}

func (r *AstronautRepository) ListAvailable() ([]model.Astronaut, error) {
	var astronauts []model.Astronaut
	err := r.db.Where("status = ?", model.Available).Find(&astronauts).Error
	return astronauts, err
}

func (r *AstronautRepository) UpdateStatus(id uint64, status model.AvailabilityStatus) error {
	return r.db.Model(&model.Astronaut{}).Where("id = ?", id).Update("status", status).Error
}

func (r *AstronautRepository) AssignMission(id uint64, missionID uint64) error {
	return r.db.Model(&model.Astronaut{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":             model.Unavailable,
		"current_mission_id": missionID,
	}).Error
}

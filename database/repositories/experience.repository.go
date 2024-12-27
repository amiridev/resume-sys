package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type ExperienceRepositoryInterface interface {
	Connection() *gorm.DB
	Create(experience models.Experience) models.Experience
	List(userId string) []models.Experience
	Update(id string, experience models.Experience) models.Experience
	Delete(id string)
}

type ExperienceRepository struct {
	DB *gorm.DB
}

func NewExperienceRepository() *ExperienceRepository {
	return &ExperienceRepository{
		DB: database.Connection(),
	}
}

func (repo *ExperienceRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *ExperienceRepository) Create(Experience models.Experience) models.Experience {
	var newExperience models.Experience
	repo.DB.Create(&Experience).Scan(&newExperience)
	return newExperience
}

func (repo *ExperienceRepository) List(userId string) []models.Experience {
	var experiences []models.Experience
	repo.DB.Table("eperiences").Find(&experiences, "user_id = ?", userId)
	return experiences
}

func (repo *ExperienceRepository) Update(id string, updatedExperience models.Experience) (models.Experience, error) {
	var existingExperience models.Experience
	result := repo.DB.First(&existingExperience, "id = ?", id)
	if result.Error != nil {
		return existingExperience, result.Error
	}
	repo.DB.Model(&existingExperience).Updates(updatedExperience)
	return existingExperience, nil
}

func (repo *ExperienceRepository) Delete(id string) error {
	result := repo.DB.Delete(&models.Experience{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found") // اگر رکوردی پیدا نشود
	}
	return nil
}

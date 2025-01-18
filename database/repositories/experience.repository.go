package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type ExperienceRepositoryInterface interface {
	Connection() *gorm.DB
	Create(experience models.Experience) (models.Experience, error)
	List(userId string) ([]models.Experience, error)
	Show(experienceId string) (models.Experience, error)
	Update(id string, experience models.Experience) (models.Experience, error)
	Delete(id string) error
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

func (repo *ExperienceRepository) Create(experience models.Experience) (models.Experience, error) {
	result := repo.DB.Create(&experience)
	if result.Error != nil {
		return models.Experience{}, result.Error
	}
	return experience, nil
}

func (repo *ExperienceRepository) List(userId string) ([]models.Experience, error) {
	var experiences []models.Experience
	result := repo.DB.Where("user_id = ?", userId).Find(&experiences)
	if result.Error != nil {
		return experiences, result.Error
	}
	return experiences, nil
}

func (repo *ExperienceRepository) Show(experienceId string) (models.Experience, error) {
	var experience models.Experience

	result := repo.DB.First(&experience, "id = ?", experienceId)

	if result.Error != nil {
		return experience, result.Error
	}

	return experience, nil
}

func (repo *ExperienceRepository) Update(id string, updatedExperience models.Experience) (models.Experience, error) {
	var existingExperience models.Experience
	result := repo.DB.First(&existingExperience, "id = ?", id)
	if result.Error != nil {
		return existingExperience, result.Error
	}

	if err := repo.DB.Model(&existingExperience).Updates(updatedExperience).Error; err != nil {
		return existingExperience, err
	}
	return existingExperience, nil
}

func (repo *ExperienceRepository) Delete(id string) error {
	result := repo.DB.Delete(&models.Experience{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found") // رکورد پیدا نشد
	}
	return nil
}

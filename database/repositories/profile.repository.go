package repositories

import (
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type ProfileRepositoryInterface interface {
	Connection() *gorm.DB
	Create(profile models.Profile) (models.Profile, error)
	List(userId string) ([]models.Profile, error)
	Update(id string, profile models.Profile) (models.Profile, error)
}

type ProfileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{
		DB: database.Connection(),
	}
}

func (repo *ProfileRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *ProfileRepository) Create(profile models.Profile) (models.Profile, error) {
	result := repo.DB.Create(&profile)
	if result.Error != nil {
		return models.Profile{}, result.Error
	}
	return profile, nil
}

func (repo *ProfileRepository) List(userId string) ([]models.Profile, error) {
	var profiles []models.Profile
	result := repo.DB.Where("user_id = ?", userId).Find(&profiles)
	if result.Error != nil {
		return nil, result.Error
	}
	return profiles, nil
}

func (repo *ProfileRepository) Update(id string, updatedProfile models.Profile) (models.Profile, error) {
	var existingProfile models.Profile
	result := repo.DB.First(&existingProfile, "id = ?", id)
	if result.Error != nil {
		return existingProfile, result.Error
	}

	if err := repo.DB.Model(&existingProfile).Updates(updatedProfile).Error; err != nil {
		return existingProfile, err
	}
	return existingProfile, nil
}

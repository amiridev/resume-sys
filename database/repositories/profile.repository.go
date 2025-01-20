package repositories

import (
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type ProfileRepositoryInterface interface {
	Connection() *gorm.DB
	Show(userId string) (models.Profile, error)
	Create(profile models.Profile) (models.Profile, error)
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

func (repo *ProfileRepository) Show(userId string) (models.Profile, error) {
	var profile models.Profile

	result := repo.DB.First(&profile, "user_id = ?", userId)

	if result.Error != nil {
		return profile, result.Error
	}

	return profile, nil
}

func (repo *ProfileRepository) Create(profile models.Profile) (models.Profile, error) {
	result := repo.DB.Create(&profile)

	if result.Error != nil {
		return models.Profile{}, result.Error
	}

	return profile, nil
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

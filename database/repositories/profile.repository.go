package repositories

import (
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type ProfileRepositoryInterface interface {
	Connection() *gorm.DB
	Create(profile models.Profile) models.Profile
	List(userId string) []models.Profile
	Update(id string, profile models.Profile) models.Profile
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

func (repo *ProfileRepository) Create(Profile models.Profile) models.Profile {
	var newProfile models.Profile
	repo.DB.Create(&Profile).Scan(&newProfile)
	return newProfile
}

func (repo *ProfileRepository) List(userId string) []models.Profile {
	var profiles []models.Profile
	repo.DB.Table("profiles").Find(&profiles, "user_id = ?", userId)
	return profiles
}

func (repo *ProfileRepository) Update(id string, updatedProfile models.Profile) (models.Profile, error) {
	var existingProfile models.Profile
	result := repo.DB.First(&existingProfile, "id = ?", id)
	if result.Error != nil {
		return existingProfile, result.Error
	}
	repo.DB.Model(&existingProfile).Updates(updatedProfile)
	return existingProfile, nil
}

package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type EducationRepositoryInterface interface {
	Connection() *gorm.DB
	Create(education models.Education) models.Education
	List(userId string) []models.Education
	Update(id string, education models.Education) models.Education
	Delete(id string)
}

type EducationRepository struct {
	DB *gorm.DB
}

func NewEducationRepository() *EducationRepository {
	return &EducationRepository{
		DB: database.Connection(),
	}
}

func (repo *EducationRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *EducationRepository) Create(education models.Education) (models.Education, error) {
	var newEducation models.Education
	result := repo.DB.Create(&education)
	if result.Error != nil {

		return newEducation, result.Error
	}

	result.Scan(&newEducation)
	return newEducation, nil
}

func (repo *EducationRepository) List(userId string) []models.Education {
	var educations []models.Education
	repo.DB.Table("educations").Find(&educations, "user_id = ?", userId)

	return educations

}

func (repo *EducationRepository) Update(id string, updatedEducation models.Education) (models.Education, error) {
	var existingEducation models.Education
	result := repo.DB.First(&existingEducation, "id = ?", id)
	if result.Error != nil {
		return existingEducation, result.Error
	}
	repo.DB.Model(&existingEducation).Updates(updatedEducation)
	return existingEducation, nil
}

func (repo *EducationRepository) Delete(id string) error {
	result := repo.DB.Delete(&models.Education{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

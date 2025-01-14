package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	Connection() *gorm.DB
	Create(project models.Project) (models.Project, error)
	List(userId string) ([]models.Project, error)
	Show(projectId string) (models.Project, error)
	Update(id string, project models.Project) (models.Project, error)
	Delete(id string) error
}

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{
		DB: database.Connection(),
	}
}

func (repo *ProjectRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *ProjectRepository) Create(project models.Project) (models.Project, error) {
	result := repo.DB.Create(&project)
	if result.Error != nil {
		return models.Project{}, result.Error
	}
	return project, nil
}

func (repo *ProjectRepository) List(userId string) ([]models.Project, error) {
	var projects []models.Project
	result := repo.DB.Where("user_id = ?", userId).Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (repo *ProjectRepository) Show(projectId string) (models.Project, error) {
	var project models.Project

	result := repo.DB.First(&project, "id = ?", projectId)

	if result.Error != nil {
		return project, result.Error
	}

	return project, nil
}

func (repo *ProjectRepository) Update(id string, updatedProject models.Project) (models.Project, error) {
	var existingProject models.Project
	result := repo.DB.First(&existingProject, "id = ?", id)
	if result.Error != nil {
		return existingProject, result.Error
	}

	if err := repo.DB.Model(&existingProject).Updates(updatedProject).Error; err != nil {
		return existingProject, err
	}
	return existingProject, nil
}

func (repo *ProjectRepository) Delete(id string) error {
	result := repo.DB.Delete(&models.Project{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

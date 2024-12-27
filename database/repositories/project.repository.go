package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	Connection() *gorm.DB
	Create(project models.Project) models.Project
	List(userId string) []models.Project
	Update(id string, project models.Project) models.Project
	Delete(id string)
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

func (repo *ProjectRepository) Create(Project models.Project) models.Project {
	var newProject models.Project
	repo.DB.Create(&Project).Scan(&newProject)
	return newProject
}

func (repo *ProjectRepository) List(userId string) []models.Project {
	var projects []models.Project
	repo.DB.Table("projects").Find(&projects, "user_id = ?", userId)
	return projects
}

func (repo *ProjectRepository) Update(id string, updatedProject models.Project) (models.Project, error) {
	var existingProject models.Project
	result := repo.DB.First(&existingProject, "id = ?", id)
	if result.Error != nil {
		return existingProject, result.Error
	}
	repo.DB.Model(&existingProject).Updates(updatedProject)
	return existingProject, nil
}

func (repo *ProjectRepository) Delete(id string) error {
	result := repo.DB.Delete(&models.Project{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

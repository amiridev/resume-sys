package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type DocumentRepositoryInterface interface {
	Connection() *gorm.DB
	Create(document models.Document) (models.Document, error)
	List(userId string) ([]models.Document, error)
	Update(id string, document models.Document) (models.Document, error)
	Delete(id string) error
}

type DocumentRepository struct {
	DB *gorm.DB
}

func NewDocumentRepository() *DocumentRepository {
	return &DocumentRepository{
		DB: database.Connection(),
	}
}

func (repo *DocumentRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *DocumentRepository) Create(document models.Document) (models.Document, error) {
	result := repo.DB.Create(&document)
	if result.Error != nil {
		return models.Document{}, result.Error
	}
	return document, nil
}

func (repo *DocumentRepository) List(userId string) ([]models.Document, error) {
	var documents []models.Document
	result := repo.DB.Where("user_id = ?", userId).Find(&documents)
	if result.Error != nil {
		return nil, result.Error
	}
	return documents, nil
}

func (repo *DocumentRepository) Update(id string, updatedDocument models.Document) (models.Document, error) {
	var existingDocument models.Document
	result := repo.DB.First(&existingDocument, "id = ?", id)
	if result.Error != nil {
		return existingDocument, result.Error
	}

	if err := repo.DB.Model(&existingDocument).Updates(updatedDocument).Error; err != nil {
		return existingDocument, err
	}
	return existingDocument, nil
}

func (repo *DocumentRepository) Delete(id string) error {
	result := repo.DB.Delete(&models.Document{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

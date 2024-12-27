package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type DocumentRepositoryInterface interface {
	Connection() *gorm.DB
	Create(document models.Document) models.Document
	List(userId string) []models.Document
	Update(id string, document models.Document) models.Document
	Delete(id string)
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

func (repo *DocumentRepository) Create(Document models.Document) models.Document {
	var newDocument models.Document
	repo.DB.Create(&Document).Scan(&newDocument)
	return newDocument
}

func (repo *DocumentRepository) List(userId string) []models.Document {
	var documents []models.Document
	repo.DB.Table("documents").Find(&documents, "user_id = ?", userId)
	return documents
}

func (repo *DocumentRepository) Update(id string, updatedDocument models.Document) (models.Document, error) {
	var existingDocument models.Document
	result := repo.DB.First(&existingDocument, "id = ?", id)
	if result.Error != nil {
		return existingDocument, result.Error
	}
	repo.DB.Model(&existingDocument).Updates(updatedDocument)
	return existingDocument, nil
}

func (repo *DocumentRepository) Delete(id string) error {
	result := repo.DB.Delete(&models.Document{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

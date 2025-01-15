package documents

import (
	"resume-sys/core"
	"resume-sys/database/models"
	"resume-sys/database/repositories"
)

type DocumentsServiceInterface interface {
	List(userId string) ([]models.Document, core.Error)
	Show(DocumentId string) (models.Document, core.Error)
	Create(userId string, dto DocumentCreateDto) (models.Document, core.Error)
	Update(DocumentId string, dto DocumentUpdateDto) core.Error
	Delete(DocumentId string) core.Error
}

type DocumentsService struct {
	repository repositories.DocumentRepositoryInterface
}

func NewDocumentsService() *DocumentsService {
	return &DocumentsService{
		repository: repositories.NewDocumentRepository(),
	}
}

func (service *DocumentsService) List(userId string) ([]models.Document, core.Error) {
	documents, err := service.repository.List(userId)

	if err != nil {
		return documents, core.Error{"reason": "Something wrong!"}
	}

	return documents, nil
}

func (service *DocumentsService) Show(documentId string) (models.Document, core.Error) {
	document, err := service.repository.Show(documentId)

	if err != nil {
		return document, core.Error{"reason": "Document not found"}
	}

	return document, nil
}

func (service *DocumentsService) Create(userId string, dto DocumentCreateDto) (models.Document, core.Error) {

	document, err := service.repository.Create(models.Document{
		UserID:          userId,
		Title:           dto.Title,
		Category:        dto.Category,
		PublicationDate: dto.PublicationDate,
		Status:          dto.Status,
		Ratings:         dto.Ratings,
		Description:     dto.Description,
	})

	if err != nil {
		return document, core.Error{"reason": "Something wrong!"}
	}

	return document, nil
}

func (service *DocumentsService) Update(documentId string, dto DocumentUpdateDto) core.Error {
	_, err := service.repository.Update(documentId, models.Document{
		Title:           dto.Title,
		Category:        dto.Category,
		PublicationDate: dto.PublicationDate,
		Status:          dto.Status,
		Ratings:         dto.Ratings,
		Description:     dto.Description,
	})

	if err != nil {
		return core.Error{"reason": "Something wrong!"}
	}

	return nil
}

func (service *DocumentsService) Delete(documentId string) core.Error {
	err := service.repository.Delete(documentId)

	if err != nil {
		return core.Error{"reason": "Document not found"}
	}

	return nil
}

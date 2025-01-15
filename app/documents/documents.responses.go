package documents

import (
	"resume-sys/database/models"
	"time"
)

type ResponseType map[string]any

type Document struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	Title           string    `json:"title"`
	Category        string    `json:"category"`
	Ratings         int       `json:"ratings"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
	PublicationDate time.Time `json:"publicationDate"`
}

// one document response

func DocumentTransform(document models.Document) Document {
	return Document{
		ID:              document.ID,
		UserID:          document.UserID,
		Title:           document.Title,
		Category:        document.Category,
		Ratings:         document.Ratings,
		Description:     document.Description,
		Status:          document.Status,
		PublicationDate: document.PublicationDate,
	}
}

type DocumentResponseType struct {
	Documents Document `json:"Document"`
}

func DocumentResponse(document models.Document) ResponseType {
	return ResponseType{
		"document": DocumentTransform(document),
	}
}

// list of document response

func DocumentsTransform(documents []models.Document) []Document {
	var data = []Document{}

	for _, document := range documents {
		data = append(data, DocumentTransform(document))
	}

	return data
}

type DocumentsResponseType struct {
	Documents []Document `json:"documents"`
}

func DocumentsResponse(documents []models.Document) ResponseType {
	return ResponseType{
		"documents": DocumentsTransform(documents),
	}
}

package documents

import "time"

type DocumentCreateDto struct {
	Title           string    `json:"title" binding:"required,min=2,max=250"`
	Category        string    `json:"category" binding:"required"`
	Ratings         int       `json:"ratings" binding:"gte=0,lte=5"`
	Description     string    `json:"description" binding:"required"`
	Status          string    `json:"status" binding:"required"`
	PublicationDate time.Time `json:"publication_date" binding:"required,datetime=2006-01-02"`
}

type DocumentUpdateDto struct {
	Title           string    `json:"title" binding:"required,min=2,max=250"`
	Category        string    `json:"category" binding:"required"`
	Ratings         int       `json:"ratings" binding:"gte=0,lte=5"`
	Description     string    `json:"description" binding:"required"`
	Status          string    `json:"status" binding:"required"`
	PublicationDate time.Time `json:"publication_date" binding:"required,datetime=2006-01-02"`
}

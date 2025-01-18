package educations

import "time"

type EducationCreateDto struct {
	Title       string    `json:"title" binding:"required,min=2,max=250"`
	Field       string    `json:"field" binding:"required,min=2,max=250"`
	Description string    `json:"description" binding:"required"`
	StartedAt   time.Time `json:"started_at" binding:"required"`
	EndedAt     time.Time `json:"ended_at" binding:"required"`
}

type EducationUpdateDto struct {
	Title       string    `json:"title" binding:"required,min=2,max=250"`
	Field       string    `json:"field" binding:"required,min=2,max=250"`
	Description string    `json:"description" binding:"required"`
	StartedAt   time.Time `json:"started_at" binding:"required"`
	EndedAt     time.Time `json:"ended_at" binding:"required"`
}

package experiences

import "time"

type ExperienceCreateDto struct {
	Title     string    `json:"title" binding:"required,min=2,max=250"`
	Company   string    `json:"company" binding:"required,min=2,max=250"`
	Status    string    `json:"status" binding:"required"`
	StartedAt time.Time `json:"started_at" binding:"required"`
	EndedAt   time.Time `json:"ended_at" binding:"required"`
}

type ExperienceUpdateDto struct {
	Title     string    `json:"title" binding:"required,min=2,max=250"`
	Company   string    `json:"company" binding:"required,min=2,max=250"`
	Status    string    `json:"status" binding:"required"`
	StartedAt time.Time `json:"started_at" binding:"required"`
	EndedAt   time.Time `json:"ended_at" binding:"required"`
}

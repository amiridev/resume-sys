package experiences

type ExperienceCreateDto struct {
	Title     string `json:"title" binding:"required,min=2,max=250"`
	Company   string `json:"company" binding:"required,min=2,max=250"`
	Status    string `json:"status" binding:"required"`
	StartedAt string `json:"started_at" binding:"required"`
	EndedAt   string `json:"ended_at" binding:"required"`
}

type ExperienceUpdateDto struct {
	Title     string `json:"title" binding:"required,min=2,max=250"`
	Company   string `json:"company" binding:"required,min=2,max=250"`
	Status    string `json:"status" binding:"required"`
	StartedAt string `json:"started_at" binding:"required"`
	EndedAt   string `json:"ended_at" binding:"required"`
}

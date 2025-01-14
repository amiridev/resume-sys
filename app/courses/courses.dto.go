package courses

type CoursCreateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=250"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type CoursUpdateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=250"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

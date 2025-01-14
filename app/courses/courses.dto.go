package courses

type CourseCreateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=250"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type CourseUpdateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=250"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

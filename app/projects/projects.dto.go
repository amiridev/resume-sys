package projects

type ProjectCreateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=250"`
	ImageUrl    string `json:"image_url" binding:"required,min=10,max=250"`
	Link        string `json:"link" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type ProjectUpdateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=250"`
	ImageUrl    string `json:"image_url" binding:"required,min=10,max=250"`
	Link        string `json:"link" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

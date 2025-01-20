package skills

type SkillCreateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=250"`
	Level       string `json:"level" binding:"required,min=2,max=250"`
	ImageUrl    string `json:"image_url" binding:"required,min=10,max=250"`
	Description string `json:"description" binding:"required"`
}

type SkillUpdateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=250"`
	Level       string `json:"level" binding:"required,min=2,max=250"`
	ImageUrl    string `json:"image_url" binding:"required,min=10,max=250"`
	Description string `json:"description" binding:"required"`
}

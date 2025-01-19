package profiles

type ProfileCreateDto struct {
	FullName    string   `json:"full_name" binding:"required,min=2,max=100"`
	UserName    string   `json:"user_name" binding:"required,min=2,max=50"`
	Mobile      string   `json:"mobile" binding:"required"`
	AvatarURL   string   `json:"avatar_url" binding:"required,url"`
	Gender      string   `json:"gender" binding:"required"`
	DateOfBirth string   `json:"date_of_birth" binding:"required"`
	City        string   `json:"city" binding:"required"`
	Country     string   `json:"country" binding:"required"`
	Language    string   `json:"language" binding:"required"`
	SocialLinks []string `json:"social_links" binding:"required"`
	Status      string   `json:"status" binding:"required"`
}

type ProfileUpdateDto struct {
	FullName    string   `json:"full_name" binding:"omitempty,min=2,max=100"`
	UserName    string   `json:"user_name" binding:"omitempty,min=2,max=50"`
	Mobile      string   `json:"mobile" binding:"omitempty"`
	AvatarURL   string   `json:"avatar_url" binding:"omitempty,url"`
	Gender      string   `json:"gender" binding:"omitempty"`
	DateOfBirth string   `json:"date_of_birth" binding:"omitempty"`
	City        string   `json:"city" binding:"omitempty"`
	Country     string   `json:"country" binding:"omitempty"`
	Language    string   `json:"language" binding:"omitempty"`
	SocialLinks []string `json:"social_links" binding:"omitempty"`
	Status      string   `json:"status" binding:"omitempty"`
}

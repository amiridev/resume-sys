package profiles

import "time"

type ProfileCreateDto struct {
	UserName    string    `json:"username" binding:"required,min=2,max=100"`
	AvatarUrl   string    `json:"avatar_url" binding:"required"`
	Gender      string    `json:"gender" binding:"omitempty"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	City        string    `json:"city" binding:"omitempty"`
	Country     string    `json:"country" binding:"omitempty"`
	Language    string    `json:"language" binding:"omitempty"`
	SocialLinks string    `json:"social_links" binding:"omitempty"`
	Status      string    `json:"status" binding:"omitempty"`
}

type ProfileUpdateDto struct {
	UserName    string    `json:"username" binding:"required,min=2,max=100"`
	AvatarUrl   string    `json:"avatar_url" binding:"required"`
	Gender      string    `json:"gender" binding:"omitempty"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	City        string    `json:"city" binding:"omitempty"`
	Country     string    `json:"country" binding:"omitempty"`
	Language    string    `json:"language" binding:"omitempty"`
	SocialLinks string    `json:"social_links" binding:"omitempty"`
	Status      string    `json:"status" binding:"omitempty"`
}

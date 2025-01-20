package profiles

import (
	"resume-sys/database/models"
	"time"
)

type ProfileResponseType map[string]any

type Profile struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Mobile      string `json:"mobile"`
	AvatarURL   string `json:"avatar_url"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Language    string `json:"language"`
	SocialLinks string `json:"social_links"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ProfileTransform(profile models.Profile) Profile {
	return Profile{
		ID:          profile.ID,
		UserID:      profile.UserID,
		UserName:    profile.UserName,
		AvatarURL:   profile.AvatarUrl,
		Gender:      profile.Gender,
		DateOfBirth: profile.DateOfBirth.Format(time.RFC3339),
		City:        profile.City,
		Country:     profile.Country,
		Language:    profile.Language,
		SocialLinks: profile.SocialLinks,
		Status:      profile.Status,
		CreatedAt:   profile.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   profile.UpdatedAt.Format(time.RFC3339),
	}
}

func ProfilesTransform(profiles []models.Profile) []Profile {
	var data []Profile
	for _, profile := range profiles {
		data = append(data, ProfileTransform(profile))
	}
	return data
}

func ProfileResponse(profile models.Profile) ProfileResponseType {
	return ProfileResponseType{
		"profile": ProfileTransform(profile),
	}
}

func ProfilesResponse(profiles []models.Profile) ProfileResponseType {
	return ProfileResponseType{
		"profiles": ProfilesTransform(profiles),
	}
}

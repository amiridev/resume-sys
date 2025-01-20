package profiles

import (
	"resume-sys/core"
	"resume-sys/database/models"
	"resume-sys/database/repositories"
)

type ProfileServiceInterface interface {
	GetProfile(userID string) (models.Profile, core.Error)
	CreateProfile(dto ProfileCreateDto, userID string) (models.Profile, core.Error)
	UpdateProfile(dto ProfileUpdateDto, userID string) (models.Profile, core.Error)
}

type ProfileService struct {
	repository repositories.ProfileRepositoryInterface
}

func NewProfileService() ProfileServiceInterface {
	return &ProfileService{
		repository: repositories.NewProfileRepository(),
	}
}

// GetProfile retrieves a user's profile by user ID.
func (s *ProfileService) GetProfile(userID string) (models.Profile, core.Error) {
	profile, err := s.repository.Show(userID)

	if err != nil {
		return profile, core.Error{"reason": "profile not found"}
	}

	return profile, nil
}

// CreateProfile creates a new profile for the user.
func (s *ProfileService) CreateProfile(dto ProfileCreateDto, userID string) (models.Profile, core.Error) {
	profile := models.Profile{
		UserID:      userID,
		UserName:    dto.UserName,
		AvatarUrl:   dto.AvatarUrl,
		Gender:      dto.Gender,
		DateOfBirth: dto.DateOfBirth,
		City:        dto.City,
		Country:     dto.Country,
		Language:    dto.Language,
		SocialLinks: dto.SocialLinks,
		Status:      dto.Status,
	}

	profile, err := s.repository.Create(profile)

	if err != nil {
		return profile, core.Error{"reason": "Something wrong!"}
	}

	return profile, nil
}

// UpdateProfile updates an existing profile with the provided data.
func (s *ProfileService) UpdateProfile(dto ProfileUpdateDto, userID string) (models.Profile, core.Error) {
	profile, err := s.repository.Show(userID)

	if err != nil {
		return profile, core.Error{"reason": "Something wrong!"}
	}

	if dto.UserName != "" {
		profile.UserName = dto.UserName
	}
	if dto.AvatarUrl != "" {
		profile.AvatarUrl = dto.AvatarUrl
	}
	if dto.Gender != "" {
		profile.Gender = dto.Gender
	}
	// if dto.DateOfBirth != "" {
	// 	profile.DateOfBirth = dto.DateOfBirth
	// }
	if dto.City != "" {
		profile.City = dto.City
	}
	if dto.Country != "" {
		profile.Country = dto.Country
	}
	if dto.Language != "" {
		profile.Language = dto.Language
	}
	if len(dto.SocialLinks) > 0 {
		profile.SocialLinks = dto.SocialLinks
	}
	if dto.Status != "" {
		profile.Status = dto.Status
	}

	_, err = s.repository.Update(profile.ID, profile)

	if err != nil {
		return profile, core.Error{"reason": "Something wrong!"}
	}

	return profile, nil
}

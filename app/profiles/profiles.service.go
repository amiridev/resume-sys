package profiles

import (
	"errors"
	"resume-sys/core"
)

type ProfileServiceInterface interface {
	GetProfile(userID string) (*Profile, error)
	CreateProfile(dto ProfileCreateDto, userID string) (*Profile, error)
	UpdateProfile(dto ProfileUpdateDto, userID string) (*Profile, error)
	GetUserDetails(userID string) (*UserDetails, error)
}

type ProfileService struct{}

func NewProfileService() ProfileServiceInterface {
	return &ProfileService{}
}

// GetProfile retrieves a user's profile by user ID.
func (s *ProfileService) GetProfile(userID string) (*Profile, error) {
	profile, err := core.GetProfileByUserID(userID)
	if err != nil {
		return nil, errors.New("profile not found")
	}
	return profile, nil
}

// CreateProfile creates a new profile for the user.
func (s *ProfileService) CreateProfile(dto ProfileCreateDto, userID string) (*Profile, error) {
	profile := &Profile{
		UserID:      userID,
		FullName:    dto.FullName,
		UserName:    dto.UserName,
		Mobile:      dto.Mobile,
		AvatarURL:   dto.AvatarURL,
		Gender:      dto.Gender,
		DateOfBirth: dto.DateOfBirth,
		City:        dto.City,
		Country:     dto.Country,
		Language:    dto.Language,
		SocialLinks: dto.SocialLinks,
		Status:      dto.Status,
	}

	err := core.SaveProfile(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// UpdateProfile updates an existing profile with the provided data.
func (s *ProfileService) UpdateProfile(dto ProfileUpdateDto, userID string) (*Profile, error) {
	profile, err := core.GetProfileByUserID(userID)
	if err != nil {
		return nil, errors.New("profile not found")
	}

	if dto.FullName != "" {
		profile.FullName = dto.FullName
	}
	if dto.UserName != "" {
		profile.UserName = dto.UserName
	}
	if dto.Mobile != "" {
		profile.Mobile = dto.Mobile
	}
	if dto.AvatarURL != "" {
		profile.AvatarURL = dto.AvatarURL
	}
	if dto.Gender != "" {
		profile.Gender = dto.Gender
	}
	if dto.DateOfBirth != "" {
		profile.DateOfBirth = dto.DateOfBirth
	}
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

	err = core.SaveProfile(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// GetUserDetails retrieves all details of a user including profile, courses, projects, etc.
func (s *ProfileService) GetUserDetails(userID string) (*UserDetails, error) {
	profile, err := core.GetProfileByUserID(userID)
	if err != nil {
		return nil, errors.New("profile not found")
	}

	courses, err := core.GetCoursesByUserID(userID)
	if err != nil {
		return nil, err
	}

	projects, err := core.GetProjectsByUserID(userID)
	if err != nil {
		return nil, err
	}

	experiences, err := core.GetExperiencesByUserID(userID)
	if err != nil {
		return nil, err
	}

	userDetails := &UserDetails{
		Profile:     profile,
		Courses:     courses,
		Projects:    projects,
		Experiences: experiences,
	}

	return userDetails, nil
}

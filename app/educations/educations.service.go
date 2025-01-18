package educations

import (
	"resume-sys/core"
	"resume-sys/database/models"
	"resume-sys/database/repositories"
)

type EducationsServiceInterface interface {
	List(userId string) []models.Education
	Show(educationId string) (models.Education, core.Error)
	Create(userId string, dto EducationCreateDto) (models.Education, core.Error)
	Update(educationId string, dto EducationUpdateDto) core.Error
	Delete(educationId string) core.Error
}

type EducationsService struct {
	repository repositories.EducationRepositoryInterface
}

func NewEducationsService() *EducationsService {
	return &EducationsService{
		repository: repositories.NewEducationRepository(),
	}
}

func (service *EducationsService) List(userId string) []models.Education {
	educations := service.repository.List(userId)
	return educations
}

func (service *EducationsService) Show(educationId string) (models.Education, core.Error) {
	education, err := service.repository.Show(educationId)

	if err != nil {
		return education, core.Error{"reason": "Education not found"}
	}

	return education, nil
}

func (service *EducationsService) Create(userId string, dto EducationCreateDto) (models.Education, core.Error) {

	education, err := service.repository.Create(models.Education{
		UserID:      userId,
		Title:       dto.Title,
		Field:       dto.Field,
		Description: dto.Description,
		StartedAt:   dto.StartedAt,
		EndedAt:     dto.EndedAt,
	})

	if err != nil {
		return education, core.Error{"reason": "Something wrong!"}
	}

	return education, nil
}

func (service *EducationsService) Update(educationId string, dto EducationUpdateDto) core.Error {
	_, err := service.repository.Update(educationId, models.Education{
		Title:       dto.Title,
		Field:       dto.Field,
		Description: dto.Description,
		StartedAt:   dto.StartedAt,
		EndedAt:     dto.EndedAt,
	})

	if err != nil {
		return core.Error{"reason": "Something wrong!"}
	}

	return nil
}

func (service *EducationsService) Delete(educationId string) core.Error {
	err := service.repository.Delete(educationId)

	if err != nil {
		return core.Error{"reason": "Education not found"}
	}

	return nil
}

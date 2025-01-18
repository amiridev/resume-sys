package experiences

import (
	"resume-sys/core"
	"resume-sys/database/models"
	"resume-sys/database/repositories"
)

type ExperiencesServiceInterface interface {
	List(userId string) ([]models.Experience, core.Error)
	Show(experienceId string) (models.Experience, core.Error)
	Create(userId string, dto ExperienceCreateDto) (models.Experience, core.Error)
	Update(experienceId string, dto ExperienceUpdateDto) core.Error
	Delete(experienceId string) core.Error
}

type ExperiencesService struct {
	repository repositories.ExperienceRepositoryInterface
}

func NewExperiencesService() *ExperiencesService {
	return &ExperiencesService{
		repository: repositories.NewExperienceRepository(),
	}
}

func (service *ExperiencesService) List(userId string) ([]models.Experience, core.Error) {
	experiences, err := service.repository.List(userId)

	if err != nil {
		return experiences, core.Error{"reason": "Something wrong!"}
	}

	return experiences, nil
}

func (service *ExperiencesService) Show(experienceId string) (models.Experience, core.Error) {
	experience, err := service.repository.Show(experienceId)

	if err != nil {
		return experience, core.Error{"reason": "Experience not found"}
	}

	return experience, nil
}

func (service *ExperiencesService) Create(userId string, dto ExperienceCreateDto) (models.Experience, core.Error) {

	experience, err := service.repository.Create(models.Experience{
		UserID:    userId,
		Title:     dto.Title,
		Company:   dto.Company,
		Status:    dto.Status,
		StartedAt: dto.StartedAt,
		EndedAt:   dto.EndedAt,
	})

	if err != nil {
		return experience, core.Error{"reason": "Something wrong!"}
	}

	return experience, nil
}

func (service *ExperiencesService) Update(experienceId string, dto ExperienceUpdateDto) core.Error {
	_, err := service.repository.Update(experienceId, models.Experience{
		Title:     dto.Title,
		Company:   dto.Company,
		Status:    dto.Status,
		StartedAt: dto.StartedAt,
		EndedAt:   dto.EndedAt,
	})

	if err != nil {
		return core.Error{"reason": "Something wrong!"}
	}

	return nil
}

func (service *ExperiencesService) Delete(experienceId string) core.Error {
	err := service.repository.Delete(experienceId)

	if err != nil {
		return core.Error{"reason": "Experience not found"}
	}

	return nil
}

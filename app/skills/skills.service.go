package skills

import (
	"resume-sys/core"
	"resume-sys/database/models"
	"resume-sys/database/repositories"
)

type SkillsServiceInterface interface {
	List(userId string) []models.Skill
	Show(skillId string) (models.Skill, core.Error)
	Create(userId string, dto SkillCreateDto) models.Skill
	Update(skillId string, dto SkillUpdateDto) core.Error
	Delete(skillId string) core.Error
}

type SkillsService struct {
	repository repositories.SkillRepositoryInterface
}

func NewSkillsService() *SkillsService {
	return &SkillsService{
		repository: repositories.NewSkillRepository(),
	}
}

func (service *SkillsService) List(userId string) []models.Skill {
	skills := service.repository.List(userId)

	return skills
}

func (service *SkillsService) Show(skillId string) (models.Skill, core.Error) {
	skill, err := service.repository.Show(skillId)

	if err != nil {
		return skill, core.Error{"reason": "Skill not found"}
	}

	return skill, nil
}

func (service *SkillsService) Create(userId string, dto SkillCreateDto) models.Skill {

	skill := service.repository.Create(models.Skill{
		UserID:      userId,
		Title:       dto.Title,
		Level:       dto.Level,
		ImageUrl:    dto.ImageUrl,
		Description: dto.Description,
	})

	return skill
}

func (service *SkillsService) Update(skillId string, dto SkillUpdateDto) core.Error {
	_, err := service.repository.Update(skillId, models.Skill{
		Title:       dto.Title,
		Level:       dto.Level,
		ImageUrl:    dto.ImageUrl,
		Description: dto.Description,
	})

	if err != nil {
		return core.Error{"reason": "Something wrong!"}

	}

	return nil
}

func (service *SkillsService) Delete(skillId string) core.Error {
	err := service.repository.Delete(skillId)

	if err != nil {
		return core.Error{"reason": "Skill not found"}
	}

	return nil
}

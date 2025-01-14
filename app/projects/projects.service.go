package projects

import (
	"resume-sys/core"
	"resume-sys/database/models"
	"resume-sys/database/repositories"
)

type ProjectsServiceInterface interface {
	List(userId string) ([]models.Project, core.Error)
	Show(projectId string) (models.Project, core.Error)
	Create(userId string, dto ProjectCreateDto) (models.Project, core.Error)
	Update(projectId string, dto ProjectUpdateDto) core.Error
	Delete(projectId string) core.Error
}

type ProjectsService struct {
	repository repositories.ProjectRepositoryInterface
}

func NewProjectsService() *ProjectsService {
	return &ProjectsService{
		repository: repositories.NewProjectRepository(),
	}
}

func (service *ProjectsService) List(userId string) ([]models.Project, core.Error) {
	projects, err := service.repository.List(userId)

	if err != nil {
		return projects, core.Error{"reason": "Something wrong!"}
	}

	return projects, nil
}

func (service *ProjectsService) Show(projectId string) (models.Project, core.Error) {
	project, err := service.repository.Show(projectId)

	if err != nil {
		return project, core.Error{"reason": "Project not found"}
	}

	return project, nil
}

func (service *ProjectsService) Create(userId string, dto ProjectCreateDto) (models.Project, core.Error) {

	project, err := service.repository.Create(models.Project{
		UserID:      userId,
		Title:       dto.Title,
		ImageUrl:    dto.ImageUrl,
		Link:        dto.Link,
		Description: dto.Description,
		Status:      dto.Status,
	})

	if err != nil {
		return project, core.Error{"reason": "Something wrong!"}
	}

	return project, nil
}

func (service *ProjectsService) Update(projectId string, dto ProjectUpdateDto) core.Error {
	_, err := service.repository.Update(projectId, models.Project{
		Title:       dto.Title,
		ImageUrl:    dto.ImageUrl,
		Link:        dto.Link,
		Description: dto.Description,
		Status:      dto.Status,
	})

	if err != nil {
		return core.Error{"reason": "Something wrong!"}
	}

	return nil
}

func (service *ProjectsService) Delete(projectId string) core.Error {
	err := service.repository.Delete(projectId)

	if err != nil {
		return core.Error{"reason": "Project not found"}
	}

	return nil
}

package courses

import (
	"resume-sys/core"
	"resume-sys/database/models"
	"resume-sys/database/repositories"
)

type CoursesServiceInterface interface {
	List(UserID string) ([]models.Course, core.Error)
	Show(courseId string) (models.Course, core.Error)
	Create(userId string, dto CourseCreateDto) (models.Course, core.Error)
	Update(courseId string, dto CourseUpdateDto) core.Error
	Delete(courseId string) core.Error
}

type CoursesService struct {
	repository repositories.CourseRepositoryInterface
}

func NewCoursesService() *CoursesService {
	return &CoursesService{
		repository: repositories.NewCourseRepository(),
	}
}

func (service *CoursesService) List(userId string) ([]models.Course, core.Error) {
	courses, err := service.repository.List(userId)

	if err != nil {
		return courses, core.Error{"reason": "Something wrong!"}
	}

	return courses, nil
}

func (service *CoursesService) Show(courseId string) (models.Course, core.Error) {
	course, err := service.repository.Show(courseId)

	if err != nil {
		return course, core.Error{"reason": "Course not found"}
	}

	return course, nil
}

func (service *CoursesService) Create(userId string, dto CourseCreateDto) (models.Course, core.Error) {
	course, err := service.repository.Create(models.Course{
		UserID:      userId,
		Title:       dto.Title,
		Category:    dto.Category,
		Description: dto.Description,
		Status:      dto.Status,
	})

	if err != nil {
		return course, core.Error{"reason": "Something wrong!"}

	}

	return course, nil
}

func (service *CoursesService) Update(courseId string, dto CourseUpdateDto) core.Error {
	_, err := service.repository.Update(courseId, models.Course{
		Title:       dto.Title,
		Category:    dto.Category,
		Description: dto.Description,
		Status:      dto.Status,
	})

	if err != nil {
		return core.Error{"reason": "Something wrong!"}
	}

	return nil
}

func (service *CoursesService) Delete(courseId string) core.Error {
	err := service.repository.Delete(courseId)

	if err != nil {
		return core.Error{"reason": "Course not found"}
	}

	return nil
}

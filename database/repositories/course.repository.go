package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type CourseRepositoryInterface interface {
	Connection() *gorm.DB
	Create(course models.Course) (models.Course, error)
	List(userId string) ([]models.Course, error)
	Update(id string, course models.Course) (models.Course, error)
	Delete(id string) error
}

type CourseRepository struct {
	DB *gorm.DB
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{
		DB: database.Connection(),
	}
}

func (repo *CourseRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *CourseRepository) Create(course models.Course) (models.Course, error) {
	result := repo.DB.Create(&course)
	if result.Error != nil {
		return models.Course{}, result.Error
	}
	return course, nil
}

func (repo *CourseRepository) List(userId string) ([]models.Course, error) {
	var courses []models.Course
	result := repo.DB.Table("courses").Where("user_id = ?", userId).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func (repo *CourseRepository) Update(id string, updatedCourse models.Course) (models.Course, error) {
	var existingCourse models.Course
	result := repo.DB.First(&existingCourse, "id = ?", id)
	if result.Error != nil {
		return existingCourse, result.Error
	}
	repo.DB.Model(&existingCourse).Updates(updatedCourse)
	return existingCourse, nil
}

func (repo *CourseRepository) Delete(id string) error {
	result := repo.DB.Delete(&models.Course{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

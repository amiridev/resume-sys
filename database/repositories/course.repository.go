package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type CourseRepositoryInterface interface {
	Connection() *gorm.DB
	Create(course models.Course) models.Course
	List(userId string) []models.Course
	Update(id string, corse models.Course) models.Course
	Delete(id string)
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

func (repo *CourseRepository) Create(Course models.Course) models.Course {
	var newCourse models.Course
	repo.DB.Create(&Course).Scan(&newCourse)
	return newCourse
}

func (repo *CourseRepository) List(userId string) []models.Course {
	var courses []models.Course
	repo.DB.Table("courses").Find(&courses, "userId = ?", userId)
	return courses

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

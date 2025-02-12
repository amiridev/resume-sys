package repositories

import (
	"fmt"
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type SkillRepositoryInterface interface {
	Connection() *gorm.DB
	Create(skill models.Skill) models.Skill
	List(userId string) []models.Skill
	Show(skillId string) (models.Skill, error)
	Update(id string, skill models.Skill) (models.Skill, error)
	Delete(id string) error
}

type SkillRepository struct {
	DB *gorm.DB
}

func NewSkillRepository() *SkillRepository {
	return &SkillRepository{
		DB: database.Connection(),
	}
}

func (repo *SkillRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *SkillRepository) Create(skill models.Skill) models.Skill {
	var newSkill models.Skill
	repo.DB.Create(&skill).Scan(&newSkill)
	return newSkill
}

func (repo *SkillRepository) List(userId string) []models.Skill {
	var skills []models.Skill
	repo.DB.Table("skills").Find(&skills, "user_id = ?", userId)
	return skills
}

func (repo *SkillRepository) Show(skillId string) (models.Skill, error) {
	var skill models.Skill

	result := repo.DB.First(&skill, "id = ?", skillId)

	if result.Error != nil {
		return skill, result.Error
	}

	return skill, nil
}

func (repo *SkillRepository) Update(id string, updatedSkill models.Skill) (models.Skill, error) {
	var existingSkill models.Skill
	result := repo.DB.First(&existingSkill, "id = ?", id)
	if result.Error != nil {
		return existingSkill, result.Error // خطا اگر رکورد پیدا نشود
	}
	repo.DB.Model(&existingSkill).Updates(updatedSkill)
	return existingSkill, nil
}

func (repo *SkillRepository) Delete(id string) error {
	result := repo.DB.Delete(&models.Skill{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found") // اگر رکوردی پیدا نشود
	}
	return nil
}

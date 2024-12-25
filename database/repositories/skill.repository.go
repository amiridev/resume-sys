package repositories

import (
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type SkillRepositoryInterface interface {
	Connection() *gorm.DB
	Create(skill models.Skill) models.Skill
	List(userId string) []models.Skill
	Update(id string, skill models.Skill) models.Skill
	Delete(id string)
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

func (repo *SkillRepository) Create(Skill models.Skill) models.Skill {
	var newSkill models.Skill
	repo.DB.Create(&Skill).Scan(&newSkill)
	return newSkill
}

func (repo *SkillRepository) List(userId string) []models.Skill {
	var skills []models.Skill
	repo.DB.Table("skills").Find(&skills, "user_id = ?", userId)
	return skills
}

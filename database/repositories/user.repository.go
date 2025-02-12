package repositories

import (
	"resume-sys/database"
	"resume-sys/database/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Connection() *gorm.DB
	Create(user models.User) models.User
	FindByMobile(mobile string) models.User
	FindByID(id string) models.User
	FindAll(exceptUserId string) []models.User
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{ //dependenci ingection
		DB: database.Connection(),
	}
}

func (repo *UserRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *UserRepository) Create(user models.User) models.User {
	var newUser models.User
	repo.DB.Create(&user).Scan(&newUser)
	return newUser
}

func (repo *UserRepository) FindByMobile(mobile string) models.User {
	var user models.User
	repo.DB.First(&user, "mobile = ?", mobile)
	return user
}

func (repo *UserRepository) FindByID(id string) models.User {
	var user models.User
	repo.DB.First(&user, "id = ?", id)
	return user
}

func (repo *UserRepository) FindAll(exceptUserId string) []models.User {
	var users []models.User
	repo.DB.Table("users").Not("id = ?", exceptUserId).Find(&users)
	return users
}

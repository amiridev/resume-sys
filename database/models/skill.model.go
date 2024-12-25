package models

import (
	"resume-sys/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type Skill struct {
	ID          string    `gorm:"type:varchar(40);primaryKey"`
	UserID      string    `gorm:"type:varchar(40);not null"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Level       string    `gorm:"type:varchar(50)"`
	ImageUrl    string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignkey:UserID"`
}

func (s *Skill) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = helpers.CUID()

	if s.Level == "" {
		s.Level = "beginner"
	}

	return
}

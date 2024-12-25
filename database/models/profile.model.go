package models

import (
	"resume-sys/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID          string    `gorm:"type:varchar(40);primaryKey"`
	UserID      string    `gorm:"type:varchar(40);not null"`
	UserName    string    `gorm:"type:varchar(100);not null"`
	AvatarUrl   string    `gorm:"type:varchar(255)"`
	Gender      string    `gorm:"type:varchar(10)"`
	DateOfBirth time.Time `gorm:"not null"`
	City        string    `gorm:"type:varchar(100)"`
	Country     string    `gorm:"type:varchar(100)"`
	Language    string    `gorm:"type:varchar(50)"`
	SocialLinks string    `gorm:"type:text"`
	Status      string    `gorm:"type:varchar(50)"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignkey:UserID"`
}

func (p *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = helpers.CUID()

	return
}

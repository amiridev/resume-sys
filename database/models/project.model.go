package models

import (
	"resume-sys/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          string    `gorm:"type:varchar(40);primaryKey"`
	UserID      string    `gorm:"type:varchar(40);not null"`
	Title       string    `gorm:"type:varchar(255);not null"`
	ImageUrl    string    `gorm:"type:varchar(255)"`
	Link        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:text"`
	Status      string    `gorm:"type:varchar(50)"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignkey:UserID"`
}

func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {

	p.ID = helpers.CUID()

	return
}

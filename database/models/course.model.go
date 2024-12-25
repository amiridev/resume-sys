package models

import (
	"resume-sys/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID          string    `gorm:"type:varchar(40);primaryKey"`
	UserID      string    `gorm:"type:varchar(40);not null" `
	Title       string    `gorm:"type:varchar(255);not null" `
	Category    string    `gorm:"type:varchar(100);not null" `
	Description string    `gorm:"type:text" `
	Status      string    `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime" `
	UpdatedAt   time.Time `gorm:"autoUpdateTime" `

	User User `gorm:"foreignKey:UserID"`
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = helpers.CUID()

	return
}

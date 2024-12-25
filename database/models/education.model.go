package models

import (
	"resume-sys/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type Education struct {
	ID          string    `gorm:"type:varchar(40);primaryKey" `
	UserID      string    `gorm:"type:varchar(40);not null" `
	Title       string    `gorm:"type:varchar(255);not null" `
	Field       string    `gorm:"type:varchar(100);not null" `
	Description string    `gorm:"type:text" `
	StartedAt   time.Time `gorm:"not null" `
	EndedAt     time.Time `gorm:"" `
	CreatedAt   time.Time `gorm:"autoCreateTime" `
	UpdatedAt   time.Time `gorm:"autoUpdateTime" `

	User User `gorm:"foreignkey:UserID"`
}

func (e *Education) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = helpers.CUID()

	return
}

package models

import (
	"resume-sys/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	ID        string    `gorm:"type:varchar(40);primaryKey" `
	UserID    string    `gorm:"type:varchar(40);not null" `
	Title     string    `gorm:"type:varchar(255);not null" `
	Company   string    `gorm:"type:varchar(255);not null" `
	Status    string    `gorm:"type:varchar(50);not null" `
	StartedAt time.Time `gorm:"not null" `
	EndedAt   time.Time `gorm:""`
	CreatedAt time.Time `gorm:"autoCreateTime" `
	UpdatedAt time.Time `gorm:"autoUpdateTime" `

	User User `gorm:"foreignkey:UserID"`
}

func (e *Experience) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = helpers.CUID()

	return
}

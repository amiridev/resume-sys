package models

import (
	"resume-sys/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type Document struct {
	ID              string    `gorm:"type:varchar(40);primaryKey"`
	UserID          string    `gorm:"type:varchar(40);not null"`
	Title           string    `gorm:"type:varchar(255);not null" `
	Category        string    `gorm:"type:varchar(100);not null" `
	Ratings         float64   `gorm:"type:decimal(3,2);default:0" `
	Description     string    `gorm:"type:text" `
	Status          string    `gorm:"type:varchar(50);not null" `
	PublicationDate time.Time `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime" `
	UpdatedAt       time.Time `gorm:"autoUpdateTime" `

	User User `gorm:"foreignKey:UserID"`
}

func (d *Document) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = helpers.CUID()

	return
}

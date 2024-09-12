package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title     string         `gorm:"not null" json:"title"`
	Body      string         `gorm:"not null" json:"body"`
}

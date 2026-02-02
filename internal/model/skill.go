package model

import (
	"time"

	"gorm.io/gorm"
)

type Skill struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name       string `json:"name" gorm:"not null"`
	Category   string `json:"category"`
	Icon       string `json:"icon"`
	Percentage int    `json:"percentage"`
}

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

	Name       string `json:"name" gorm:"not null"` // e.g. "Golang"
	Category   string `json:"category"`             // e.g. "Backend", "Frontend", "Database"
	Icon       string `json:"icon"`                 // e.g. URL ya Icon Class string
	Percentage int    `json:"percentage"`           // e.g. 80 (Agar progress bar dikhana ho)
}

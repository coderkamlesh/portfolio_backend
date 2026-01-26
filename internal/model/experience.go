package model

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Company     string `json:"company" gorm:"not null"`  // e.g. "Google"
	Position    string `json:"position" gorm:"not null"` // e.g. "Senior Go Developer"
	Duration    string `json:"duration"`                 // e.g. "Jan 2023 - Present"
	Description string `json:"description"`              // Work details
	Location    string `json:"location"`                 // e.g. "Remote" or "Bangalore"
	Logo        string `json:"logo"`                     // Company Logo URL
	SortOrder   int    `json:"sort_order" gorm:"default:0"`
}

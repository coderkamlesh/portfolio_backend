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

	Company     string `json:"company" gorm:"not null"`
	Position    string `json:"position" gorm:"not null"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Logo        string `json:"logo"`
	SortOrder   int    `json:"sort_order" gorm:"default:0"`
}

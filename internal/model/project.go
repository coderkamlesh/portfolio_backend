package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	TechStack   string `json:"tech_stack"` // e.g. "Go, Gin, Turso"
	RepoLink    string `json:"repo_link"`
	LiveLink    string `json:"live_link"`
	ImageURL    string `json:"image_url"`
	SortOrder   int    `gorm:"default:0" json:"sort_order"` // Project upar/neeche dikhane ke liye
}

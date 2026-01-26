package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Auth Fields
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"` // JSON me password nahi bhejenge

	// Hero Section / Portfolio Fields
	FullName     string `json:"full_name"`
	JobTitle     string `json:"job_title"`   // e.g., "Golang Backend Developer"
	Description  string `json:"description"` // Short bio for Hero section
	ResumeLink   string `json:"resume_link"` // Link to PDF
	GithubLink   string `json:"github_link"`
	LinkedinLink string `json:"linkedin_link"`
	AvatarURL    string `json:"avatar_url"`
}

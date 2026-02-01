package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Auth Fields
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`

	// Hero Section / Portfolio Fields
	FullName     string         `json:"full_name"`
	JobTitle     string         `json:"job_title"`
	Description  string         `json:"description"`
	ResumeLink   string         `json:"resume_link"`
	GithubLink   string         `json:"github_link"`
	LinkedinLink string         `json:"linkedin_link"`
	AvatarURL    string         `json:"avatar_url"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

package dto

// Login ke liye request structure
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Admin Setup / Profile Update ke liye request structure
type SetupAdminRequest struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6"` // Validations bhi laga sakte hain
	FullName     string `json:"full_name" binding:"required"`
	JobTitle     string `json:"job_title"`
	Description  string `json:"description"`
	ResumeLink   string `json:"resume_link"`
	GithubLink   string `json:"github_link"`
	LinkedinLink string `json:"linkedin_link"`
	AvatarURL    string `json:"avatar_url"`
}
type UpdateProfileRequest struct {
	FullName     string `json:"full_name"`
	JobTitle     string `json:"job_title"`
	Description  string `json:"description"`
	ResumeLink   string `json:"resume_link"`
	GithubLink   string `json:"github_link"`
	LinkedinLink string `json:"linkedin_link"`
	AvatarURL    string `json:"avatar_url"`
}

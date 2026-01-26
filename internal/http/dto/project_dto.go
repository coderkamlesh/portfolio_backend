package dto

type CreateProjectRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	TechStack   string `json:"tech_stack"`
	RepoLink    string `json:"repo_link"`
	LiveLink    string `json:"live_link"`
	ImageURL    string `json:"image_url"`
	SortOrder   int    `json:"sort_order"`
}

type UpdateProjectRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	TechStack   string `json:"tech_stack"`
	RepoLink    string `json:"repo_link"`
	LiveLink    string `json:"live_link"`
	ImageURL    string `json:"image_url"`
	SortOrder   int    `json:"sort_order"`
}

package dto

type CreateExperienceRequest struct {
	Company     string `json:"company" binding:"required"`
	Position    string `json:"position" binding:"required"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Logo        string `json:"logo"`
	SortOrder   int    `json:"sort_order"`
}

type UpdateExperienceRequest struct {
	Company     string `json:"company"`
	Position    string `json:"position"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Logo        string `json:"logo"`
	SortOrder   int    `json:"sort_order"`
}

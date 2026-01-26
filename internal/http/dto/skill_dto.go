package dto

type CreateSkillRequest struct {
	Name       string `json:"name" binding:"required"`
	Category   string `json:"category" binding:"required"`
	Icon       string `json:"icon"`
	Percentage int    `json:"percentage"`
}

type UpdateSkillRequest struct {
	Name       string `json:"name"`
	Category   string `json:"category"`
	Icon       string `json:"icon"`
	Percentage int    `json:"percentage"`
}

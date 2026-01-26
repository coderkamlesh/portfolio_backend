package repository

import (
	"github.com/coderkamlesh/portfolio_backend/config"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
)

type ProjectRepository struct{}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{}
}

// Create
func (r *ProjectRepository) Create(project *model.Project) error {
	return config.DB.Create(project).Error
}

// Read All (Public - Sorted by SortOrder)
func (r *ProjectRepository) GetAll() ([]model.Project, error) {
	var projects []model.Project
	// SortOrder ke hisaab se ascending order me layenge
	result := config.DB.Order("sort_order asc").Find(&projects)
	return projects, result.Error
}

// Find By ID (Internal helper)
func (r *ProjectRepository) FindByID(id uint) (*model.Project, error) {
	var project model.Project
	result := config.DB.First(&project, id)
	return &project, result.Error
}

// Update
func (r *ProjectRepository) Update(project *model.Project) error {
	return config.DB.Save(project).Error
}

// Delete
func (r *ProjectRepository) Delete(id uint) error {
	return config.DB.Delete(&model.Project{}, id).Error
}

package repository

import (
	"github.com/coderkamlesh/portfolio_backend/config"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
)

type ExperienceRepository struct{}

func NewExperienceRepository() *ExperienceRepository {
	return &ExperienceRepository{}
}

func (r *ExperienceRepository) Create(exp *model.Experience) error {
	return config.DB.Create(exp).Error
}

func (r *ExperienceRepository) GetAll() ([]model.Experience, error) {
	var experiences []model.Experience
	// SortOrder ke hisaab se sort karenge (Ascending: 1 pehle, 2 baad mein)
	result := config.DB.Order("sort_order asc").Find(&experiences)
	return experiences, result.Error
}

func (r *ExperienceRepository) FindByID(id uint) (*model.Experience, error) {
	var exp model.Experience
	result := config.DB.First(&exp, id)
	return &exp, result.Error
}

func (r *ExperienceRepository) Update(exp *model.Experience) error {
	return config.DB.Save(exp).Error
}

func (r *ExperienceRepository) Delete(id uint) error {
	return config.DB.Delete(&model.Experience{}, id).Error
}

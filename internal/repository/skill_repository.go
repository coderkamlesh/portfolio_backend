package repository

import (
	"github.com/coderkamlesh/portfolio_backend/config"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
)

type SkillRepository struct{}

func NewSkillRepository() *SkillRepository {
	return &SkillRepository{}
}

func (r *SkillRepository) Create(skill *model.Skill) error {
	return config.DB.Create(skill).Error
}

func (r *SkillRepository) GetAll() ([]model.Skill, error) {
	var skills []model.Skill
	result := config.DB.Order("category asc").Find(&skills)
	return skills, result.Error
}

func (r *SkillRepository) FindByID(id uint) (*model.Skill, error) {
	var skill model.Skill
	result := config.DB.First(&skill, id)
	return &skill, result.Error
}

func (r *SkillRepository) Update(skill *model.Skill) error {
	return config.DB.Save(skill).Error
}

func (r *SkillRepository) Delete(id uint) error {
	return config.DB.Delete(&model.Skill{}, id).Error
}

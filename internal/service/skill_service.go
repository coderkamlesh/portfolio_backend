package service

import (
	"errors"

	"github.com/coderkamlesh/portfolio_backend/internal/http/dto"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
	"github.com/coderkamlesh/portfolio_backend/internal/repository"
)

type SkillService struct {
	Repo *repository.SkillRepository
}

func NewSkillService(repo *repository.SkillRepository) *SkillService {
	return &SkillService{Repo: repo}
}

func (s *SkillService) GetAllSkills() ([]model.Skill, error) {
	return s.Repo.GetAll()
}

func (s *SkillService) CreateSkill(req dto.CreateSkillRequest) (*model.Skill, error) {
	skill := model.Skill{
		Name:       req.Name,
		Category:   req.Category,
		Icon:       req.Icon,
		Percentage: req.Percentage,
	}
	if err := s.Repo.Create(&skill); err != nil {
		return nil, err
	}
	return &skill, nil
}

func (s *SkillService) UpdateSkill(id uint, req dto.UpdateSkillRequest) (*model.Skill, error) {
	skill, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("skill not found")
	}

	skill.Name = req.Name
	skill.Category = req.Category
	skill.Icon = req.Icon
	skill.Percentage = req.Percentage

	if err := s.Repo.Update(skill); err != nil {
		return nil, err
	}
	return skill, nil
}

func (s *SkillService) DeleteSkill(id uint) error {
	if _, err := s.Repo.FindByID(id); err != nil {
		return errors.New("skill not found")
	}
	return s.Repo.Delete(id)
}

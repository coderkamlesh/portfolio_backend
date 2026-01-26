package service

import (
	"errors"

	"github.com/coderkamlesh/portfolio_backend/internal/http/dto"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
	"github.com/coderkamlesh/portfolio_backend/internal/repository"
)

type ExperienceService struct {
	Repo *repository.ExperienceRepository
}

func NewExperienceService(repo *repository.ExperienceRepository) *ExperienceService {
	return &ExperienceService{Repo: repo}
}

func (s *ExperienceService) GetAllExperiences() ([]model.Experience, error) {
	return s.Repo.GetAll()
}

func (s *ExperienceService) CreateExperience(req dto.CreateExperienceRequest) (*model.Experience, error) {
	exp := model.Experience{
		Company:     req.Company,
		Position:    req.Position,
		Duration:    req.Duration,
		Description: req.Description,
		Location:    req.Location,
		Logo:        req.Logo,
		SortOrder:   req.SortOrder,
	}
	if err := s.Repo.Create(&exp); err != nil {
		return nil, err
	}
	return &exp, nil
}

func (s *ExperienceService) UpdateExperience(id uint, req dto.UpdateExperienceRequest) (*model.Experience, error) {
	exp, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("experience not found")
	}

	exp.Company = req.Company
	exp.Position = req.Position
	exp.Duration = req.Duration
	exp.Description = req.Description
	exp.Location = req.Location
	exp.Logo = req.Logo
	exp.SortOrder = req.SortOrder

	if err := s.Repo.Update(exp); err != nil {
		return nil, err
	}
	return exp, nil
}

func (s *ExperienceService) DeleteExperience(id uint) error {
	if _, err := s.Repo.FindByID(id); err != nil {
		return errors.New("experience not found")
	}
	return s.Repo.Delete(id)
}

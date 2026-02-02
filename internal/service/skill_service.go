package service

import (
	"errors"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/coderkamlesh/portfolio_backend/internal/http/dto"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
	"github.com/coderkamlesh/portfolio_backend/internal/repository"
	"github.com/coderkamlesh/portfolio_backend/internal/utils"
)

type SkillService struct {
	Repo *repository.SkillRepository // <--- Yahan * add kiya
	Cld  *cloudinary.Cloudinary
}

// Update Constructor to accept Pointer
func NewSkillService(repo *repository.SkillRepository, cld *cloudinary.Cloudinary) *SkillService { // <--- Yahan bhi * add kiya
	return &SkillService{Repo: repo, Cld: cld}
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

func (s *SkillService) UpdateSkillIcon(id uint, file *multipart.FileHeader) (string, error) {
	skill, err := s.Repo.FindByID(id)
	if err != nil {
		return "", err
	}

	// Delete Old Icon if exists
	if skill.Icon != "" {
		_ = utils.DeleteFromCloudinary(s.Cld, skill.Icon)
	}

	// Upload New Icon
	newURL, err := utils.UploadToCloudinary(s.Cld, file, "skills")
	if err != nil {
		return "", err
	}

	// Update DB
	skill.Icon = newURL
	err = s.Repo.Update(skill)

	return newURL, err
}

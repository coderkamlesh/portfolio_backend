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

type ProjectService struct {
	Repo *repository.ProjectRepository
	Cld  *cloudinary.Cloudinary
}

func NewProjectService(repo *repository.ProjectRepository, cld *cloudinary.Cloudinary) *ProjectService {
	return &ProjectService{Repo: repo, Cld: cld}
}

func (s *ProjectService) UpdateProjectImage(id uint, file *multipart.FileHeader) (string, error) {
	// 1. Project find karo
	project, err := s.Repo.FindByID(id)
	if err != nil {
		return "", err
	}

	// 2. Old Image Delete
	if project.ImageURL != "" {
		_ = utils.DeleteFromCloudinary(s.Cld, project.ImageURL)
	}

	// 3. New Image Upload
	newURL, err := utils.UploadToCloudinary(s.Cld, file, "projects")
	if err != nil {
		return "", err
	}

	// 4. DB Update
	project.ImageURL = newURL
	err = s.Repo.Update(project)

	return newURL, err
}

// 1. Get All Projects (Public)
func (s *ProjectService) GetAllProjects() ([]model.Project, error) {
	return s.Repo.GetAll()
}

// 2. Create Project (Admin)
func (s *ProjectService) CreateProject(req dto.CreateProjectRequest) (*model.Project, error) {
	project := model.Project{
		Title:       req.Title,
		Description: req.Description,
		TechStack:   req.TechStack,
		RepoLink:    req.RepoLink,
		LiveLink:    req.LiveLink,
		ImageURL:    req.ImageURL,
		SortOrder:   req.SortOrder,
	}

	if err := s.Repo.Create(&project); err != nil {
		return nil, err
	}
	return &project, nil
}

// 3. Update Project (Admin)
func (s *ProjectService) UpdateProject(id uint, req dto.UpdateProjectRequest) (*model.Project, error) {
	// Pehle check karo project exist karta hai ya nahi
	project, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("project not found")
	}

	// Update fields
	project.Title = req.Title
	project.Description = req.Description
	project.TechStack = req.TechStack
	project.RepoLink = req.RepoLink
	project.LiveLink = req.LiveLink
	project.ImageURL = req.ImageURL
	project.SortOrder = req.SortOrder

	if err := s.Repo.Update(project); err != nil {
		return nil, err
	}
	return project, nil
}

// 4. Delete Project (Admin)
func (s *ProjectService) DeleteProject(id uint) error {
	// Check existence
	_, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("project not found")
	}
	return s.Repo.Delete(id)
}

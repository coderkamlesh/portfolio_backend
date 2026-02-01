package handler

import (
	"net/http"

	"github.com/coderkamlesh/portfolio_backend/internal/http/dto" // Import DTO package
	"github.com/coderkamlesh/portfolio_backend/internal/model"
	"github.com/coderkamlesh/portfolio_backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// POST /api/login
func (h *AuthHandler) Login(c *gin.Context) {
	// DTO ka use
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// POST /api/setup
func (h *AuthHandler) SetupAdmin(c *gin.Context) {
	// DTO ka use
	var req dto.SetupAdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mapping DTO to Domain Model
	user := model.User{
		Email:        req.Email,
		Password:     req.Password, // DTO se password aayega, Model me store hoga (fir service hash karegi)
		FullName:     req.FullName,
		JobTitle:     req.JobTitle,
		Description:  req.Description,
		ResumeLink:   req.ResumeLink,
		GithubLink:   req.GithubLink,
		LinkedinLink: req.LinkedinLink,
		AvatarURL:    req.AvatarURL,
	}

	if err := h.Service.RegisterAdmin(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create admin: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin created successfully"})
}

// GET /api/hero (Ye same rahega kyunki isme input nahi hai)
func (h *AuthHandler) GetHeroInfo(c *gin.Context) {
	user, err := h.Service.GetPortfolioHero()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PUT /api/admin/hero
func (h *AuthHandler) UpdateHero(c *gin.Context) {
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Model mapping
	user := model.User{
		FullName:     req.FullName,
		JobTitle:     req.JobTitle,
		Description:  req.Description,
		ResumeLink:   req.ResumeLink,
		GithubLink:   req.GithubLink,
		LinkedinLink: req.LinkedinLink,
	}

	if err := h.Service.UpdateProfile(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// PUT /api/admin/hero/avatar
func (h *AuthHandler) UpdateAvatar(c *gin.Context) {
	// File receive karo
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	// Service call logic
	url, err := h.Service.UpdateAvatar(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar updated", "url": url})
}

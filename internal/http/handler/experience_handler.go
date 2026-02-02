package handler

import (
	"net/http"
	"strconv"

	"github.com/coderkamlesh/portfolio_backend/internal/http/dto"
	"github.com/coderkamlesh/portfolio_backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ExperienceHandler struct {
	Service *service.ExperienceService
}

func NewExperienceHandler(service *service.ExperienceService) *ExperienceHandler {
	return &ExperienceHandler{Service: service}
}

// GET /api/experiences
func (h *ExperienceHandler) GetAll(c *gin.Context) {
	experiences, err := h.Service.GetAllExperiences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch experiences"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": experiences})
}

// POST /api/admin/experiences
func (h *ExperienceHandler) Create(c *gin.Context) {
	var req dto.CreateExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exp, err := h.Service.CreateExperience(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add experience"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": exp})
}

// PUT /api/admin/experiences/:id
func (h *ExperienceHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req dto.UpdateExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exp, err := h.Service.UpdateExperience(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": exp})
}

// DELETE /api/admin/experiences/:id
func (h *ExperienceHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Service.DeleteExperience(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Experience deleted successfully"})
}

// Add this method to ExperienceHandler
func (h *ExperienceHandler) UpdateLogo(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Logo file is required"})
		return
	}

	url, err := h.Service.UpdateExperienceLogo(uint(id), file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience logo updated", "url": url})
}

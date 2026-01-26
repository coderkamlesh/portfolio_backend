package handler

import (
	"net/http"
	"strconv"

	"github.com/coderkamlesh/portfolio_backend/internal/http/dto"
	"github.com/coderkamlesh/portfolio_backend/internal/service"
	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	Service *service.SkillService
}

func NewSkillHandler(service *service.SkillService) *SkillHandler {
	return &SkillHandler{Service: service}
}

// GET /api/skills
func (h *SkillHandler) GetAll(c *gin.Context) {
	skills, err := h.Service.GetAllSkills()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch skills"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": skills})
}

// POST /api/admin/skills
func (h *SkillHandler) Create(c *gin.Context) {
	var req dto.CreateSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := h.Service.CreateSkill(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add skill"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": skill})
}

// PUT /api/admin/skills/:id
func (h *SkillHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req dto.UpdateSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := h.Service.UpdateSkill(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": skill})
}

// DELETE /api/admin/skills/:id
func (h *SkillHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Service.DeleteSkill(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill deleted successfully"})
}

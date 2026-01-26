package handler

import (
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/coderkamlesh/portfolio_backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	Cld *cloudinary.Cloudinary
}

func NewUploadHandler(cld *cloudinary.Cloudinary) *UploadHandler {
	return &UploadHandler{Cld: cld}
}

// POST /api/admin/upload
func (h *UploadHandler) UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	folder := c.DefaultQuery("folder", "portfolio")

	url, err := utils.UploadToCloudinary(h.Cld, file, folder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"url":     url,
	})
}

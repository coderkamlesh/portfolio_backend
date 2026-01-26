package utils

import (
	"context"
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// 2MB Limit
const MaxFileSize = 2 * 1024 * 1024

func UploadToCloudinary(cld *cloudinary.Cloudinary, file *multipart.FileHeader, folder string) (string, error) {
	// 1. Size Validation
	if file.Size > MaxFileSize {
		return "", errors.New("file size exceeds 2MB limit")
	}

	// 2. File Type Validation (Extension check)
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" && ext != ".svg" {
		return "", errors.New("invalid file type. allowed: jpg, jpeg, png, webp, svg")
	}

	// 3. Open File
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 4. Upload to Cloudinary
	ctx := context.Background()
	resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{
		Folder: "portfolio_go/" + folder, // Cloudinary me folder banega
	})

	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
func DeleteFromCloudinary(cld *cloudinary.Cloudinary, fileURL string) error {
	if fileURL == "" {
		return nil
	}

	// Helper function call karke Public ID nikalo
	publicID := GetPublicIDFromURL(fileURL)
	if publicID == "" {
		return nil
	}

	ctx := context.Background()
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	return err
}

// 2. URL se Public ID extract karne ka helper
// URL Example: https://res.cloudinary.com/.../upload/v123/portfolio_go/avatars/image.jpg
// Public ID: portfolio_go/avatars/image
func GetPublicIDFromURL(url string) string {
	// "portfolio_go" hamara folder prefix hai jo humne upload time diya tha
	if !strings.Contains(url, "portfolio_go") {
		return ""
	}

	// URL split karke last parts nikalo
	parts := strings.Split(url, "portfolio_go/")
	if len(parts) < 2 {
		return ""
	}

	// Extension (.jpg, .png) remove karo
	pathWithExt := "portfolio_go/" + parts[1]
	lastDotIndex := strings.LastIndex(pathWithExt, ".")
	if lastDotIndex == -1 {
		return pathWithExt
	}

	return pathWithExt[:lastDotIndex]
}

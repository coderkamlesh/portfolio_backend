package config

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

func SetupCloudinary() *cloudinary.Cloudinary {
	// Ensure config is loaded
	if Envs == nil {
		log.Fatal("Config not loaded. Call LoadConfig() first.")
	}

	cld, err := cloudinary.NewFromParams(
		Envs.CLOUDINARY_CLOUD_NAME,
		Envs.CLOUDINARY_API_KEY,
		Envs.CLOUDINARY_API_SECRET,
	)
	if err != nil {
		log.Fatal("Failed to initialize Cloudinary:", err)
	}
	return cld
}

package service

import (
	"errors"
	"mime/multipart"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
	"github.com/coderkamlesh/portfolio_backend/internal/repository"
	"github.com/coderkamlesh/portfolio_backend/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
	Cld  *cloudinary.Cloudinary // <--- Added
}

func NewAuthService(repo *repository.UserRepository, cld *cloudinary.Cloudinary) *AuthService {
	return &AuthService{Repo: repo, Cld: cld}
}

// Update Avatar Logic
func (s *AuthService) UpdateAvatar(file *multipart.FileHeader) (string, error) {
	// 1. Current User nikalo
	user, err := s.Repo.GetFirstUser()
	if err != nil {
		return "", err
	}

	// 2. Agar purana avatar hai, toh Cloudinary se delete karo
	if user.AvatarURL != "" {
		_ = utils.DeleteFromCloudinary(s.Cld, user.AvatarURL)
	}

	// 3. Naya Avatar Upload karo
	newURL, err := utils.UploadToCloudinary(s.Cld, file, "avatars")
	if err != nil {
		return "", err
	}

	// 4. DB Update karo
	user.AvatarURL = newURL
	err = s.Repo.UpdateUser(user)

	return newURL, err
}

// Login Logic
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Password compare karo
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// JWT Token Generate karo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Hero Section Data Fetch
func (s *AuthService) GetPortfolioHero() (*model.User, error) {
	return s.Repo.GetFirstUser()
}

// Seed Admin (Pehli baar user banane ke liye)
func (s *AuthService) RegisterAdmin(user *model.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	return s.Repo.CreateUser(user)
}

// Update Admin Profile (Hero Section)
func (s *AuthService) UpdateProfile(user *model.User) error {
	// Existing user find karo (Assume ID 1 is admin, or pass ID)
	existingUser, err := s.Repo.GetFirstUser()
	if err != nil {
		return err
	}

	// Fields update karo
	existingUser.FullName = user.FullName
	existingUser.JobTitle = user.JobTitle
	existingUser.Description = user.Description
	existingUser.ResumeLink = user.ResumeLink
	existingUser.GithubLink = user.GithubLink
	existingUser.LinkedinLink = user.LinkedinLink
	existingUser.AvatarURL = user.AvatarURL

	// Repo update logic (User Repo me Update method add karna padega)
	return s.Repo.UpdateUser(existingUser)
}

package repository

import (
	"github.com/coderkamlesh/portfolio_backend/config"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Admin login ke liye email find karna
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := config.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

// Portfolio display ke liye (assume karte hain ek hi admin/user hai)
func (r *UserRepository) GetFirstUser() (*model.User, error) {
	var user model.User
	result := config.DB.First(&user)
	return &user, result.Error
}

// Initial setup ke liye user create karna
func (r *UserRepository) CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}
func (r *UserRepository) UpdateUser(user *model.User) error {
	return config.DB.Save(user).Error
}

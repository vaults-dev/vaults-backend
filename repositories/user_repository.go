package repositories

import (
	"github.com/vaults-dev/vaults-backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	gorm *gorm.DB
}

type UserRepositoryInterface interface {
	CreateUser(user *models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}

func NewUserRepository(gorm *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		gorm: gorm,
	}
}

func (r *UserRepository) CreateUser(user *models.User) (models.User, error) {
	err := r.gorm.Create(user).Error
	if err != nil {
		// TODO LOGGER
	}

	return *user, err
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := r.gorm.First(&user, "email=?", email).Error
	if err != nil {
		// TODO LOGGER
	}

	return user, err
}

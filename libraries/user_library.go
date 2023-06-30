package libraries

import (
	"fmt"

	"github.com/vaults-dev/vaults-backend/models"
	"github.com/vaults-dev/vaults-backend/repositories"
	"github.com/vaults-dev/vaults-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserLibrary struct {
	repo repositories.UserRepositoryInterface
}

type UserLibraryInterface interface {
	SignUp(params models.SignUp) error
	Login(params models.Login) (interface{}, error)
}

func NewUserLibrary(repo repositories.UserRepositoryInterface) UserLibraryInterface {
	return &UserLibrary{
		repo: repo,
	}
}

func (l *UserLibrary) SignUp(request models.SignUp) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fmt.Errorf("failed generate hash, %v", err.Error())
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashPass),
	}

	err = l.repo.CreateUser(&user)
	if err != nil {
		return fmt.Errorf("failed create user to db, %v", err.Error())
	}

	return nil
}

func (l *UserLibrary) Login(request models.Login) (interface{}, error) {
	user, _ := l.repo.GetUserByEmail(request.Email)
	if user.Email == "" {
		return nil, fmt.Errorf("wrong email or pass")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong email or pass")
	}

	jwt, err := utils.GenerateTokenForUser(user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed generate jwt, %v", err.Error())
	}

	return string(jwt), nil
}

package services

import (
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/models"
	"github.com/kousuketk/crudGo/app/repository"
)

type UserServices struct{}

var userRepo repository.UserRepository

func (UserServices) GetUsers() ([]models.User, error) {
	result, err := userRepo.GetUsers()

	return result, err
}

func (UserServices) GetUserById(id int) (models.User, error) {
	result, err := userRepo.GetUserById(id)

	return result, err
}

func (UserServices) CreateUser(user models.User) (models.User, error) {
	hashedPassword, _ := middlewares.CreatePassword(user.PasswordDigest)
	user.PasswordDigest = hashedPassword
	result, err := userRepo.CreateUser(user)

	return result, err
}

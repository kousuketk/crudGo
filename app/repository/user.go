package repository

import (
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/models"
)

type UserRepository struct{}

func (UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := middlewares.DB().Find(&users)

	return users, err.Error
}

func (UserRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := middlewares.DB().Where("id = ?", id).First(&user)

	return user, err.Error
}

func (UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := middlewares.DB().Where("email = ?", email).First(&user)

	return user, err.Error
}

func (UserRepository) CreateUser(user models.User) (models.User, error) {
	err := middlewares.DB().Create(&user)

	return user, err.Error
}

func (UserRepository) UpdateUser(id int, user models.User) error {
	err := middlewares.DB().Where("id = ?", id).Updates(&user)

	return err.Error
}

func (UserRepository) DeleteUser(id int) error {
	err := middlewares.DB().Table("users").Delete(id)

	return err.Error
}

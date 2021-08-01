package services

import (
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/models"
)

type UserServices struct{}

func (UserServices) GetUsers() ([]*models.User, error) {
	var users []*models.User
	result := middlewares.DB().Find(&users)

	return users, result.Error
}

func (UserServices) GetUserById(ID int) (models.User, error) {
	var user models.User
	result := middlewares.DB().Where("id = ?", ID).First(&user)

	return user, result.Error
}

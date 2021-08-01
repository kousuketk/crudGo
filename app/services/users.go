package services

import (
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/models"
)

type UserServices struct{}

func (UserServices) GetUsers() []*models.User {
	var users []*models.User
	// users := make([]models.User, 0)
	err := middlewares.DB().Find(&users)
	if err != nil {
		panic(err)
	}

	return users
}

func (UserServices) GetUserById(ID int) models.User {
	var user models.User
	err := middlewares.DB().Where("id = ?", ID).First(&user)
	if err != nil {
		panic(err)
	}

	return user
}

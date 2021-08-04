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

func (UserServices) GetUserById(Id int) (models.User, error) {
	var user models.User
	result := middlewares.DB().Where("id = ?", Id).First(&user)

	return user, result.Error
}

func (UserServices) GetUserByEmail(Email string) (models.User, error) {
	var user models.User
	result := middlewares.DB().Where("email = ?", Email).First(&user)

	return user, result.Error
}

func (UserServices) CreateUser(name string, selfIntroduction string, email string, passwordDigest string, address string, phoneNumber string) (models.User, error) {
	user := models.User{
		Name:             name,
		SelfIntroduction: selfIntroduction,
		Email:            email,
		PasswordDigest:   passwordDigest,
		Address:          address,
		PhoneNumber:      phoneNumber}
	result := middlewares.DB().Create(&user)

	return user, result.Error
}

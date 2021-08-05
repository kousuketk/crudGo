package services

import (
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/models"
)

type UserServices struct{}

func (UserServices) GetUsers() ([]models.User, error) {
	var users []models.User
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

func (UserServices) CreateUser(user models.User) (models.User, error) {
	hashedPassword, _ := middlewares.CreatePassword(user.PasswordDigest)
	user.PasswordDigest = hashedPassword
	result := middlewares.DB().Create(&user)

	return user, result.Error
}

func (UserServices) UpdateUser(userId int, user models.User) error {
	hashedPassword, _ := middlewares.CreatePassword(user.PasswordDigest)
	user.PasswordDigest = hashedPassword
	result := middlewares.DB().Where("id = ?", userId).Updates(&user)

	return result.Error
}

func (UserServices) DeleteUser(userId int) error {
	result := middlewares.DB().Delete(&models.User{}, userId)

	return result.Error
}

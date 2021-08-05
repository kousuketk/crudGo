package services

import (
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/models"
	"github.com/kousuketk/crudGo/app/repository"
)

type MeServices struct{}

var meRepo repository.UserRepository

func (MeServices) Index(id int) (models.User, error) {
	result, err := meRepo.GetUserById(id)

	return result, err
}

// パスワード変更はできるけど検討しないといけない
func (MeServices) UpdateMe(id int, user models.User) error {
	hashedPassword, _ := middlewares.CreatePassword(user.PasswordDigest)
	user.PasswordDigest = hashedPassword
	err := meRepo.UpdateUser(id, user)

	return err
}

func (MeServices) DeleteMe(id int) error {
	err := meRepo.DeleteUser(id)

	return err
}

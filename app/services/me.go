package services

import (
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/models"
)

type MeServices struct{}

func (MeServices) Index(id int) (models.User, error) {
	result, err := userRepo.GetUserById(id)

	return result, err
}

// パスワード変更はできるけど検討しないといけない
func (MeServices) UpdateMe(id int, user models.User) error {
	hashedPassword, _ := middlewares.CreatePassword(user.PasswordDigest)
	user.PasswordDigest = hashedPassword
	err := userRepo.UpdateUser(id, user)

	return err
}

func (MeServices) DeleteMe(id int) error {
	err := userRepo.DeleteUser(id)

	return err
}

package middlewares

import (
	"github.com/kousuketk/crudGo/app/models"
)

func CreatePassword(rowPassword string) (string, error) {
	return "test", nil
}

func VerifyPassword(user models.User, password string) (bool, error) {
	if user.PasswordDigest == password {
		return true, nil
	}
	return false, nil
}

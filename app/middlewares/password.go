package middlewares

import (
	"github.com/kousuketk/crudGo/app/models"
	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(rowPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rowPassword), 10)
	return string(hashedPassword), err
}

func VerifyPassword(user models.User, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	if err == nil {
		return true, nil
	}
	return false, err
}

package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsers(db *gorm.DB) ([]*User, error) {
	users := []*User{}
	result := db.Find(&users)

	return users, result.Error
}

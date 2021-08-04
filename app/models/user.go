package models

import (
	"reflect"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `json:"name"`
	SelfIntroduction string `json:"selfIntroduction"`
	Email            string `json:"email"`
	PasswordDigest   string `json:"passwordDigest"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phoneNumber"`
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}

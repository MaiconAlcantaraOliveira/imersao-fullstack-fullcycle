package model

import (
	"github.com/asaskevich/govalidator"
)

type User struct {
	ID    string `json:"id" valid:"uuid" valid:"required"`
	Name  string `json:"name" valid:"required"`
	Email string `json:"email" valid:"required"`
}

func (user *User) NewUser() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

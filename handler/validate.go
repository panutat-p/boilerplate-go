package handler

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`
	Email     string `validate:"required,email"`
	Gender    string `validate:"oneof=male female"`
}

func ValidateUser(user User) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(user)
	if err != nil {
		return err
	}

	return nil
}
package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/guregu/null/v5"

	"boilerplate-go/pkg"
)

var validate *validator.Validate

type Fruit struct {
	Name    string      `validate:"required"`
	Price   int         `validate:"required"`
	Factory null.String `validate:"omitempty,max=5"`
}

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate = pkg.RegisterNullTypes(validate)

	f := Fruit{
		Name:    "Apple",
		Price:   100,
		Factory: null.StringFrom("Robert"),
	}

	err := validate.Struct(f)
	if err != nil {
		panic(err)
	}
}

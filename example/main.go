package main

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/guregu/null/v5"
)

var validate *validator.Validate

type Fruit struct {
	Name    string      `validate:"required"`
	Price   int         `validate:"required"`
	Factory null.String `validate:"max=20"`
}

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate = RegisterNullTypes(validate)

	f := Fruit{
		Name:    "Apple",
		Price:   100,
		Factory: null.StringFrom("aaaaabbbbbcccccddddde"),
	}

	err := validate.Struct(f)
	if err != nil {
		panic(err)
	}
}

func RegisterNullTypes(validate *validator.Validate) *validator.Validate {
    validate.RegisterCustomTypeFunc(ValidateNull, null.String{})
    validate.RegisterCustomTypeFunc(ValidateNull, null.Float{})
    validate.RegisterCustomTypeFunc(ValidateNull, null.Bool{})
    validate.RegisterCustomTypeFunc(ValidateNull, null.Time{})
    validate.RegisterCustomTypeFunc(ValidateNull, null.Int{})
    validate.RegisterCustomTypeFunc(ValidateNull, null.Int16{})
    validate.RegisterCustomTypeFunc(ValidateNull, null.Int32{})
    validate.RegisterCustomTypeFunc(ValidateNull, null.Int64{})
    return validate
}

func ValidateNull(field reflect.Value) interface{} {
	// Get the concrete value
	v := field.Interface()

	switch t := v.(type) {
	case null.String:
		if t.Valid {
			return t.String
		}
	case null.Float:
		if t.Valid {
			return t.Float64
		}
	case null.Bool:
		if t.Valid {
			return t.Bool
		}
	case null.Time:
		if t.Valid {
			return t.Time
		}

	case null.Int16:
		if t.Valid {
			return t.Int16
		}
	case null.Int32:
		if t.Valid {
			return t.Int32
		}
	case null.Int:
		if t.Valid {
			return t.Int64
		}
	}
	return nil
}

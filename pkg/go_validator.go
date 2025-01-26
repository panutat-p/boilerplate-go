package pkg

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/guregu/null/v5"
	"github.com/shopspring/decimal"
)

func RegisterNullTypes(validate *validator.Validate) *validator.Validate {
	validate.RegisterCustomTypeFunc(
		ValidateNull,
		null.String{},
		null.Bool{},
		null.Float{},
		null.Int16{},
		null.Int32{},
		null.Int{},
		null.Time{},
	)
	return validate
}

func RegisterDecimalTypes(validate *validator.Validate) *validator.Validate {
	validate.RegisterCustomTypeFunc(
		ValidateDecimal,
		decimal.Decimal{},
		decimal.NullDecimal{},
	)
	return validate
}

func ValidateNull(field reflect.Value) interface{} {
	v := field.Interface()

	switch t := v.(type) {
	case null.String:
		if t.Valid {
			return t.String
		}
	case null.Bool:
		if t.Valid {
			return t.Bool
		}
	case null.Float:
		if t.Valid {
			return t.Float64
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
	case null.Time:
		if t.Valid {
			return t.Time
		}
	}
	return nil
}

func ValidateDecimal(field reflect.Value) interface{} {
	v := field.Interface()
	switch t := v.(type) {
	case decimal.Decimal:
		return t.String()
	case decimal.NullDecimal:
		if !t.Valid {
			return nil
		}
		return t.Decimal.String()
	}
	return nil
}

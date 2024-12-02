package model

import (
	"github.com/shopspring/decimal"
)

type Fruit struct {
	Name  string          `json:"name"`
	Color string          `json:"color"`
	Price decimal.Decimal `json:"price"`
}

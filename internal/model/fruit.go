package model

import (
	"github.com/guregu/null/v5"
	"github.com/shopspring/decimal"
)

type Fruit struct {
	Name    string          `json:"name"`
	Color   string          `json:"color"`
	Price   decimal.Decimal `json:"price"`
	Factory null.String     `json:"factory"`
}

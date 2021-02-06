package models

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	Code         string          `json:"code"`
	Name         string          `json:"name"`
	PricePerUnit decimal.Decimal `json:"price"`
}

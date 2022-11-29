package models

import (
	"fmt"
	"net/http"
)

type Product struct {
	ProductID   int64    `json:"product_id"`
	ProductName string   `json:"product_name"`
	UnitPrice   *float64 `json:"unit_price"`
	InStock     int      `json:"in_stock"`
}

type ProductList struct {
	Products []Product `json:"products"`
}

func (p *Product) Bind(r *http.Request) error {
	if p.ProductName == "" {
		return fmt.Errorf("product_name is a required field")
	}

	if p.UnitPrice == nil {
		return fmt.Errorf("unit_price is a required field")
	}

	return nil
}

func (*ProductList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Product) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

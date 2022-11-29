package models

import (
	"fmt"
	"net/http"
)

type OrderDetail struct {
	OrderDetailID int64    `json:"order_detail_id"`
	OrderID       *int64   `json:"order_id"`
	ProductID     *int64   `json:"product_id"`
	Quantity      *int     `json:"quantity"`
	UnitPrice     *float64 `json:"unit_price"`
	Discount      *float64 `json:"discount"`
}

type OrderDetailList struct {
	OrderDetails []OrderDetail `json:"order_details"`
}

func (o *OrderDetail) Bind(r *http.Request) error {
	if o.OrderID == nil {
		return fmt.Errorf("order_id is a required field")
	}

	if o.ProductID == nil {
		return fmt.Errorf("product_id is a required field")
	}

	if o.Quantity == nil {
		return fmt.Errorf("quantity is a required field")
	}

	if o.UnitPrice == nil {
		return fmt.Errorf("unit_price is a required field")
	}

	if o.Discount == nil {
		return fmt.Errorf("discount is a required field")
	}

	return nil
}

func (*OrderDetailList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*OrderDetail) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

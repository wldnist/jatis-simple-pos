package models

import (
	"fmt"
	"net/http"
	"time"
)

type Order struct {
	OrderID             int64     `json:"order_id"`
	CustomerID          *int64    `json:"customer_id"`
	EmployeeID          *int64    `json:"employee_id"`
	OrderDate           time.Time `json:"order_date"`
	PurchaseOrderNumber string    `json:"purchase_order_number"`
	ShipDate            time.Time `json:"ship_date"`
	ShippingMethodID    *int64    `json:"shipping_method_id"`
	FreightCharge       *float64  `json:"freight_charge"`
	Taxes               *float64  `json:"taxes"`
	PaymentReceived     int       `json:"payment_received"`
	Comment             string    `json:"comment"`
}

type OrderList struct {
	Orders []Order `json:"orders"`
}

func (o *Order) Bind(r *http.Request) error {
	if o.CustomerID == nil {
		return fmt.Errorf("customer_id is a required field")
	}

	if o.EmployeeID == nil {
		return fmt.Errorf("employee_id is a required field")
	}

	if o.PurchaseOrderNumber == "" {
		return fmt.Errorf("purchase_order_number is a required field")
	}

	if o.ShippingMethodID == nil {
		return fmt.Errorf("shipping_method_id is a required field")
	}

	if o.FreightCharge == nil {
		return fmt.Errorf("freight_charge is a required field")
	}

	if o.Taxes == nil {
		return fmt.Errorf("taxes is a required field")
	}

	return nil
}

func (*OrderList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Order) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type BulkOrderReq struct {
	Url string `json:"url"`
}

type BulkOrderRes struct {
	IDs []int64 `json:"ids"`
}

func (b *BulkOrderReq) Bind(r *http.Request) error {
	if b.Url == "" {
		return fmt.Errorf("url is a required field")
	}

	return nil
}

func (*BulkOrderReq) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*BulkOrderRes) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

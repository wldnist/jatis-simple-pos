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

type BulkOrderDetailReq struct {
	Url string `json:"url"`
}

type BulkOrderDetailRes struct {
	IDs []int64 `json:"ids"`
}

func (b *BulkOrderDetailReq) Bind(r *http.Request) error {
	if b.Url == "" {
		return fmt.Errorf("url is a required field")
	}

	return nil
}

func (*BulkOrderDetailReq) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*BulkOrderDetailRes) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type CompleteOrderDetailsSpec struct {
	OrderID        int64   `json:"order_id"`
	ProductID      int64   `json:"product_id"`
	ProductName    string  `json:"product_name"`
	Quantity       int     `json:"quantity"`
	UnitPrice      float64 `json:"unit_price"`
	Discount       float64 `json:"discount"`
	CustomerName   string  `json:"customer_name"`
	EmployeeName   string  `json:"employee_name"`
	ShippingMethod string  `json:"shipping_method"`
	SubTotal       float64 `json:"sub_total"`
}

type CompleteOrderDetails struct {
	CustomerName   string             `json:"customer_name"`
	EmployeeName   string             `json:"employee_name"`
	ShippingMethod string             `json:"shipping_method"`
	TotalPayment   float64            `json:"total_payment"`
	OrderDetails   []OrderDetailsSpec `json:"order_details"`
}

type OrderDetailsSpec struct {
	OrderID     int64   `json:"order_id"`
	ProductID   int64   `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Discount    float64 `json:"discount"`
	SubTotal    float64 `json:"sub_total"`
}

type GetCompleteAllOrderDetailsRes struct {
	Result []CompleteOrderDetails `json:"result"`
}

func (*GetCompleteAllOrderDetailsRes) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

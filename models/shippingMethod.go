package models

import (
	"fmt"
	"net/http"
)

type ShippingMethod struct {
	ShippingMethodID int64  `json:"shipping_method_id"`
	ShippingMethod   string `json:"shipping_method"`
}

type ShippingMethodList struct {
	ShippingMethods []ShippingMethod `json:"shipping_methods"`
}

func (s *ShippingMethod) Bind(r *http.Request) error {
	if s.ShippingMethod == "" {
		return fmt.Errorf("shipping_method is a required field")
	}

	return nil
}

func (*ShippingMethodList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*ShippingMethod) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type BulkShippingMethodReq struct {
	Url string `json:"url"`
}

type BulkShippingMethodRes struct {
	IDs []int64 `json:"ids"`
}

func (b *BulkShippingMethodReq) Bind(r *http.Request) error {
	if b.Url == "" {
		return fmt.Errorf("url is a required field")
	}

	return nil
}

func (*BulkShippingMethodReq) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*BulkShippingMethodRes) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

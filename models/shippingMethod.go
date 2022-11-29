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

package models

import (
	"fmt"
	"net/http"
)

type Customer struct {
	CustomerID          int64  `json:"customer_id"`
	CompanyName         string `json:"company_name"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	BillingAddress      string `json:"billing_address"`
	City                string `json:"city"`
	StateOrProvince     string `json:"state_or_province"`
	ZIPCode             string `json:"zip_code"`
	Email               string `json:"email"`
	CompanyWebsite      string `json:"company_website"`
	PhoneNumber         string `json:"phone_number"`
	FaxNumber           string `json:"fax_number"`
	ShipAddress         string `json:"ship_address"`
	ShipCity            string `json:"ship_city"`
	ShipStateOrProvince string `json:"ship_state_or_province"`
	ShipZIPCode         string `json:"ship_zip_code"`
	ShipPhoneNumber     string `json:"ship_phone_number"`
}

type CustomerList struct {
	Customers []Customer `json:"customers"`
}

func (c *Customer) Bind(r *http.Request) error {
	if c.CompanyName == "" {
		return fmt.Errorf("company_name is a required field")
	}

	if c.FirstName == "" {
		return fmt.Errorf("first_name is a required field")
	}

	if c.BillingAddress == "" {
		return fmt.Errorf("billing_address is a required field")
	}

	if c.City == "" {
		return fmt.Errorf("city is a required field")
	}

	if c.StateOrProvince == "" {
		return fmt.Errorf("state_or_province is a required field")
	}

	if c.ZIPCode == "" {
		return fmt.Errorf("zip_code is a required field")
	}

	if c.Email == "" {
		return fmt.Errorf("email is a required field")
	}

	if c.PhoneNumber == "" {
		return fmt.Errorf("phone_number is a required field")
	}

	if c.ShipAddress == "" {
		return fmt.Errorf("ship_address is a required field")
	}

	if c.ShipCity == "" {
		return fmt.Errorf("ship_city is a required field")
	}

	if c.ShipStateOrProvince == "" {
		return fmt.Errorf("ship_state_or_province is a required field")
	}

	if c.ShipZIPCode == "" {
		return fmt.Errorf("ship_zip_code is a required field")
	}

	if c.ShipPhoneNumber == "" {
		return fmt.Errorf("ship_phone_number is a required field")
	}

	return nil
}

func (*CustomerList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Customer) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type BulkCustomerReq struct {
	Url string `json:"url"`
}

type BulkCustomerRes struct {
	IDs []int64 `json:"ids"`
}

func (b *BulkCustomerReq) Bind(r *http.Request) error {
	if b.Url == "" {
		return fmt.Errorf("url is a required field")
	}

	return nil
}

func (*BulkCustomerReq) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*BulkCustomerRes) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

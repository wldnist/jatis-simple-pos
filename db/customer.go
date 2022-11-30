package db

import (
	"log"

	"github.com/wldnist/jatis-simple-pos/models"
)

func (db Database) AddCustomer(customer *models.Customer) error {
	var customerID int64
	sql := "INSERT INTO customers (company_name, first_name, last_name, billing_address, city, state_or_province, zip_code, email, phone_number, ship_address, ship_city, ship_state_or_province, ship_zip_code, ship_phone_number) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	res, err := db.Conn.Exec(sql,
		customer.CompanyName,
		customer.FirstName,
		customer.LastName,
		customer.BillingAddress,
		customer.City,
		customer.StateOrProvince,
		customer.ZIPCode,
		customer.Email,
		customer.PhoneNumber,
		customer.ShipAddress,
		customer.ShipCity,
		customer.ShipStateOrProvince,
		customer.ShipZIPCode,
		customer.ShipPhoneNumber,
	)
	if err != nil {
		return err
	}

	customerID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	customer.CustomerID = customerID
	return nil
}

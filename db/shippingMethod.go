package db

import (
	"log"

	"github.com/wldnist/jatis-simple-pos/models"
)

func (db Database) AddShippingMethod(shippingMethod *models.ShippingMethod) error {
	var shippingMethodID int64
	sql := "INSERT INTO shipping_methods (shipping_method) VALUES (?)"
	res, err := db.Conn.Exec(sql, shippingMethod.ShippingMethod)
	if err != nil {
		return err
	}

	shippingMethodID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	shippingMethod.ShippingMethodID = shippingMethodID
	return nil
}

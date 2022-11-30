package db

import (
	"fmt"
	"log"

	"github.com/wldnist/jatis-simple-pos/models"
)

func (db Database) AddShippingMethod(shippingMethod *models.ShippingMethod) error {
	var shippingMethodID int64
	// query := `INSERT INTO shipping_methods (shipping_method) VALUES (?)`
	// err := db.Conn.QueryRow(query, shippingMethod.ShippingMethod).Scan()
	// insert, err := db.Conn.Query("INSERT INTO shipping_methods (shipping_method) VALUES (?)", shippingMethod.ShippingMethod)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(insert)

	sql := "INSERT INTO shipping_methods (shipping_method) VALUES (?)"
	res, err := db.Conn.Exec(sql, shippingMethod.ShippingMethod)
	if err != nil {
		return err
	}

	shippingMethodID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(shippingMethodID)

	shippingMethod.ShippingMethodID = shippingMethodID
	return nil
}

package db

import (
	"log"

	"github.com/wldnist/jatis-simple-pos/models"
)

func (db Database) AddProduct(product *models.Product) error {
	var productID int64
	sql := "INSERT INTO products (product_name, unit_price, in_stock) VALUES (?,?,?)"
	res, err := db.Conn.Exec(sql, product.ProductName, product.UnitPrice, product.InStock)
	if err != nil {
		return err
	}

	productID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	product.ProductID = productID
	return nil
}

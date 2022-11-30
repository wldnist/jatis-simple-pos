package db

import (
	"log"

	"github.com/wldnist/jatis-simple-pos/models"
)

func (db Database) AddOrderDetail(orderDetail *models.OrderDetail) error {
	var orderDetailID int64
	sql := "INSERT INTO order_details (order_id, product_id, quantity, unit_price, discount) VALUES (?,?,?,?,?)"
	res, err := db.Conn.Exec(sql,
		orderDetail.OrderID,
		orderDetail.ProductID,
		orderDetail.Quantity,
		orderDetail.UnitPrice,
		orderDetail.Discount,
	)
	if err != nil {
		return err
	}

	orderDetailID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	orderDetail.OrderDetailID = orderDetailID
	return nil
}

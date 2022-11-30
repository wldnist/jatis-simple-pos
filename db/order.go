package db

import (
	"log"

	"github.com/wldnist/jatis-simple-pos/models"
)

func (db Database) GetAllOrders() (*models.OrderList, error) {
	list := &models.OrderList{}
	rows, err := db.Conn.Query("SELECT * FROM orders ORDER BY order_id ASC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var order models.Order
		err := rows.Scan(
			&order.OrderID,
			&order.CustomerID,
			&order.EmployeeID,
			&order.OrderDate,
			&order.PurchaseOrderNumber,
			&order.ShipDate,
			&order.ShippingMethodID,
			&order.FreightCharge,
			&order.Taxes,
			&order.PaymentReceived,
			&order.Comment)
		if err != nil {
			return list, err
		}

		list.Orders = append(list.Orders, order)
	}

	return list, nil
}

func (db Database) AddOrder(order *models.Order) error {
	var orderID int64
	sql := "INSERT INTO orders (customer_id, employee_id, purchase_order_number, shipping_method_id, freight_charge, taxes) VALUES (?,?,?,?,?,?)"
	res, err := db.Conn.Exec(sql,
		order.CustomerID,
		order.EmployeeID,
		order.PurchaseOrderNumber,
		order.ShippingMethodID,
		order.FreightCharge,
		order.Taxes,
	)
	if err != nil {
		return err
	}

	orderID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	order.OrderID = orderID
	return nil
}

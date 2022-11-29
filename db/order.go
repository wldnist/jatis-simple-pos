package db

import (
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

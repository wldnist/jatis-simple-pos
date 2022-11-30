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

func (db Database) GetCompleteAllOrderDetails() ([]*models.CompleteOrderDetailsSpec, error) {
	list := []*models.CompleteOrderDetailsSpec{}
	rows, err := db.Conn.Query(`
		SELECT od.order_id, 
			od.product_id, 
			p.product_name, 
			SUM(od.quantity) as quantity, 
			SUM(od.unit_price) AS unit_price, 
			SUM(od.discount) AS discount, 
			CONCAT(c.first_name, " ", c.last_name) AS customer_name, 
			CONCAT(e.first_name, " ", e.last_name) AS employee_name, 
			sm.shipping_method, 
			SUM(((od.quantity * od.unit_price) - od.discount)) AS sub_total
		FROM order_details od
		LEFT JOIN products p ON od.product_id = p.product_id
		LEFT JOIN orders o ON od.order_id = o.order_id
		LEFT JOIN customers c ON o.customer_id = c.customer_id
		LEFT JOIN employees e ON o.employee_id = e.employee_id
		LEFT JOIN shipping_methods sm ON o.shipping_method_id = sm.shipping_method_id
		GROUP BY od.order_id, od.product_id`)
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var completeOrderDetailsSpec models.CompleteOrderDetailsSpec
		err := rows.Scan(
			&completeOrderDetailsSpec.OrderID,
			&completeOrderDetailsSpec.ProductID,
			&completeOrderDetailsSpec.ProductName,
			&completeOrderDetailsSpec.Quantity,
			&completeOrderDetailsSpec.UnitPrice,
			&completeOrderDetailsSpec.Discount,
			&completeOrderDetailsSpec.CustomerName,
			&completeOrderDetailsSpec.EmployeeName,
			&completeOrderDetailsSpec.ShippingMethod,
			&completeOrderDetailsSpec.SubTotal)
		if err != nil {
			return list, err
		}

		list = append(list, &completeOrderDetailsSpec)
	}

	return list, nil
}

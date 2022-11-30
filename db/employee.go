package db

import (
	"log"

	"github.com/wldnist/jatis-simple-pos/models"
)

func (db Database) AddEmployee(employee *models.Employee) error {
	var employeeID int64
	sql := "INSERT INTO employees (first_name, last_name, title, work_phone) VALUES (?,?,?,?)"
	res, err := db.Conn.Exec(sql, employee.FirstName, employee.LastName, employee.Title, employee.WorkPhone)
	if err != nil {
		return err
	}

	employeeID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	employee.EmployeeID = employeeID
	return nil
}

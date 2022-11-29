package models

import (
	"fmt"
	"net/http"
)

type Employee struct {
	EmployeeID int64  `json:"employee_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Title      string `json:"title"`
	WorkPhone  string `json:"work_phone"`
}

type EmployeeList struct {
	Employees []Employee `json:"employees"`
}

func (e *Employee) Bind(r *http.Request) error {
	if e.FirstName == "" {
		return fmt.Errorf("first_name is a required field")
	}

	if e.Title == "" {
		return fmt.Errorf("title is a required field")
	}

	if e.WorkPhone == "" {
		return fmt.Errorf("work_phone is a required field")
	}

	return nil
}

func (*EmployeeList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Employee) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

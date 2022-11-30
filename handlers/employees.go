package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/wldnist/jatis-simple-pos/helpers"
	"github.com/wldnist/jatis-simple-pos/models"
)

func employees(router chi.Router) {
	router.Post("/bulk", createBulkEmployees)
}

func createBulkEmployees(w http.ResponseWriter, r *http.Request) {
	bulkEmployeeReq := &models.BulkEmployeeReq{}
	if err := render.Bind(r, bulkEmployeeReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	data, err := helpers.ReadCSVFromUrl(bulkEmployeeReq.Url)
	if err != nil {
		panic(err)
	}

	result := []int64{}
	employee := &models.Employee{}
	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		employee.FirstName = row[0]
		employee.LastName = row[1]
		employee.Title = row[2]
		employee.WorkPhone = row[3]
		if err := dbInstance.AddEmployee(employee); err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}

		result = append(result, employee.EmployeeID)
	}

	bulkEmployeeRes := &models.BulkEmployeeRes{}
	bulkEmployeeRes.IDs = result
	if err := render.Render(w, r, bulkEmployeeRes); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

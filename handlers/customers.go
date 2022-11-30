package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/wldnist/jatis-simple-pos/helpers"
	"github.com/wldnist/jatis-simple-pos/models"
)

func customers(router chi.Router) {
	router.Post("/bulk", createBulkCustomers)
}

func createBulkCustomers(w http.ResponseWriter, r *http.Request) {
	bulkCustomerReq := &models.BulkCustomerReq{}
	if err := render.Bind(r, bulkCustomerReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	data, err := helpers.ReadCSVFromUrl(bulkCustomerReq.Url)
	if err != nil {
		panic(err)
	}

	result := []int64{}
	customer := &models.Customer{}
	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		customer.CompanyName = row[0]
		customer.FirstName = row[1]
		customer.LastName = row[2]
		customer.BillingAddress = row[3]
		customer.City = row[4]
		customer.StateOrProvince = row[5]
		customer.ZIPCode = row[6]
		customer.Email = row[7]
		customer.PhoneNumber = row[8]
		customer.ShipAddress = row[9]
		customer.ShipCity = row[10]
		customer.ShipStateOrProvince = row[11]
		customer.ShipZIPCode = row[12]
		customer.ShipPhoneNumber = row[13]
		if err := dbInstance.AddCustomer(customer); err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}

		result = append(result, customer.CustomerID)
	}

	bulkCustomerRes := &models.BulkCustomerRes{}
	bulkCustomerRes.IDs = result
	if err := render.Render(w, r, bulkCustomerRes); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

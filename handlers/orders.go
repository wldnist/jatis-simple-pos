package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/wldnist/jatis-simple-pos/helpers"
	"github.com/wldnist/jatis-simple-pos/models"
)

func orders(router chi.Router) {
	router.Get("/", getAllOrders)
	router.Post("/bulk", createBulkOrders)
}

func getAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := dbInstance.GetAllOrders()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, orders); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createBulkOrders(w http.ResponseWriter, r *http.Request) {
	bulkOrderReq := &models.BulkOrderReq{}
	if err := render.Bind(r, bulkOrderReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	data, err := helpers.ReadCSVFromUrl(bulkOrderReq.Url)
	if err != nil {
		panic(err)
	}

	result := []int64{}
	order := &models.Order{}
	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		row0, _ := strconv.Atoi(row[0])
		customerID := int64(row0)
		row1, _ := strconv.Atoi(row[1])
		employeeID := int64(row1)
		row3, _ := strconv.Atoi(row[3])
		shippingMethodID := int64(row3)
		freightCharge, _ := strconv.ParseFloat(row[4], 64)
		taxes, _ := strconv.ParseFloat(row[5], 64)

		order.CustomerID = &customerID
		order.EmployeeID = &employeeID
		order.PurchaseOrderNumber = row[2]
		order.ShippingMethodID = &shippingMethodID
		order.FreightCharge = &freightCharge
		order.Taxes = &taxes
		if err := dbInstance.AddOrder(order); err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}

		result = append(result, order.OrderID)
	}

	bulkOrderRes := &models.BulkOrderRes{}
	bulkOrderRes.IDs = result
	if err := render.Render(w, r, bulkOrderRes); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

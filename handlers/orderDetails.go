package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/wldnist/jatis-simple-pos/helpers"
	"github.com/wldnist/jatis-simple-pos/models"
)

func orderDetails(router chi.Router) {
	router.Post("/bulk", createBulkOrderDetails)
}

func createBulkOrderDetails(w http.ResponseWriter, r *http.Request) {
	bulkOrderDetailReq := &models.BulkOrderDetailReq{}
	if err := render.Bind(r, bulkOrderDetailReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	data, err := helpers.ReadCSVFromUrl(bulkOrderDetailReq.Url)
	if err != nil {
		panic(err)
	}

	result := []int64{}
	orderDetail := &models.OrderDetail{}
	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		row0, _ := strconv.Atoi(row[0])
		orderID := int64(row0)
		row1, _ := strconv.Atoi(row[1])
		productID := int64(row1)
		row2, _ := strconv.Atoi(row[2])
		unitPrice, _ := strconv.ParseFloat(row[3], 64)
		discount, _ := strconv.ParseFloat(row[4], 64)

		orderDetail.OrderID = &orderID
		orderDetail.ProductID = &productID
		orderDetail.Quantity = &row2
		orderDetail.UnitPrice = &unitPrice
		orderDetail.Discount = &discount
		if err := dbInstance.AddOrderDetail(orderDetail); err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}

		result = append(result, orderDetail.OrderDetailID)
	}

	bulkOrderDetailRes := &models.BulkOrderDetailRes{}
	bulkOrderDetailRes.IDs = result
	if err := render.Render(w, r, bulkOrderDetailRes); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

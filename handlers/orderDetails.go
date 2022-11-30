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
	router.Get("/complete", getCompleteAllOrderDetails)
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

func getCompleteAllOrderDetails(w http.ResponseWriter, r *http.Request) {
	getCompleteAllOrderDetails, err := dbInstance.GetCompleteAllOrderDetails()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	getCompleteAllOrderDetailsRes := &models.GetCompleteAllOrderDetailsRes{}
	orderIDs := make(map[int64]bool)
	for _, row := range getCompleteAllOrderDetails {
		if exist := orderIDs[row.OrderID]; !exist {
			orderIDs[row.OrderID] = true
		}
	}

	result := []models.CompleteOrderDetails{}
	for orderID, _ := range orderIDs {
		completeOrderDetails := models.CompleteOrderDetails{}
		orderDetails := []models.OrderDetailsSpec{}
		customerName := ""
		employeeName := ""
		shippingMethod := ""
		totalPayment := float64(0)
		for _, getCompleteAllOrderDetail := range getCompleteAllOrderDetails {
			if getCompleteAllOrderDetail.OrderID == orderID {
				customerName = getCompleteAllOrderDetail.CustomerName
				employeeName = getCompleteAllOrderDetail.EmployeeName
				shippingMethod = getCompleteAllOrderDetail.ShippingMethod
				totalPayment += getCompleteAllOrderDetail.SubTotal
				orderDetailsSpec := models.OrderDetailsSpec{
					OrderID:     getCompleteAllOrderDetail.OrderID,
					ProductID:   getCompleteAllOrderDetail.ProductID,
					ProductName: getCompleteAllOrderDetail.ProductName,
					Quantity:    getCompleteAllOrderDetail.Quantity,
					UnitPrice:   getCompleteAllOrderDetail.UnitPrice,
					Discount:    getCompleteAllOrderDetail.Discount,
					SubTotal:    getCompleteAllOrderDetail.SubTotal,
				}

				orderDetails = append(orderDetails, orderDetailsSpec)
			}
		}

		completeOrderDetails.CustomerName = customerName
		completeOrderDetails.EmployeeName = employeeName
		completeOrderDetails.ShippingMethod = shippingMethod
		completeOrderDetails.TotalPayment = totalPayment
		completeOrderDetails.OrderDetails = orderDetails

		result = append(result, completeOrderDetails)
	}

	getCompleteAllOrderDetailsRes.Result = result
	if err := render.Render(w, r, getCompleteAllOrderDetailsRes); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

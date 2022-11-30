package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/wldnist/jatis-simple-pos/helpers"
	"github.com/wldnist/jatis-simple-pos/models"
)

func products(router chi.Router) {
	router.Post("/bulk", createBulkProducts)
}

func createBulkProducts(w http.ResponseWriter, r *http.Request) {
	bulkProductReq := &models.BulkProductReq{}
	if err := render.Bind(r, bulkProductReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	data, err := helpers.ReadCSVFromUrl(bulkProductReq.Url)
	if err != nil {
		panic(err)
	}

	result := []int64{}
	product := &models.Product{}
	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		unitPrice := float64(0)
		inStock := int(0)
		unitPrice, _ = strconv.ParseFloat(row[1], 64)
		inStock, _ = strconv.Atoi(row[2])

		product.ProductName = row[0]
		product.UnitPrice = &unitPrice
		product.InStock = inStock
		if err := dbInstance.AddProduct(product); err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}

		result = append(result, product.ProductID)
	}

	bulkProductRes := &models.BulkProductRes{}
	bulkProductRes.IDs = result
	if err := render.Render(w, r, bulkProductRes); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

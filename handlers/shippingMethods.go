package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/wldnist/jatis-simple-pos/helpers"
	"github.com/wldnist/jatis-simple-pos/models"
)

func shippingMethods(router chi.Router) {
	router.Post("/bulk", createBulkShippingMethods)
}

func createBulkShippingMethods(w http.ResponseWriter, r *http.Request) {
	bulkShippingMethodReq := &models.BulkShippingMethodReq{}
	if err := render.Bind(r, bulkShippingMethodReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	data, err := helpers.ReadCSVFromUrl(bulkShippingMethodReq.Url)
	if err != nil {
		panic(err)
	}

	result := []int64{}
	shippingMethod := &models.ShippingMethod{}
	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		shippingMethod.ShippingMethod = row[0]
		if err := dbInstance.AddShippingMethod(shippingMethod); err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}

		result = append(result, shippingMethod.ShippingMethodID)
	}

	bulkShippingMethodRes := &models.BulkShippingMethodRes{}
	bulkShippingMethodRes.IDs = result
	if err := render.Render(w, r, bulkShippingMethodRes); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

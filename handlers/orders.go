package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func orders(router chi.Router) {
	router.Get("/", getAllOrders)
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

package product

import (
	"net/http"

	"github.com/AmiyoKm/basic_http/middleware"
)

func (h *Handler) HttpRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /api/products", manager.With(http.HandlerFunc(h.getProducts), manager.Authentication))
	mux.Handle("GET /api/products/{id}", manager.With(http.HandlerFunc(h.getProductByID), manager.Authentication))
	mux.Handle("POST /api/products", manager.With(http.HandlerFunc(h.createProduct), manager.Authentication))
	mux.Handle("PUT /api/products/{id}", manager.With(http.HandlerFunc(h.updateProduct), manager.Authentication))
	mux.Handle("DELETE /api/products/{id}", manager.With(http.HandlerFunc(h.deleteProduct), manager.Authentication))
}

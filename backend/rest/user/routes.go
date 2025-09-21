package user

import (
	"net/http"

	"github.com/AmiyoKm/basic_http/middleware"
)

func (h *Handler) HttpRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /api/users/register", manager.With(http.HandlerFunc(h.RegisterUser)))
	mux.Handle("POST /api/users/login", manager.With(http.HandlerFunc(h.LoginUser)))
}

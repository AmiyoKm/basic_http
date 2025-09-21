package middleware

import (
	"net/http"

	"github.com/AmiyoKm/basic_http/config"
)

type Manager struct {
	middlewares []Middleware
	cfg         *config.Config
}

func NewManager(cfg *config.Config) *Manager {
	return &Manager{
		middlewares: make([]Middleware, 0),
		cfg:         cfg,
	}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {

	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}

	for i := len(m.middlewares) - 1; i >= 0; i-- {
		next = m.middlewares[i](next)
	}

	return next
}

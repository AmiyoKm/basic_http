package main

import "net/http"

type Manager struct {
	middlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		middlewares: make([]Middleware, 0),
	}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {


	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}

	for _, middleware := range m.middlewares {
		next = middleware(next)
	}

	return next
}

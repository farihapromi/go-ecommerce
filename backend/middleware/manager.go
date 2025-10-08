package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler
type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}
func (mngr *Manager) With(middleware ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		h := handler

	}

}

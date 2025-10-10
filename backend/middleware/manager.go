package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler
type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}
func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)

}
func (mngr *Manager) With(middlewares []Middleware, handler http.Handler) http.Handler {

	h := handler
	//middlewares =[hudai,logger] [0,1]
	// middleware.Logger http.HandlerFunc(handlers.GetProducts)))
	// for i := len(middlewares) - 1; i >= 0; i-- {
	// 	middleware := middlewares[i]
	// 	n = middleware(n)

	// }
	//h=corsWithPrefligh(logger(Mux))
	for _, middleware := range middlewares {
		h = middleware(h)

	}
	// for _, globalMiddleware := range mngr.globalMiddlewares {
	// 	h = globalMiddleware(h)
	// }

	return h

}
func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler
	//middlewares =[hudai,logger] [0,1]
	// middleware.Logger http.HandlerFunc(handlers.GetProducts)))
	// for i := len(middlewares) - 1; i >= 0; i-- {
	// 	middleware := middlewares[i]
	// 	n = middleware(n)

	// }
	//h=corsWithPrefligh(logger(Mux))
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)

	}
	// for _, globalMiddleware := range mngr.globalMiddlewares {
	// 	h = globalMiddleware(h)
	// }

	return h

}

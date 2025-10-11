package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /rahim", manager.WrapMux(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /route", manager.WrapMux(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /products", manager.WrapMux(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.WrapMux(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productId}", manager.WrapMux(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("PUT /products/{productId}", manager.WrapMux(http.HandlerFunc(handlers.UpdateProducts)))

}

package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /rahim", manager.With(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /route", manager.With(http.HandlerFunc(handlers.Test)))

	mux.Handle(
		"GET /products",
		manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProducts)))
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProducts)))

}

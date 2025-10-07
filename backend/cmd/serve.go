package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"

	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET / route", middleware.Logger(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productId}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))

	fmt.Println("Server running on :8080")

	// Use cors-enabled globalRouter
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

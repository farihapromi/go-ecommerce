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
	manager := middleware.NewManager()

	mux.Handle("GET /rahim", manager.With(http.HandlerFunc(handlers.Test),
		middleware.Logger))

	mux.Handle("GET / route", manager.With(http.HandlerFunc(handlers.Test),
		middleware.Logger))

	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts), middleware.Logger))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middleware.Logger))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(handlers.GetProductByID), middleware.Logger))

	fmt.Println("Server running on :8080")

	// Use cors-enabled globalRouter
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

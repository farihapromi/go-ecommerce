package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	cntrl := func(w http.ResponseWriter, r *http.Request) {
		log.Println("ami handler:middle a print hobo")
	}
	handler := http.HandlerFunc(cntrl)
	loggerMiddlerware := middleware.Logger(handler)
	mux.Handle("GET route", loggerMiddlerware)

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server running on :8080")

	// Use cors-enabled globalRouter
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

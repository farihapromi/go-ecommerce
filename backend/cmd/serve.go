package cmd

import (
	"ecommerce/middleware"
	"fmt"

	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.CorsWithPreflight)
	initRoutes(mux, manager)

	fmt.Println("Server running on :8080")

	// Use cors-enabled globalRouter
	// globalRouter := middleware.CorsWithPreflight(mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

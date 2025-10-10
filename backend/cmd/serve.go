package cmd

import (
	"ecommerce/middleware"
	"fmt"

	"net/http"
)

func Serve() {

	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.CorsWithPreflight)
	mux := http.NewServeMux()
	initRoutes(mux, manager)

	fmt.Println("Server running on :8080")

	// Use cors-enabled globalRouter
	// globalRouter := middleware.CorsWithPreflight(mux)
	wrappedMux := manager.WrapMux(mux)
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

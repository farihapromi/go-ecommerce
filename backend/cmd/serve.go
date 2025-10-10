package cmd

import (
	"ecommerce/middleware"
	"fmt"

	"net/http"
)

func Serve() {

	manager := middleware.NewManager()
	// manager.Use(middleware.Logger, middleware.CorsWithPreflight)
	mux := http.NewServeMux()
	initRoutes(mux, manager)
	// globalMiddlewares := []middleware.Middleware{
	// 	middleware.CorsWithPreflight, middleware.Logger,
	// }m
	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	fmt.Println("Server running on :8080")

	// Use cors-enabled globalRouter
	// globalRouter := middleware.CorsWithPreflight(mux)
	//CorsWithPreflight(hudai(logger(mux)))
	// wrappedMux := manager.WrapMux(mux, middleware.Logger, middleware.CorsWithPreflight)

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

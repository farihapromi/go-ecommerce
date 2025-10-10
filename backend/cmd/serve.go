package cmd

import (
	"ecommerce/middleware"
	"fmt"

	"net/http"
)

func Serve() {

	manager := middleware.NewManager()

	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	mux := http.NewServeMux()
	initRoutes(mux, manager)
	wrappedMux := manager.WrapMux(mux)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

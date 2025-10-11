package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"os"
	"strconv"

	"net/http"
)

func Serve() {
	cnf := config.GetConfig()

	manager := middleware.NewManager()

	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	mux := http.NewServeMux()
	initRoutes(mux, manager)
	wrappedMux := manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(cnf.HttpPort) //type casting .int to asci
	fmt.Println("Server running on port ", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}

}

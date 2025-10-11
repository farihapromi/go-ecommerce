package config

import (
	"fmt"
	"os"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

// load configuration variable from env
func LoadConfig() {
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Sevice name is required")
		os.Exit(1)

	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)

	}

}

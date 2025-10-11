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

func LoadConfig() {
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("vSevice name is required")

	}

}

package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int64
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
	port, err := strconv.ParseInt(httpPort, 10, 64) //decimasl base 10, 64 base
	if err != nil {
		fmt.Println("port must be integar:", err)
		os.Exit(1)
	}
	cnf := Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}

}

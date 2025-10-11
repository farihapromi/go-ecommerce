package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

// data segment a thalbe  const
// code segment a thakbe var
var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

//keep the variable lowercae so that other package can not access this
//capital hole public

// load configuration variable from env
func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env variables", err)
		os.Exit(1)
	}
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
	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
	}

}
func GetConfig() Config {
	loadConfig()
	return configurations
}

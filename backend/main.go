package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
	"fmt"
)

func main() {
	cnf := config.GetConfig()
	fmt.Println(cnf.HttpPort)
	fmt.Println(cnf.ServiceName)
	cmd.Serve()
}

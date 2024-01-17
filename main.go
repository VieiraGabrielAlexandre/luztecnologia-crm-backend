package main

import (
	"fmt"

	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/config"
	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/routes"
)

func main() {
	fmt.Printf("-----------\n")
	fmt.Printf("Loading ENV VARS\n")
	config.SetEnv()
	fmt.Printf("-----------\n")

	fmt.Printf("-----------\n")
	fmt.Printf("Loading DATABASE\n")
	config.ConnectDatabase()
	fmt.Printf("-----------\n")

	fmt.Printf("-----------\n")
	fmt.Printf("Loading Routes\n")
	routes.HandleRequests()
	fmt.Printf("-----------\n")
}

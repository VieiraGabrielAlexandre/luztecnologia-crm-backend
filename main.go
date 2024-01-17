package main

import (
	"fmt"
	"sync"

	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/config"
	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/routes"
)

var wg sync.WaitGroup

func main() {

	fmt.Printf("-----------\n")
	fmt.Printf("Loading Routes\n")
	wg.Add(1)
	go carregarRotas()
	fmt.Printf("-----------\n")

	fmt.Printf("-----------\n")
	fmt.Printf("Loading ENV VARS\n")
	config.SetEnv()
	fmt.Printf("-----------\n")

	fmt.Printf("-----------\n")
	fmt.Printf("Loading DATABASE\n")
	config.ConnectDatabase()
	fmt.Printf("-----------\n")
	wg.Wait()
}
func carregarRotas() {
	var wg sync.WaitGroup
	wg.Add(1)
	routes.HandleRequests()
	wg.Done()
}

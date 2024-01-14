package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func SetEnv() {
	godotenv.Load("./.env")
	fmt.Println("Success")
}

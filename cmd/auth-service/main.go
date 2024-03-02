package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sladkoezhkovo/auth-service/internal/configs"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/.yml", "path to config")
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading environment")
	}

	flag.Parse()

	config := configs.SetupConfig(configPath)

	fmt.Println(config)
}

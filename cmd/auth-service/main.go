package main

import (
	"flag"
	"fmt"
	"github.com/sladkoezhkovo/auth-service/internal/configs"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/.yml", "path to config")
}

func main() {
	flag.Parse()

	config := configs.SetupConfig(configPath)

	fmt.Println(config)
}

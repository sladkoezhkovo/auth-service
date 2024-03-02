package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/sladkoezhkovo/auth-service/internal/configs"
	jwtservice "github.com/sladkoezhkovo/auth-service/internal/service/jwt-service"
	"github.com/sladkoezhkovo/auth-service/internal/storage/redis"
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

	redisStorage := redis.New(&config.Redis)
	_ = jwtservice.New(&config.Jwt, redisStorage)
}

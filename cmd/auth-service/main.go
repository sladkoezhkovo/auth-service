package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	api "github.com/sladkoezhkovo/auth-service/api/auth"
	"github.com/sladkoezhkovo/auth-service/internal/configs"
	"github.com/sladkoezhkovo/auth-service/internal/grpc/auth"
	jwtservice "github.com/sladkoezhkovo/auth-service/internal/service/jwt-service"
	roleservice "github.com/sladkoezhkovo/auth-service/internal/service/role-service"
	userservice "github.com/sladkoezhkovo/auth-service/internal/service/user-service"
	"github.com/sladkoezhkovo/auth-service/internal/storage/pg"
	"github.com/sladkoezhkovo/auth-service/internal/storage/redis"
	"github.com/sladkoezhkovo/lib"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	configPath string
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
)

func init() {
	flag.StringVar(&configPath, "config", "configs/.yml", "path to config")
}

func main() {

	flag.Parse()

	var config configs.Config

	if err := lib.SetupConfig(configPath, &config); err != nil {
		panic(fmt.Errorf("cannot read config: %s", err))
	}

	if config.App.Env == EnvLocal {
		fmt.Printf("loading dotenv\n")
		if err := godotenv.Load(".env.jwt", ".env.pg", ".env.redis"); err != nil {
			panic(fmt.Errorf("godotenv.Load: %s", err))
		}
	}

	fmt.Println(config)

	db, err := pg.Setup(&config.Pg)
	if err != nil {
		panic(fmt.Errorf("error connecting database: %s", err))
	}
	fmt.Println("postgres connected")

	userRepo := pg.NewUserRepository(db)

	//TODO ADD CRUD FOR ROLES
	roleRepo := pg.NewRoleRepository(db)

	redisStorage := redis.New(&config.Redis)

	jwtService := jwtservice.New(&config.Jwt, redisStorage)
	userService := userservice.New(userRepo)
	roleService := roleservice.New(roleRepo)

	server := grpc.NewServer()
	handler := auth.NewServer(userService, jwtService, roleService)

	api.RegisterAuthServiceServer(server, handler)

	go func(s *grpc.Server) {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.App.Port))
		if err != nil {
			panic(fmt.Errorf("cannot bind port %d", config.App.Port))
		}

		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}(server)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	<-stopChan

	server.GracefulStop()
}

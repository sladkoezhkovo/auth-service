build:
	go build -v ./cmd/auth-service

debug:
	make build
	.\auth-service.exe -config configs/local.yml

protoc:
	protoc proto/auth.proto --go_out=. --go-grpc_out=.

deploy:
	 docker-compose build && docker-compose down && docker-compose up -d && make up

up:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5436/auth?sslmode=disable' up

down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5436/auth?sslmode=disable' down
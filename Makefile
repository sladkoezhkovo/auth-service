build:
	go build -v ./cmd/auth-service

protoc:
	protoc proto/auth.proto --go_out=. --go-grpc_out=.

deploy:
	 docker-compose build && docker-compose down && docker-compose up -d
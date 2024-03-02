build:
	go build -v ./cmd/auth-service

deploy:
	 docker-compose build && docker-compose down && docker-compose up -d
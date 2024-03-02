FROM golang:1.22.0-alpine3.19 AS builder

RUN apk update --no-cache
RUN go mod download
COPY . .
RUN go build -o main ./cmd/app.go

FROM alpine

RUN apk update --no-cache
COPY --from=builder ./main ./app
COPY ./configs ./configs

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /app/wait
RUN chmod +x /app/wait

CMD ./wait && ./main
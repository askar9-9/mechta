# syntax=docker/dockerfile:1
FROM golang:1.23-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем бинарник из cmd/main.go
RUN go build -o app ./cmd/main.go

EXPOSE 8080 8081

CMD ["./app"]

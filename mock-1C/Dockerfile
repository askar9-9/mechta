FROM golang:1.23.3-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o mock1c ./cmd/main.go

EXPOSE 9900
ENTRYPOINT ["./mock1c"]

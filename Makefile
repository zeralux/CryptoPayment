LOCATION:= bin

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest

build: swagger
	go mod tidy
	swag init
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/main main.go

start: build
	go run main.go




LOCATION:= bin

build:
	go mod tidy
	swag init
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/main main.go

start: build
	go run main.go




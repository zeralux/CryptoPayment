LOCATION:= bin

refresh:
	go mod tidy

swagger: refresh
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init

build: refresh swagger
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/main main.go

start: build
	go run main.go




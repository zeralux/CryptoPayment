LOCATION:= bin
YML:=application-local.yml

install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/krashanoff/copypasta@latest

refresh_mod:
	go mod tidy

refresh_swagger:
	swag init

build: install refresh_mod refresh_swagger
	copypasta -i $(YML) -o $(LOCATION)/$(YML) -v
	go build -ldflags="-s -w" -o $(LOCATION)/main main.go

start: build
	go run main.go




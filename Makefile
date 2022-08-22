LOCATION:= bin
YML:=application-local.yml

install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/krashanoff/copypasta@latest

refresh:
	swag init
	go mod tidy

build: install refresh
	copypasta -i $(YML) -o $(LOCATION)/$(YML) -v
	go build -ldflags="-s -w" -o $(LOCATION)/main main.go

start: build
	go run main.go




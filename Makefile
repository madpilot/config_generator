build:
	go build -o bin/generate_config cmd/generate_config/main.go

default:
	build

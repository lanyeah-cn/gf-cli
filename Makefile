
.PHONY: default

default:
	@echo "USAGE: make build"


build:
	go build -o ./gf-cli main.go
	

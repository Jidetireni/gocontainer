# Makefile for gocontainer project

.PHONY: all setup build run clean examples example-simple example-network test

all: setup build

setup:
	sudo scripts/setup.sh

build:
	go build -o bin/gocontainer ./src/main.go
	go build -o bin/simple-container ./examples/simple-container.go
	go build -o bin/network-container ./examples/network-container.go

run: build
	sudo ./bin/gocontainer /bin/bash

example-simple: build
	sudo ./bin/simple-container

example-network: build
	sudo ./bin/network-container

test:
	go test -v ./src/...

clean:
	rm -f bin/gocontainer
	rm -f bin/simple-container
	rm -f bin/network-container
	
# Usage instructions
help:
	@echo "Make targets:"
	@echo "  all           - Setup environment and build binaries"
	@echo "  setup         - Set up container environment (requires sudo)"
	@echo "  build         - Build all binaries"
	@echo "  run           - Run the main container with bash"
	@echo "  example-simple- Run simple container example"
	@echo "  example-network- Run network namespace container example"
	@echo "  clean         - Remove built binaries"
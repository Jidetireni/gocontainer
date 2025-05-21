# Makefile for gocontainer project

.PHONY: build run clean

build:
	go build -o gocontainer ./src/main.go

run: build
	./gocontainer

clean:
	rm -f gocontainer
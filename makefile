.PHONY: build
build:
	go build -v -o ./bin/ ./cmd/

run:
	make build
	clear
	./bin/cmd

DEFAULT_GOAL: build
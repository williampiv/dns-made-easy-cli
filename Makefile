NAME := dns-made-easy-cli
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

check: vet lint
.PHONY:check

build:
	go build -o build/$(NAME)
.PHONY: build

build-all:
	GOOS=linux GOARCH=amd64 go build -o build/$(NAME)-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o build/$(NAME)-linux-arm64
	GOOS=darwin GOARCH=amd64 go build -o build/$(NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o build/$(NAME)-darwin-arm64
	GOOS=windows GOARCH=amd64 go build -o build/$(NAME)-windows-amd64
.PHONY: build-all

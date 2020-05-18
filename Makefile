.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run the app
	go run ./cmd/sokoban/

build: ## Build binary in the local env
	go build -mod=vendor -i -v -o sokoban-bin ./cmd/sokobango get /

govet: ## Run go vet on the project
	go vet ./...

test: ## Run go tests
	go test -race -v ./...

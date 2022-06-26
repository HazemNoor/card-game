include .env
export

.PHONY: help

.DEFAULT_GOAL := help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run the tests
	go test -race -v $(shell go list ./...)

build: ## Build the binary
	go build -o bin/$(APP_NAME) cmd/api/*

start: build ## Build and Start the binary
	bin/$(APP_NAME) -http-address=$(HTTP_ADDRESS) -redis-address=$(REDIS_HOST):$(REDIS_PORT) -redis-password=$(REDIS_PASSWORD)

start-redis: ## Start redis server
	docker run --name $(APP_NAME)-redis -p $(REDIS_PORT):6379 -d redis:7

stop-redis: ## Stop redis server
	docker stop $(APP_NAME)-redis && docker rm $(APP_NAME)-redis

restart-redis: stop-redis start-redis ## Restart redis server

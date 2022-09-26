#!make
ifneq (,$(wildcard ./build/.env))
    include ./build/.env
    export $(shell sed 's/=.*//' ./build/.env)
endif

lint:
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.49.0 golangci-lint run

run:
	go run cmd/main.go

docker-run:
	COMPOSE_HTTP_TIMEOUT=120 docker-compose -f build/dev/docker-compose.yml up

docker-down:
	COMPOSE_HTTP_TIMEOUT=120 docker-compose -f build/dev/docker-compose.yml down

init:
	docker network create potaynik-net
	docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions

deps:
	go mod tidy
	go mod vendor

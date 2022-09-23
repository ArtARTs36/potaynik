#!make
ifneq (,$(wildcard ./build/.env))
    include ./build/.env
    export $(shell sed 's/=.*//' ./build/.env)
endif

lint:
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.49.0 golangci-lint run

run:
	docker-compose -f build/dev/docker-compose.yml up

docker-create-network:
	docker network create potaynik-net

deps:
	go mod tidy
	go mod vendor

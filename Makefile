#!make
ifneq (,$(wildcard ./build/.env))
    include ./build/.env
    export $(shell sed 's/=.*//' ./build/.env)
endif

run:
	go run cmd/main.go

docker-run:
	docker-compose -f build/dev/docker-compose.yml up

docker-create-network:
	docker network create potaynik-net

deps:
	go mod tidy
	go mod vendor

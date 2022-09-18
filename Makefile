#!make
ifneq (,$(wildcard ./build/.env))
    include ./build/.env
    export $(shell sed 's/=.*//' ./build/.env)
endif

run:
	go run cmd/main.go

deps:
	go mod tidy
	go mod vendor

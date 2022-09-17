run:
	go run cmd/main.go

deps:
	go mod tidy
	go mod vendor

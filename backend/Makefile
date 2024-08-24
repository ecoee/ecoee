GIT_SHA := $(shell git rev-parse --short HEAD)
.PHONY: generate format test lint cover

generate:
	go generate ./...

format:
	go fmt ./...
	goimports -w .

test:
	go test -v ./...

build:
	docker build -t ecoee:$(GIT_SHA) .

run:
	docker run -p 8080:8080 ecoee:$(GIT_SHA)

lint:
	docker compose -f docker-compose-lint.yaml up
	docker compose -f docker-compose-lint.yaml down

cover:
	go test -p=1 -race ./... -coverprofile=coverage.out -covermode=atomic

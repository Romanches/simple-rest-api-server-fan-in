.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY: all lint test build run clean mod vendor

all: run

vendor:
	go mod vendor

# go linter
lint:
	golangci-lint run

test:
	go test ./... -race

test-coverage:
	go test ./... -race -coverprofile /tmp/coverage.out && go tool cover -html=/tmp/coverage.out

make mock-server:
	go run cmd/mock_server/main.go

# run local
run:
	go run ./cmd/.

build: vendor
	CGO_ENABLED=0 go build -o go-app ./cmd/.

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-app ./cmd/.

clean:
	go clean ./...
	rm -f go-app

rebuild: clean build

docker-build: build-linux
	docker build -t go-app .


docker-run: docker-build
	docker run -tid --rm go-app


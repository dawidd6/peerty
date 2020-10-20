VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse HEAD)

build:
	@mkdir -p bin
	go build -ldflags "-s -w -X github.com/dawidd6/p2p/pkg/version.Version=$(VERSION)" -o bin ./cmd/...

test:
	go test -v -count=1 ./...

proto:
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/*/*.proto
	protoc --go_out=. --go_opt=paths=source_relative pkg/*/*.proto

image:
	docker build -t p2p .
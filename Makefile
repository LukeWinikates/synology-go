.PHONY: test integration-test setup
build/synoctl: build/ $(shell find . -iname "*.go")
	go build -o $@ cmd/main.go cmd/version.go

build/:
	-mkdir -p build

test:
	go test -v ./...
	golangci-lint run

integration-test:
	go test -v ./...  --tags=integration

setup:
	 brew install golangci-lint

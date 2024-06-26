.PHONY: test integration-test setup
build/synoctl: build/ $(shell find . -iname "*.go")
	go build -o $@ cmd/main.go

build/:
	-mkdir -p build

test:
	golangci-lint run
	go test -v ./...

integration-test:
	go test -v ./...  --tags=integration

setup:
	 brew install golangci-lint

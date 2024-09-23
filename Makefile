all: test build

build:
	go mod tidy
	go build .

test:
	go test -v ./...

default: install
	@GO111MODULE=on go build -o bin/packet-mon ./main.go

install:
	@go mod download

test: install
	@go test -v ./...

.PHONY: default install test


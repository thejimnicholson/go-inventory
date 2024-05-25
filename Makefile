.DEFAULT_GOAL := bin/go-inventory

bin/go-inventory: $(wildcard *.go)
	go build -v -ldflags="-X main.Version=$(shell git describe --always --long --dirty=+)" -o bin/go-inventory cmd/go-inventory.go

test:
	go test -v ./...

clean:
	rm bin/go-inventory

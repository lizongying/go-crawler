.PHONY: all

all: tidy tls testSpider

module := $(shell head -n 1 go.mod)
module := $(subst module ,,${module})

tidy:
	go mod tidy

tls:
	go vet ./cmd/tls
	go build -ldflags "-s -w" -o ./releases/tls ./cmd/tls

testSpider:
	go vet ./cmd/testSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test" -o ./releases/testSpider ./cmd/testSpider

.PHONY: all

all: tidy tls mitm test_spider test_compress_spider test_decode_spider test_file_spider test_item_spider multi_spider

module := $(shell head -n 1 go.mod)
module := $(subst module ,,${module})

shell:
	@echo 'SHELL='$(SHELL)

tidy:
	go mod tidy

tls:
	go vet ./tools/tls
	go build -ldflags "-s -w" -o ./releases/tls ./tools/tls

mitm:
	go vet ./tools/mitm
	go build -ldflags "-s -w" -o ./releases/mitm ./tools/mitm

test_spider:
	go vet ./cmd/test_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test" -o ./releases/test_spider ./cmd/test_spider

test_compress_spider:
	go vet ./cmd/test_compress_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-compress" -o ./releases/test_compress_spider ./cmd/test_compress_spider

test_decode_spider:
	go vet ./cmd/test_decode_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-decode" -o ./releases/test_decode_spider ./cmd/test_decode_spider

test_file_spider:
	go vet ./cmd/test_file_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-file" -o ./releases/test_file_spider ./cmd/test_file_spider

test_item_spider:
	go vet ./cmd/test_item_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-item" -o ./releases/test_item_spider ./cmd/test_item_spider

multi_spider:
	go vet ./cmd/multi_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-item" -o ./releases/multi_spider ./cmd/multi_spider
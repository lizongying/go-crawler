.PHONY: all

all: tidy tls mitm testSpider testCompressSpider testDecodeSpider testFileSpider testItemSpider

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

testSpider:
	go vet ./cmd/testSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test" -o ./releases/testSpider ./cmd/testSpider

testCompressSpider:
	go vet ./cmd/testCompressSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-compress" -o ./releases/testCompressSpider ./cmd/testCompressSpider

testDecodeSpider:
	go vet ./cmd/testDecodeSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-decode" -o ./releases/testDecodeSpider ./cmd/testDecodeSpider

testFileSpider:
	go vet ./cmd/testFileSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-file" -o ./releases/testFileSpider ./cmd/testFileSpider

testItemSpider:
	go vet ./cmd/testItemSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-item" -o ./releases/testItemSpider ./cmd/testItemSpider
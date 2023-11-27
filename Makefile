.PHONY: all

all: web_ui web_server tidy tls_generator mitm test_spider test_compress_spider test_decode_spider test_file_spider test_item_spider multi_spider

module := $(shell head -n 1 go.mod)
module := $(subst module ,,${module})
branch := $(shell git rev-parse --abbrev-ref HEAD)
commit := $(shell git rev-parse --short HEAD)
commit_time := $(shell git log -1 --format=%ct)

shell:
	@echo 'SHELL='$(SHELL)
	@echo 'branch='$(branch)
	@echo 'commit='$(commit)
	@echo 'commit_time='$(commit_time)

tidy:
	go mod tidy

tls_generator:
	go vet ./tools/tls_generator
	go build -ldflags "-s -w" -o ./releases/tls_generator ./tools/tls_generator

tls_generator_more:
	go vet ./tools/tls_generator
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/tls_generator_linux_amd64 ./tools/tls_generator
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/tls_generator_linux_arm64 ./tools/tls_generator
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/tls_generator_darwin_amd64 ./tools/tls_generator
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/tls_generator_darwin_arm64 ./tools/tls_generator
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/tls_generator_windows_amd64.exe ./tools/tls_generator

mitm:
	go vet ./tools/mitm
	go build -ldflags "-s -w" -o ./releases/mitm ./tools/mitm

test_spider:
	go vet ./cmd/test_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/test_spider ./cmd/test_spider

test_compress_spider:
	go vet ./cmd/test_compress_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-compress -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/test_compress_spider ./cmd/test_compress_spider

test_decode_spider:
	go vet ./cmd/test_decode_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-decode -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/test_decode_spider ./cmd/test_decode_spider

test_file_spider:
	go vet ./cmd/test_file_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-file -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/test_file_spider ./cmd/test_file_spider

test_item_spider:
	go vet ./cmd/test_item_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-item -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/test_item_spider ./cmd/test_item_spider

multi_spider:
	go vet ./cmd/multi_spider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test-item -X $(module)/pkg/crawler.buildBranch=$(branch) -X $(module)/pkg/crawler.buildCommit=$(commit) -X $(module)/pkg/crawler.buildTime=$(commit_time)" -o ./releases/multi_spider ./cmd/multi_spider

web_ui:
	rm -rf ./static/dist
	npm run build --prefix ./web/ui

web_server:
	go vet ./tools/web_server
	go build -ldflags "-s -w" -o ./releases/web_server ./tools/web_server

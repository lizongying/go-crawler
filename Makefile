.PHONY: all

all: tidy testSpider

module := $(shell head -n 1 go.mod)
module := $(subst module ,,${module})

tidy:
	go mod tidy

testSpider:
	go vet ./example/testSpider
	go build -ldflags "-s -w -X $(module)/pkg/logger.name=test" -o ./releases/testSpider ./example/testSpider

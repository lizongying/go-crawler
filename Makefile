.PHONY: all

all:  youtubeSpider


youtubeSpider:
	go mod tidy
	go vet ./example/youtubeSpider
	go build -ldflags "-s -w" -o ./releases/youtubeSpider ./cmd/youtubeSpider

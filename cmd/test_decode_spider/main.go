package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/test_decode_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(test_decode_spider.NewSpider).Run()
}

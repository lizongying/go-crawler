package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/test_compress_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(test_compress_spider.NewSpider).Run()
}

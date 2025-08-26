package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/test_httpbin_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(test_httpbin_spider.NewSpider).Run()
}

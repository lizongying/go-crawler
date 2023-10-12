package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/test_method_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(test_method_spider.NewSpider).Run()
}

package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/test_scheduler_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(test_scheduler_spider.NewSpider).Run()
}

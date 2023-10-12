package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/test_file_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(test_file_spider.NewSpider).Run()
}

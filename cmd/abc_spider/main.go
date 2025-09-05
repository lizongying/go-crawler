package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/abc_spider"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewSimpleApp(abc_spider.NewSpider).Run(pkg.WithDefaultMocks())
}

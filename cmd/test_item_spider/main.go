package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/test_item_spider"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
)

func main() {
	app.NewApp(test_item_spider.NewSpider).Run(pkg.WithMockServerRoutes(mock_servers.NewRouteOk))
}

package main

import (
	mockServersInternal "github.com/lizongying/go-crawler/internal/mock_servers"
	"github.com/lizongying/go-crawler/internal/spiders/test_spider"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp(test_spider.NewSpider).Run(pkg.WithMockServerRoutes(mockServersInternal.NewRouteCustom))
}

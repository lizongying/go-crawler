package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/test_error_spider"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
)

func main() {
	app.NewApp(test_error_spider.NewSpider).Run(
		pkg.WithMockServerRoutes(mock_servers.NewRouteOk, mock_servers.NewRouteBadGateway),
	)
}

package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/test_from_request_spider"
	"github.com/lizongying/go-crawler/internal/spiders/test_must_ok_spider"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
)

// go run cmd/multi_spider/*.go -c example.yml
func main() {
	app.NewApp(
		test_from_request_spider.NewSpider,
		test_must_ok_spider.NewSpider,
	).Run(
		pkg.WithMockServerRoutes(mock_servers.NewRouteOk),
	)
}

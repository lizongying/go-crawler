package main

import (
	"github.com/lizongying/go-crawler/internal/spiders/{{.Name}}_spider"
	"github.com/lizongying/go-crawler/pkg/app"
)

func main() {
	app.NewApp({{.Name}}_spider.NewSpider).Run()
}

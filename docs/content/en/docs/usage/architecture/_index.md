---
weight: 1
title: Basic Architecture
---

## Basic Architecture

* Crawler：Within the Crawler, there can be multiple Spiders, and it manages the startup and shutdown of the Spiders.
* Spider: Spider integrates components such as Downloader, Exporter, and Scheduler. In the Spider, you can initiate requests and parse content. You need to set a unique name for each
  Spider.`spider.WithOptions(pkg.WithName("example"))`

    ```go
    package main
    
    import (
        "github.com/lizongying/go-crawler/pkg"
        "github.com/lizongying/go-crawler/pkg/app"
    )
    
    type Spider struct {
        pkg.Spider
        logger pkg.Logger
    }
    
    // some spider funcs
    
    func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
        spider = &Spider{
            Spider: baseSpider,
            logger: baseSpider.GetLogger(),
        }
        spider.WithOptions(
            pkg.WithName("test"),
        )
        return
    }
    
    func main() {
        app.NewApp(NewSpider).Run()
    }

    ```

* Job
* Task
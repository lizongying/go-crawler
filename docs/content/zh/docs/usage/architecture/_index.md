---
weight: 1
title: 基本架构
---

## 基本架构

* Crawler：Crawler里可以有多个Spider，同时管理Spider的启动和关闭等。
* Spider：集成了Downloader、Exporter、Scheduler等组件。在Spider里可以发起请求和解析内容。 您需要为每个Spider设置一个唯一名称。
  `spider.WithOptions(pkg.WithName("example"))`

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
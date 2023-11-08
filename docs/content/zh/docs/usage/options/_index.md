---
weight: 2
title: 选项
---

## 选项

选项是指可以在代码中配置的条目

### Spider选项

* `WithName` 设置唯一名称。
* `WithHost` 设置host，用于基于host的过滤或robot.txt的支持。
* `WithPlatforms` 设置浏览器平台。
* `WithBrowsers` 设置浏览器。
* `WithFilter` 设置过滤器。
* `WithDownloader` 设置下载器。
* `WithExporter` 设置导出器。
* `WithMiddleware` 设置中间件。
* `WithStatsMiddleware` 设置统计中间件，用于记录和统计爬虫的性能和运行情况。
* `WithDumpMiddleware` 设置打印中间件，打印request或者response。
* `WithProxyMiddleware` 设置代理中间件，用于使用代理服务器进行爬取。
* `WithRobotsTxtMiddleware` 设置开启robots.txt支持中间件，用于遵守网站的 robots.txt 规则。
* `WithFilterMiddleware` 设置过滤器中间件，用于过滤已处理的请求。
* `WithFileMiddleware` 设置文件中间件，用于处理文件下载请求。
* `WithImageMiddleware` 设置图像中间件，用于处理图像下载请求。
* `WithHttpMiddleware` 设置 HTTP 中间件。
* `WithRetryMiddleware` 设置重试中间件，用于在请求失败时进行自动重试。
* `WithUrlMiddleware` 设置 URL 中间件。
* `WithReferrerMiddleware` 设置 Referrer 中间件，用于自动设置请求的 Referrer 头。
* `WithCookieMiddleware` 设置 Cookie 中间件，用于处理请求和响应中的 Cookie，自动在接下来的请求设置之前的 Cookie。
* `WithRedirectMiddleware` 设置重定向中间件，用于自动处理请求的重定向，跟随重定向链接并获取最终响应。
* `WithChromeMiddleware` 设置 Chrome 中间件，用于模拟 Chrome 浏览器。
* `WithHttpAuthMiddleware` 设置开启HTTP认证中间件，用于处理需要认证的网站。
* `WithCompressMiddleware` 设置压缩中间件，用于处理请求和响应的压缩。当爬虫发送请求或接收响应时，该中间件可以自动处理压缩算法，解压缩请求或响应的内容。
* `WithDecodeMiddleware` 设置解码中间件，用于处理请求和响应的解码操作。该中间件可以处理请求或响应中的编码内容。
* `WithDeviceMiddleware` 设置开启设备模拟中间件。
* `WithCustomMiddleware` 设置自定义中间件，允许用户定义自己的中间件组件。
* `WithRecordErrorMiddleware` 设置错误记录中间件，请求和解析如果出错会被记录。
* `WithPipeline` 设置Pipeline，用于处理爬取的数据并进行后续操作。
* `WithDumpPipeline` 设置打印管道，用于打印待保存的数据。
* `WithFilePipeline` 设置文件管道，用于处理爬取的文件数据，将文件保存到指定位置。
* `WithImagePipeline` 设置图像管道，用于处理爬取的图像数据，将保存图像到指定位置。
* `WithFilterPipeline` 设置过滤器管道，用于过滤爬取过的数据。
* `WithCsvPipeline` 设置 CSV 数据处理管道，将爬取的数据保存为 CSV 格式。
* `WithJsonLinesPipeline` 设置 JSON Lines 数据处理管道，将爬取的数据保存为 JSON Lines 格式。
* `WithMongoPipeline` 设置 MongoDB 数据处理管道，将爬取的数据保存到 MongoDB 数据库。
* `WithSqlitePipeline` 设置 Sqlite 数据处理管道，将爬取的数据保存到 Sqlite 数据库。
* `WithMysqlPipeline` 设置 MySQL 数据处理管道，将爬取的数据保存到 MySQL 数据库。
* `WithKafkaPipeline` 设置 Kafka 数据处理管道，将爬取的数据发送到 Kafka 消息队列。
* `WithCustomPipeline` 设置自定义数据处理管道。
* `WithRetryMaxTimes` 设置请求的最大重试次数。
* `WithRedirectMaxTimes` 设置请求的最大跳转次数。
* `WithTimeout` 设置请求的超时时间。
* `WithInterval` 设置请求的间隔时间。
* `WithOkHttpCodes` 设置正常的HTTP状态码。

### crawler选项

* `WithLogger` 设置日志。
* `WithMockServerRoutes` 设置模拟服务Route，包括内置或自定义的。不需要配置`mock_server.enable: true`
* `WithItemDelay` 设置数据保存间隔。
* `WithItemConcurrency` 设置数据保存并行数量。
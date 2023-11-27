---
weight: 2
title: Options
---

## Options

Options refer to items that can be configured in the code.

### Spider Options

* `WithName`: Set a unique name for the spider.
* `WithHost`: Set the host for filtering based on the host or to support robots.txt.
* `WithPlatforms`: Set the browser platforms.
* `WithBrowsers`: Set the browsers.
* `WithFilter`: Set the filter.
* `WithDownloader`: Set the downloader.
* `WithExporter`: Set the exporter.
* `WithMiddleware`: Set middleware components.
* `WithStatsMiddleware`: Set the statistics middleware to record and monitor the performance and runtime of the
  spider.
* `WithDumpMiddleware`: Set the dump middleware to print requests or responses.
* `WithProxyMiddleware`: Set the proxy middleware to use proxy servers for crawling.
* `WithRobotsTxtMiddleware`: Set the middleware to enable robots.txt support, ensuring compliance with websites'
  robots.txt rules.
* `WithFilterMiddleware`: Set the filter middleware to filter processed requests.
* `WithFileMiddleware`: Set the file middleware to handle file download requests.
* `WithImageMiddleware`: Set the image middleware to handle image download requests.
* `WithHttpMiddleware`: Set the HTTP middleware.
* `WithRetryMiddleware`: Set the retry middleware for automatic retries in case of request failures.
* `WithUrlMiddleware`: Set the URL middleware.
* `WithReferrerMiddleware`: Set the referrer middleware to automatically set the Referrer header for requests.
* `WithCookieMiddleware`: Set the cookie middleware to handle cookies in requests and responses, automatically
  preserving cookies for subsequent requests.
* `WithRedirectMiddleware`: Set the redirect middleware to automatically handle request redirections, following the
  redirect links to obtain the final response.
* `WithChromeMiddleware`: Set the Chrome middleware to simulate the Chrome browser.
* `WithHttpAuthMiddleware`: Enable the HTTP authentication middleware to handle websites that require
  authentication.
* `WithCompressMiddleware`: Set the compress middleware to handle compression in requests and responses. When the
  crawler sends requests or receives responses, this middleware can automatically handle compression algorithms,
  decompressing the content of requests or responses.
* `WithDecodeMiddleware`: Set the decode middleware to handle decoding operations in requests and responses. This
  middleware can handle encoding content in requests or responses.
* `WithDeviceMiddleware`: Enable the device simulation middleware.
* `WithCustomMiddleware`: Set the custom middleware, allowing users to define their own middleware components.
* `WithRecordErrorMiddleware` Set up error logging middleware, request and parsing will be logged if there is an error
* `WithPipeline`: Set the Pipeline to process the crawled data and perform subsequent operations.
* `WithDumpPipeline`: Set the dump pipeline to print data to be saved.
* `WithFilePipeline`: Set the file pipeline to handle crawled file data and save files to a specified location.
* `WithImagePipeline`: Set the image pipeline to handle crawled image data and save images to a specified location.
* `WithFilterPipeline`: Set the filter pipeline to filter crawled data.
* `WithCsvPipeline`: Set the CSV data processing pipeline to save crawled data in CSV format.
* `WithJsonLinesPipeline`: Set the JSON Lines data processing pipeline to save crawled data in JSON Lines format.
* `WithMongoPipeline`: Set the MongoDB data processing pipeline to save crawled data to a MongoDB database.
* `WithSqlitePipeline`: Set the Sqlite data processing pipeline to save crawled data to a Sqlite database.
* `WithMysqlPipeline`: Set the MySQL data processing pipeline to save crawled data to a MySQL database.
* `WithKafkaPipeline`: Set the Kafka data processing pipeline to send crawled data to a Kafka message queue.
* `WithCustomPipeline`: Set the custom data processing pipeline.
* `WithRetryMaxTimes`: Set the maximum number of retries for requests.
* `WithRedirectMaxTimes` Set the maximum number of redirect for requests.
* `WithTimeout`: Set the timeout for requests.
* `WithInterval`: Set the interval between requests.
* `WithOkHttpCodes`: Set the normal HTTP status codes.

### Crawler Options

* `WithLogger`: Set the logger.
* `WithMockServerRoutes` Configure development service routes, including built-in or custom ones. You don't need to
  set `mock_server.enable: true` to enable the mock Server.
* `WithItemDelay` sets the data saving interval.
* `WithItemConcurrency` sets the data saving parallelism.
* `WithCDP` initial browser.
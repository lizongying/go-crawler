# go-crawler

A web crawling framework implemented in Golang, it is simple to write and delivers powerful performance. It comes with a
wide range of practical middleware and supports various parsing and storage methods. Additionally, it supports
distributed deployment.

[go-crawler](https://github.com/lizongying/go-crawler)

[document](https://pkg.go.dev/github.com/lizongying/go-crawler)

[中文](./README_CN.md)

## Feature

* Simple to write, yet powerful in performance.
* Built-in various practical middleware for easier development.
* Supports multiple parsing methods for simpler page parsing.
* Supports multiple storage methods for more flexible data storage.
* Provides numerous configuration options for richer customization.
* Allows customizations for components, providing more freedom for feature extensions.
* Includes a built-in mock Server for convenient debugging and development.
* It supports distributed deployment.

## Install

```shell
go get -u github.com/lizongying/go-crawler

# Currently, the framework is updated frequently. It's advisable to use the latest version, such as:
go get -u github.com/lizongying/go-crawler@latest
go get -u github.com/lizongying/go-crawler@0cfd21a

```

## Usage

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

* Basic Architecture
    * Spider: In the Spider, you can initiate requests and parse content. You need to set a unique name for each
      Spider.`spider.WithOptions(pkg.WithName("example"))`
    * BaseSpider: BaseSpider integrates components such as Spider, Downloader, Exporter, and Scheduler, serving as the
      processing center of the crawler.
    * Crawler：Within the Crawler, there can be multiple Spiders, and it manages the startup and shutdown of the Spiders.

* Spider Options
    * `WithName`: Set a unique name for the spider.
    * `WithHost`: Set the host for filtering based on the host or to support robots.txt.
    * `WithPlatforms`: Set the browser platforms.
    * `WithBrowsers`: Set the browsers.
    * `WithLogger`: Set the logger.
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
    * `WithPipeline`: Set the Pipeline to process the crawled data and perform subsequent operations.
    * `WithDumpPipeline`: Set the dump pipeline to print data to be saved.
    * `WithFilePipeline`: Set the file pipeline to handle crawled file data and save files to a specified location.
    * `WithImagePipeline`: Set the image pipeline to handle crawled image data and save images to a specified location.
    * `WithFilterPipeline`: Set the filter pipeline to filter crawled data.
    * `WithCsvPipeline`: Set the CSV data processing pipeline to save crawled data in CSV format.
    * `WithJsonLinesPipeline`: Set the JSON Lines data processing pipeline to save crawled data in JSON Lines format.
    * `WithMongoPipeline`: Set the MongoDB data processing pipeline to save crawled data to a MongoDB database.
    * `WithMysqlPipeline`: Set the MySQL data processing pipeline to save crawled data to a MySQL database.
    * `WithKafkaPipeline`: Set the Kafka data processing pipeline to send crawled data to a Kafka message queue.
    * `WithCustomPipeline`: Set the custom data processing pipeline.
    * `WithRetryMaxTimes`: Set the maximum number of retries for requests.
    * `WithTimeout`: Set the timeout for requests.
    * `WithInterval`: Set the interval between requests.
    * `WithOkHttpCodes`: Set the normal HTTP status codes.

* Crawler Options
    * `WithMockServerRoute` Configure development service routes, including built-in or custom ones. You need to
      set `mock_server.enable: true` to enable the mock Server.

* Item

  The Item is used to store data that needs to be exported and some other auxiliary information.
  The built-in Items in the framework cover major storage methods such as files, databases, and message queues.
  `pkg.Item` is an interface and cannot be used directly. `pkg.ItemUnimplemented` implements all methods of `pkg.Item`.
  If a custom Item needs to implement `pkg.Item`, it can be composed with `pkg.ItemUnimplemented`. For example:
    ```go
    type ItemNone struct {
        pkg.ItemUnimplemented
    }
    ```
    * Item has some common methods:
        * `Name() pkg.ItemName`: Get the specific type of the Item, such
          as `pkg.ItemNone`, `pkg.ItemCsv`, `pkg.ItemJsonl`, `pkg.ItemMongo`, `pkg.ItemMysql`, `pkg.ItemKafka`, etc.,
          which is used for deserializing the Item to the specific Item implementation.
        * `SetReferrer(string)`: Set the referrer, which can be used to record the source of the request. Generally,
          there is no need to set it manually as it is automatically set by the `ReferrerMiddleware`.
        * `Referrer() string`: Get the referrer.
        * `SetUniqueKey(string)`: Set the unique key, which can be used for filtering and other unique purposes.
        * `UniqueKey() string`: Get the unique key.
        * `SetId(any)`: Set the ID, mainly used as the primary key when saving data. One difference from `UniqueKey` is
          that `Id` may be generated in the Response and may not be obtained when making the request.
        * `Id() any`: Get the ID.
        * `SetData(any)`: Set the data, which is the complete data to be stored. For standardization, it is required to
          be a pointer type. When storing data in different destinations, the data needs to be set in different formats.
        * `Data() any`: Get the data.
        * `SetFilesRequest([]pkg.Request)`: Set the requests for downloading files. This is a slice and can be used to
          download multiple files.
        * `FilesRequest() []pkg.Request`: Get the requests for downloading files.
        * `SetFiles([]pkg.File)`: Set the downloaded files using this method.
        * `Files() []pkg.File`: Get the downloaded files.
        * `SetImagesRequest([]pkg.Request)`: Set the requests for downloading images. This is a slice and can be used to
          download multiple images.
        * `ImagesRequest() []pkg.Request`: Get the requests for downloading images.
        * `SetImages([]pkg.Image)`: Set the downloaded images using this method.
        * `Images() []pkg.Image`: Get the downloaded images.
    * Built-in Item Implementations: The framework provides some built-in Item implementations, such
      as `pkg.ItemNone`, `pkg.ItemCsv`, `pkg.ItemJsonl`, `pkg.ItemMongo`, `pkg.ItemMysql`, `pkg.ItemKafka`, etc.
      You can return an Item as needed and enable the corresponding Pipeline. For example:
      ```go
      err = s.YieldItem(ctx, items.NewItemMongo(s.collection, true).
      SetUniqueKey(extra.Keyword).
      SetId(extra.Keyword).
      SetData(&data))
      
      ```
      ```go
      spider.WithOptions(pkg.WithMongoPipeline())
      ```
        * pkg.ItemNone: This Item does not implement any other methods and is mainly used for debugging.
            * `items.NewItemNone()`
        * pkg.ItemCsv: Saves data to a CSV file.
            * `items.NewItemCsv(filename string)`: filename is the name of the file to be saved, without the extension.
        * pkg.ItemJsonl: Saves data to a JSONL file.
            * `items.NewItemJsonl(filename string)`: filename is the name of the file to be saved, without the
              extension.
        * pkg.ItemMongo: Saves data to MongoDB.
            * `items.NewItemMongo(collection string, update bool)`: collection is the MongoDB collection, and update
              indicates whether to update the data if it already exists in MongoDB.
        * pkg.ItemMysql: Saves data to MySQL.
            * `items.NewItemMysql(table string, update bool)`: table is the MySQL table, and update indicates whether to
              update the data if it already exists in MySQL.
        * pkg.ItemKafka: Sends data to Kafka.
            * `items.NewItemKafka(topic string)`: topic is the Kafka topic.
* Middleware and Pipeline include built-in ones, commonly used custom ones (internal/middlewares, internal/pipelines),
  and custom ones defined within the spider's module.
  Please make sure that the order values for different middleware and pipelines are not duplicated. If there are
  duplicate order values, the later middleware or pipeline will replace the earlier ones.

* Middleware

  In the framework, built-in middleware has pre-defined `order` values that are multiples of 10, such as 10, 20, 30, and
  so on. To avoid conflicts with the `order` values of built-in middleware, it is recommended to choose
  different `order` values when defining custom middleware.

  When customizing middleware, arrange them in the expected execution order based on their functionalities and
  requirements. Make sure that middleware with lower `order` values is executed first, followed by middleware with
  higher `order` values.

  Built-in middleware and custom middleware can use the default `order` values. If you need to change the
  default `order` value, include the `pkg.WithMiddleware(new(middleware), order)` option in the `NewApp` function to
  enable the middleware with the specified `order` value.

  The following are the built-in middleware with their respective `order` values:

    * custom: 10
        - Custom middleware.
        - To enable this middleware, include the `pkg.WithCustomMiddleware(new(CustomMiddleware))` option in
          the `NewApp` function.

    * dump: 20
        - Console dump middleware used for printing detailed information of item.data, including request and response
          details.
        - You can control whether to enable this middleware by configuring the `enable_dump_middleware` option, which is
          enabled by default.
        - To enable this middleware, include the `pkg.WithDumpMiddleware()` option in the `NewApp` function.

    * proxy: 30
        - Proxy switch middleware used for switching proxies for requests.
        - You can control whether to enable this middleware by configuring the `enable_proxy_middleware` option, which
          is enabled by default.
        - To enable this middleware, include the `pkg.WithProxyMiddleware()` option in the `NewApp` function.

    * robotsTxt: 40
        - Robots.txt support middleware for handling robots.txt files of websites.
        - You can control whether to enable this middleware by configuring the `enable_robots_txt_middleware` option,
          which is disabled by default.
        - To enable this middleware, include the `pkg.WithRobotsTxtMiddleware()` option in the `NewApp` function.

    * filter: 50
        - Request deduplication middleware used for filtering duplicate requests. By default, items are added to the
          deduplication queue only after being successfully saved.
        - You can control whether to enable this middleware by configuring the `enable_filter_middleware` option, which
          is enabled by default.
        - To enable this middleware, include the `pkg.WithFilterMiddleware()` option in the `NewApp` function.

    * file: 60
        - Automatic file information addition middleware used for automatically adding file information to requests.
        - You can control whether to enable this middleware by configuring the `enable_file_middleware` option, which is
          disabled by default.
        - To enable this middleware, include the `pkg.WithFileMiddleware()` option in the `NewApp` function.

    * image: 70
        - Automatic image information addition middleware used for automatically adding image information to requests.
        - You can control whether to enable this middleware by configuring the `enable_image_middleware` option, which
          is disabled by default.
        - To enable this middleware, include the `pkg.WithImageMiddleware()` option in the `NewApp` function.

    * retry: 80
        - Request retry middleware used for retrying requests when they fail.
        - The default maximum number of retries is 10. You can control whether to enable this middleware by configuring
          the `enable_retry_middleware` option, which is enabled by default.
        - To enable this middleware, include the `pkg.WithRetryMiddleware()` option in the `NewApp` function.

    * url: 90
        - URL length limiting middleware used for limiting the length of requests' URLs.
        - You can control whether to enable this middleware and set the maximum URL length by configuring
          the `enable_url_middleware` and `url_length_limit` options, respectively. Both options are enabled and set to
          2083 by default.
        - To enable this middleware, include the `pkg.WithUrlMiddleware()` option in the `NewApp` function.

    * referrer: 100
        - Automatic referrer addition middleware used for automatically adding the referrer to requests.
        - You can choose different referrer policies based on the `referrer_policy` configuration
          option. `DefaultReferrerPolicy` includes the request source, while `NoReferrerPolicy` does not include the
          request source.
        - You can control whether to enable this middleware by configuring the `enable_referrer_middleware` option,
          which is enabled by default.
        - To enable this middleware, include the `pkg.WithReferrerMiddleware()` option in the `NewApp` function.

    * cookie: 110
        - Automatic cookie addition middleware used for automatically adding cookies returned from previous requests to
          subsequent requests.
        - You can control whether to enable this middleware by configuring the `enable_cookie_middleware` option, which
          is enabled by default.
        - To enable this middleware, include the `pkg.WithCookieMiddleware()` option in the `NewApp` function.

    * redirect: 120
        - Website redirection middleware used for handling URL redirection. By default, it supports 301 and 302
          redirects.
        - You can control whether to enable this middleware and set the maximum number of redirections by
          configuring
          the `enable_redirect_middleware` and `redirect_max_times` options, respectively. Both options are enabled
          and set
          to 1 by default.
        - To enable this middleware, include the `pkg.WithRedirectMiddleware()` option in the `NewApp` function.

    * chrome: 130
        - Chrome simulation middleware used for simulating a Chrome browser.
        - You can control whether to enable this middleware by configuring the `enable_chrome_middleware` option,
          which is
          enabled by default.
        - To enable this middleware, include the `pkg.WithChromeMiddleware()` option in the `NewApp` function.

    * httpAuth: 140
        - HTTP authentication middleware used for performing HTTP authentication by providing a username and
          password.
        - You need to set the username and password in the specific request. You can control whether to enable this
          middleware by configuring the `enable_http_auth_middleware` option, which is disabled by default.
        - To enable this middleware, include the `pkg.WithHttpAuthMiddleware()` option in the `NewApp` function.

    * compress: 150
        - Gzip/deflate decompression middleware used for handling response compression encoding.
        - You can control whether to enable this middleware by configuring the `enable_compress_middleware` option,
          which is
          enabled by default.
        - To enable this middleware, include the `pkg.WithCompressMiddleware()` option in the `NewApp` function.

    * decode: 160
        - Chinese decoding middleware used for decoding responses with GBK, GB2312, and Big5 encodings.
        - You can control whether to enable this middleware by configuring the `enable_decode_middleware` option,
          which is
          enabled by default.
        - To enable this middleware, include the `pkg.WithDecodeMiddleware()` option in the `NewApp` function.

    * device: 170
        - Modify request device information middleware used for modifying the device information of requests,
          including
          request headers and TLS information. Currently, only User-Agent random switching is supported.
        - You need to set the device range (Platforms) and browser range (Browsers).
        - Platforms: Windows/Mac/Android/Iphone/Ipad/Linux
        - Browsers: Chrome/Edge/Safari/FireFox
        - You can control whether to enable this middleware by configuring the `enable_device_middleware` option,
          which is
          disabled by default.
        - To enable this middleware, include the `pkg.WithDeviceMiddleware()` option in the `NewApp` function.

    * http: 200
        - Create request middleware used for creating HTTP requests.
        - You can control whether to enable this middleware by configuring the `enable_http_middleware` option,
          which is
          enabled by default.
        - To enable this middleware, include the `pkg.WithHttpMiddleware()` option in the `NewApp` function.

    * stats: 210
        - Data statistics middleware used for collecting statistics on requests, responses, and processing in the
          spider.
        - You can control whether to enable this middleware by configuring the `enable_stats_middleware` option,
          which is
          enabled by default.
        - To enable this middleware, include the `pkg.WithStatsMiddleware()` option in the `NewApp` function.

* Pipeline

  Pipelines are used for stream processing of items, such as data filtering and data storage. By configuring different
  pipelines, you can conveniently process items and save the results to different targets, such as the console, files,
  databases, or message queues.

  Built-in pipelines and custom pipelines use the default `order` value. If you need to change the default `order`
  value, you can include the `pkg.WithPipeline(new(pipeline), order)` option in the `NewApp` function to enable the
  pipeline with the specified `order` value.

  The following are the built-in pipelines with their respective `order` values:

    * dump: 10
        - Used to print detailed information of items to the console.
        - You can control whether to enable this pipeline by configuring `enable_dump_pipeline`, which is enabled by
          default.
        - To enable this pipeline, include the `pkg.WithDumpPipeline()` option in the `NewApp` function.

    * file: 20
        - Used to download files and save them to items.
        - You can control whether to enable this pipeline by configuring `enable_file_pipeline`, which is enabled by
          default.
        - To enable this pipeline, include the `pkg.WithFilePipeline()` option in the `NewApp` function.

    * image: 30
        - Used to download images and save them to items.
        - You can control whether to enable this pipeline by configuring `enable_image_pipeline`, which is enabled by
          default.
        - To enable this pipeline, include the `pkg.WithImagePipeline()` option in the `NewApp` function.

    * filter: 200
        - Used for item filtering.
        - It can be used for deduplicating requests when filter middleware is enabled.
        - By default, items are only added to the deduplication queue after they are successfully saved.
        - You can control whether to enable this pipeline by configuring `enable_filter_pipeline`, which is enabled by
          default.
        - To enable this pipeline, include the `pkg.WithFilterPipeline()` option in the `NewApp` function.

    * csv: 101
        - Used to save results to CSV files.
        - You need to set the `FileName` in the `ItemCsv`, which specifies the name of the file to be saved (without the
          .csv extension).
        - You can use the tag `column:""` to define the column names of the CSV file.
        - You can control whether to enable this pipeline by configuring `enable_csv_pipeline`, which is disabled by
          default.
        - To enable this pipeline, include the `pkg.WithCsvPipeline()` option in the `NewApp` function.

    * jsonLines: 102
        - Used to save results to JSON Lines files.
        - You need to set the `FileName` in the `ItemJsonl`, which specifies the name of the file to be saved (without
          the .jsonl extension).
        - You can use the tag `json:""` to define the fields of the JSON Lines file.
        - You can control whether to enable this pipeline by configuring `enable_json_lines_pipeline`, which is disabled
          by default.
        - To enable this pipeline, include the `pkg.WithJsonLinesPipeline()` option in the `NewApp` function.

    * mongo: 103
        - Used to save results to MongoDB.
        - You need to set the `Collection` in the `ItemMongo`, which specifies the name of the collection to be saved.
        - You can use the tag `bson:""` to define the fields of the MongoDB document.
        - You can control whether to enable this pipeline by configuring `enable_mongo_pipeline`, which is disabled by
          default.
        - To enable this pipeline, include the `pkg.WithMongoPipeline()` option in the `NewApp` function.

    * mysql: 104
        - Used to save results to MySQL.
        - You need to set the `Table` in the `ItemMysql`, which specifies the name of the table to be saved.
        - You can use the tag `column:""` to define the column names of the MySQL table.
        - You can control whether to enable this pipeline by configuring `enable_mysql_pipeline`, which is disabled by
          default.
        - To enable this pipeline, include the `pkg.WithMysqlPipeline()` option in the `NewApp` function.

    * kafka: 105
        - Used to save results to Kafka.
        - You need to set the `Topic` in the `ItemKafka`, which specifies the name of the topic to be saved.
        - You can use the tag `json:""` to define the fields of the Kafka message.
        - You can control whether to enable this pipeline by configuring `enable_kafka_pipeline`, which is disabled by
          default.
        - To enable this pipeline, include the `pkg.WithKafkaPipeline()` option in the `NewApp` function.

    * custom: 110
        - Custom pipeline.
        - To enable this pipeline, include the `pkg.WithCustomPipeline(new(CustomPipeline))` option in the `NewApp`
          function.

* Request

    * `SetFingerprint(string) Request` Many websites nowadays implement security measures based on SSL fingerprints. By
      setting this parameter, you can perform disguising. If the fingerprint is `pkg.Browser`, the framework will
      automatically select a suitable fingerprint for this browser. If the fingerprint is in the ja3 format, the
      framework will apply this SSL fingerprint. If the fingerprint is empty, the framework will choose based on the
      user-agent. Note that the framework will only make modifications when `enable_ja3 = true`, and it uses the default
      SSL configuration of the Go programming language.
    * `SetClient(Client) Request` Some websites might implement security measures based on browser fingerprints, or the
      requests might require complex simulation. In such cases, it's recommended to use a headless browser. By setting
      the Client to `pkg.Browser`, the framework will automatically enable a headless browser for processing.
    * `SetAjax(bool) Request` If you need to use a headless browser and the request is an AJAX request, please set
      this option to true. The framework will handle the request as an XHR (XMLHttpRequest) request. You may also
      need to set the referrer.

* Response

  The `Response` wraps around `http.Response` and provides the following functionalities:
    - `Xpath() (*xpath.Selector, error)` `MustXpath() *xpath.Selector`
      Returns an XPath selector, for specific syntax, please refer
      to [go-xpath](https://github.com/lizongying/go-xpath).
    - `Css() (*css.Selector, error)` `MustCss() *css.Selector`
      Returns a CSS selector, for specific syntax, please refer to [go-css](https://github.com/lizongying/go-css).
    - `Json() (*gson.Selector, error)` `MustJson() gjson.Result`
      Returns a gjson selector, for specific syntax, please refer to [go-json](https://github.com/lizongying/go-json).
    - `Re() (*re.Selector, error)` `MustRe() *re.Selector`
      Returns a regular expression selector, for specific syntax, please refer
      to [go-re](https://github.com/lizongying/go-re).
    - `AllLink() []*url.URL`
      Retrieves all links from the response.
    - `BodyText() string`
      Retrieves the cleaned text content without HTML tags, the handling may be rough.
    - `AbsoluteURL(relativeUrl string) (absoluteURL *url.URL, err error)`
      Retrieves the absolute URL for a given relative URL.

* Signals

  By using signals, it's possible to capture crawler events and perform certain actions. You need to register
  the `func(pkg.Spider)` method for the following signals:

    * SpiderOpened: Indicates the start of the spider. You can register it using `RegisterSpiderOpened`.
    * SpiderClosed: Indicates the end of the spider. You can register it using `RegisterSpiderClosed`.

* Parsing

  The framework comes with several built-in parsing modules. You can choose the one that suits your specific spider's
  needs to extract the required data from web pages.

    * Css Selector: [go-css](https://github.com/lizongying/go-css) is a library for handling css selectors. You
      can use the css selector syntax to
      extract data from HTML by calling the `response.Css()` method.
    * XPath Selector: [go-xpath](https://github.com/lizongying/go-xpath) is a library for XPath selection. You can use
      XPath expressions to extract data from HTML
      by calling the `response.Xpath()` method.
    * GJSON Selector: [go-json](https://github.com/lizongying/go-json) is a library for handling JSON. You can use the
      gjson syntax to extract data from JSON by calling
      the `response.Json()` method.
    * Regular Expression Selector: [go-re](https://github.com/lizongying/go-re) is a library for handling regular
      expressions. You can use regular expressions
      to extract data from the response by calling the `response.Re()` method.

* Proxy

  You can set up tunnel proxies by using tools like [go-proxy](https://github.com/lizongying/go-proxy) to provide random
  proxy switching functionality, transparent to the caller. You can integrate these proxy tools into your spider
  framework to automatically switch proxies when making requests. The random switching tunnel proxy provides convenience
  and ease of use to the caller. In the future, other calling methods may be added, such as maintaining the original
  proxy address, to provide greater flexibility to meet different proxy requirements.

    * Proxy Configuration in Spider

      Currently, only random switching of proxies is supported in the spider configuration.

* File and Image Downloads
    * If you want to save files to object storage like S3, you need to perform the corresponding configuration.
    * File Download
        * Set Files Requests in Item: In the Item, you need to set Files requests, which include a list of requests for
          downloading files. You can use the `item.SetFilesRequest([]pkg.Request{...})` method to set the list of
          requests.
        * Item.data: Your Item.data field needs to implement a slice of `pkg.File` to store the downloaded file results.
          The name of this field must be "Files," for example: `type DataFile struct { Files []*media.File }`.
    * Image Download
        * Set Images Requests in Item: In the Item, you need to set Images requests, which include a list of requests
          for downloading images. You can use the `item.SetImagesRequest([]pkg.Request{...})` method to set the list of
          requests.
        * Item.data: Your Item.data field needs to implement a slice of `pkg.Image` to store the downloaded image
          results. The name of this field must be "Images," for
          example: `type DataImage struct { Images []*media.Image }`.

* Spider Structure

  It is recommended to have one spider for each website (or sub-website) or each specific business. You don't need to
  split it too finely, nor do you need to include all websites and businesses in one spider. You can package each spider
  separately or combine multiple spiders together to reduce the number of files. However, during execution, only one
  spider can be started.

  ```go
  app.NewApp(NewExample1Spider, NewExample2Spider).Run()
  ```

* Local MockServer

  To facilitate development and debugging, the framework comes with a built-in local MockServer that can be enabled by
  setting `mock_server.enable: true` in the configuration. By using the local MockServer, you can more easily simulate
  and
  observe network requests and responses, as well as handle custom route logic. This provides developers with a
  convenient tool to quickly locate and resolve issues.

  You can customize routes by implementing the `pkg.Route` interface and registering them with the MockServer in the
  spider by calling `AddMockServerRoutes(...pkg.Route)`.

    * The MockServer supports both HTTP and HTTPS, and you can specify the MockServer's URL by setting the `mock_server`
      option. `http://localhost:8081` represents using HTTP protocol, and `https://localhost:8081` represents using
      HTTPS protocol.
    * By default, the MockServer displays JA3 fingerprints. JA3 is an algorithm used for TLS client fingerprinting, and
      it shows information about the TLS version and cipher suites used by the client when establishing a connection
      with the server.
    * You can use the tls tool to generate the server's private key and certificate for use with HTTPS in the
      MockServer.
      The tls tool can help you generate self-signed certificates for local development and testing environments.
    * The MockServer includes multiple built-in routes that provide rich functionalities to simulate various network
      scenarios and assist in development and debugging. You can choose the appropriate route based on your needs and
      configure it in the MockServer to simulate specific network responses and behaviors.

        * BadGatewayRoute: Simulates returning a 502 status code.
        * Big5Route: Simulates using the big5 encoding.
        * CookieRoute: Simulates returning cookies.
        * DeflateRoute: Simulates using Deflate compression.
        * FileRoute: Simulates outputting files.
        * Gb2312Route: Simulates using the gb2312 encoding.
        * Gb18030Route: Simulates using the gb18030 encoding.
        * GbkRoute: Simulates using the gbk encoding.
        * GzipRoute: Simulates using gzip compression.
        * HelloRoute: Prints the header and body information of the request.
        * HtmlRoute: simulates the return of HTML static files. You can place HTML files inside the `/static/html/`
          directory for web parsing testing purposes, eliminating the need for redundant requests.
        * HttpAuthRoute: Simulates http-auth authentication.
        * InternalServerErrorRoute: Simulates returning a 500 status code.
        * OkRoute: Simulates normal output, returning a 200 status code.
        * RateLimiterRoute: Simulates rate limiting, currently based on all requests and not differentiated by users.
          Can be used in conjunction with HttpAuthRoute.
        * RedirectRoute: Simulates a 302 temporary redirect, requires enabling OkRoute simultaneously.
        * RobotsTxtRoute: Returns the robots.txt file.

### Startup

By configuring environment variables or parameters, you can start the crawler more flexibly, including selecting the
configuration file, specifying the spider's name, defining the initial method, passing additional parameters, and
setting the startup mode.

```shell
spider -c example.yml -n example -f TestOk -m once
```

* Configuration file path, must be configured. It is recommended to use different configuration files for different
  environments.
    * Environment variable `CRAWLER_CONFIG_FILE`
    * Startup parameter `-c`
* Spider name, must be configured.
    * Environment variable `CRAWLER_SPIDER_NAME`
    * Startup parameter `-n`
* Initial method, default is "Test". Please note that the case must be consistent.
    * Environment variable `CRAWLER_SPIDER_FUNC`
    * Startup parameter `-f`
* Additional parameters, this parameter is optional. It is recommended to use a JSON string. The parameters will be
  passed to the initial method.
    * Environment variable `CRAWLER_SPIDER_ARGS`
    * Startup parameter `-a`
* Startup mode, default is "manual". You can use different modes as needed
    * Environment variable `CRAWLER_SPIDER_MODE`
    * Startup parameter `-m`
    * You can use different modes as needed:
    * Optional values
        * `manual`: Executes manually (default is no execution); can be managed through the API.
        * `loop`: Executes repeatedly.
        * `once`: Executes only once.
        * `cron`: Executes at scheduled intervals.
* Scheduled task. This configuration is only applied when the mode is set to "cron", such as "1s/2i/3h/4d/5m/6w"
    * Environment variable `CRAWLER_SPIDER_SPEC`
    * Startup parameter `-s`

### Configuration

In the configuration file, you can set global configurations that apply to all spiders. However, some configurations can
be modified and overridden in individual spiders or specific requests.
The configuration file needs to be specified at startup using environment variables or parameters. Here are the
configuration parameters:

* `env: dev`. In the `dev` environment, data will not be written to the database.
* `bot_name: crawler` Project Name

Database Configuration:

* `mongo_enable:` Whether to enable MongoDB.
* `mongo.example.uri:` MongoDB URI.
* `mongo.example.database:` MongoDB database name.
* `mysql_enable:` Whether to enable MySQL.
* `mysql.example.uri:` MySQL URI.
* `mysql.example.database:` MySQL database name.
* `redis_enable:` Whether to enable Redis.
* `redis.example.addr:` Redis address.
* `redis.example.password:` Redis password.
* `redis.example.db:` Redis database number.
* `s3_enable:` Whether to enable S3 object storage (e.g., COS, OSS, MinIO, etc.).
* `s3.example.endpoint:` S3 endpoint.
* `s3.example.region:` S3 region.
* `s3.example.id:` S3 access ID.
* `s3.example.key:` S3 access key.
* `s3.example.bucket:` S3 bucket name.
* `kafka_enable:` Whether to enable Kafka.
* `kafka.example.uri:` Kafka URI.

Log Configuration:

* `log.filename:` Log file path. You can use "{name}" to replace it with a parameter from `-ldflags`.
* `log.long_file:` If set to true (default), it logs the full file path.
* `log.level:` Log level, options are DEBUG/INFO/WARN/ERROR.

* `mock_server`: Mock Server
    * `enable: false`: Whether to enable the mock Server.
    * `host: https://localhost:8081`: The address of the mock Server.
    * `client_auth: 0` Client authentication type, 0 means no authentication.

Middleware and Pipeline Configuration:

* `enable_stats_middleware:` Whether to enable the statistics middleware, enabled by default.
* `enable_dump_middleware:` Whether to enable the dump middleware for printing requests and responses, enabled by
  default.
* `enable_filter_middleware:` Whether to enable the filter middleware, enabled by default.
* `enable_file_middleware:` Whether to enable the file handling middleware, enabled by default.
* `enable_image_middleware:` Whether to enable the image handling middleware, enabled by default.
* `enable_http_middleware:` Whether to enable the HTTP request middleware, enabled by default.
* `enable_retry_middleware:` Whether to enable the request retry middleware, enabled by default.
* `enable_referrer_middleware:` Whether to enable the Referrer middleware, enabled by default.
* `referrer_policy:` Set the Referrer policy, options are DefaultReferrerPolicy (default) and NoReferrerPolicy.
* `enable_http_auth_middleware:` Whether to enable the HTTP authentication middleware, disabled by default.
* `enable_cookie_middleware:` Whether to enable the Cookie middleware, enabled by default.
* `enable_url_middleware:` Whether to enable the URL length limiting middleware, enabled by default.
* `url_length_limit:` Maximum length limit for URLs, default is 2083.
* `enable_compress_middleware:` Whether to enable the response decompression middleware (gzip, deflate), enabled by
  default.
* `enable_decode_middleware:` Whether to enable the Chinese decoding middleware (GBK, GB2312, Big5 encodings), enabled
  by default.
* `enable_redirect_middleware:` Whether to enable the redirect middleware, enabled by default.
* `redirect_max_times:` Maximum number of times to follow redirects, default is 10.
* `enable_chrome_middleware:` Whether to enable the Chrome simulation middleware, enabled by default.
* `enable_device_middleware:` Whether to enable the device simulation middleware, disabled by default.
* `enable_proxy_middleware:` Whether to enable the proxy middleware, enabled by default.
* `enable_robots_txt_middleware:` Whether to enable the robots.txt support middleware, disabled by default.
* `enable_dump_pipeline:` Whether to enable the print item pipeline, enabled by default.
* `enable_file_pipeline:` Whether to enable the file download pipeline, enabled by default.
* `enable_image_pipeline:` Whether to enable the image download pipeline, enabled by default.
* `enable_filter_pipeline:` Whether to enable the filter pipeline, enabled by default.
* `enable_csv_pipeline:` Whether to enable the CSV pipeline, disabled by default.
* `enable_json_lines_pipeline:` Whether to enable the JSON Lines pipeline, disabled by default.
* `enable_mongo_pipeline:` Whether to enable the MongoDB pipeline, disabled by default.
* `enable_mysql_pipeline:` Whether to enable the MySQL pipeline, disabled by default.
* `enable_kafka_pipeline:` Whether to enable the Kafka pipeline, disabled by default.
* `enable_priority_queue:` Whether to enable the priority queue, enabled by default, currently only supports Redis.

Other Configurations:

* `proxy.example`: Proxy configuration.
* `request.concurrency`: Number of concurrent requests.
* `request.interval`: Request interval time in milliseconds. Default is 1000 milliseconds (1 second).
* `request.timeout`: Request timeout in seconds. Default is 60 seconds (1 minute).
* `request.ok_http_codes`: Normal HTTP status codes for requests.
* `request.retry_max_times`: Maximum number of retries for requests. Default is 10.
* `request.http_proto`: HTTP protocol for requests. Default is `2.0`.
* `enable_ja3`: Whether to modify/print JA3 fingerprints. Default is disabled.
* `scheduler`: Scheduler method. Default is `memory` (memory-based scheduling). Options are `memory`, `redis`, `kafka`.
  Selecting `redis` or `kafka` enables cluster scheduling.
* `filter`: Filter method. Default is `memory` (memory-based filtering). Options are `memory`, `redis`.
  Selecting `redis` enables cluster filtering.

## Web Page Parsing Based on Field Tags

In this framework, the returned data is a struct. We only need to add parsing rule tags to the fields, and the framework
will automatically perform web page parsing, making it appear very clean and concise.

For some simple web scraping tasks, this approach is more convenient and efficient. Especially when you need to create a
large number of generic web scrapers, you can directly configure these tags for parsing.

For example:

```go
type DataRanks struct {
Data []struct {
Name           string  `_json:"name"`
FullName       string  `_json:"fullname"`
Code           string  `_json:"code"`
MarketBalue    int     `_json:"market_value"`
MarketValueUsd int     `_json:"market_value_usd"`
Marketcap      int     `_json:"marketcap"`
Turnoverrate   float32 `_json:"turnoverrate"`
} `_json:"data"`
}
```

You can set the root parsing for `data` as `_json:"data"`, meaning that the fields inside it are all parsed under the
root. For example, `_json:"name"`.

You can mix and match root and sub-tags, for instance, use XPath for the root and JSON for the sub-tags.

You can use the following tags:

* `_json:""` for gjson format
* `_xpath:""` for XPath format
* `_css:""` for CSS format
* `_re:""` for regular expression (regex) format

## Some Question

* In some frameworks, there is a presence of `start_urls`. How is it set up in this framework?

  In this framework, this approach has been removed. It's possible to explicitly create requests within the initial
  method and perform additional processing on those requests, which can actually be more convenient.
    ```go
    startUrls := []string{"/a.html", "/b.html"}
    for _, v:=range startUrls {
		if err = s.YieldRequest(ctx, request.NewRequest().
            SetUrl(fmt.Sprintf("https://a.com%s", v)).
            SetCallBack(s.Parse)); err != nil {
            s.logger.Error(err)
        }
    }

    ```
* What are the ways to improve spider performance?

  To improve the performance of the spider, you can consider disabling some unused middleware or pipelines to reduce
  unnecessary processing and resource consumption. Before disabling any middleware or pipeline, please assess its actual
  impact on the spider's performance. Ensure that disabling any part will not have a negative impact on the
  functionality.

* Why isn't item implemented as a distributed queue?

  The crawler processes its own items, and there is no need to handle items from other crawlers.
  Therefore, while the framework has reserved the architecture for distributed queues, it does not use external queues
  to replace the in-memory queue used by the program.
  If there are performance issues with processing, it is recommended to output the results to a queue.

* How to Set the Request Priority?

  Priorities are allowed to range from 0 to 2147483647.
  Priority 0 is the highest and will be processed first.
  Currently, only Redis-based priority queues are supported.

    ```go
    request.SetPriority(0)
    ```

* When will the crawler end?

  The crawler will end and the program will close when the following conditions are met under normal circumstances:

    1. All requests and parsing methods have been executed.
    2. The item queue is empty.
    3. The request queue is empty.

  When these conditions are fulfilled, the crawler has completed its tasks and will terminate.

* How to prevent the spider from stopping?

  Simply return `pkg.DontStopErr` in the `Stop` method.

    ```go
    package main
    
    import "github.com/lizongying/go-crawler/pkg"
    
    func (s *Spider) Stop(ctx context.Context) (err error) {
        if err = s.Spider.Stop(ctx); err != nil {
            s.logger.Error(err)
            return
        }
    
        err = pkg.DontStopErr
        s.logger.Error(err)
        return
    }

    ```

* Which should be used in the task queue: `request`, `extra`, or `unique_key`?

  Firstly, it should be noted that these three terms are concepts within this framework:
    * `request` contains all the fields of a request, including URL, method, headers, and may have undergone middleware
      processing. The drawback is that it occupies more space, making it somewhat wasteful as a queue value.
    * `extra` is a structured field within the request and, in the framework's design, it contains information that can
      construct a unique request (in most cases). For instance, a list page under a category may include the category ID
      and page number. Similarly, a detail page may include a detail ID. To ensure compatibility with various languages,
      the storage format in the queue is JSON, which is more space-efficient. It's recommended to use this option.
    * `unique_key` is a unique identifier for a request within the framework and is a string. While it can represent
      uniqueness in some cases, it can become cumbersome when requiring a combination of multiple fields to be unique –
      such as in the case of list pages or detail pages involving both a category and an ID. If memory is constrained (
      e.g., in Redis usage), it can be used. However, for greater generality, using `extra` might be more convenient.

  Enqueuing:
    * `YieldExtra` or `MustYieldExtra`

  Dequeuing:
    * `GetExtra` or `MustGetExtra`

* Whether to use `Must[method]`, such as `MustYieldRequest`?

  `Must[method]` is more concise, but it might be less convenient for troubleshooting errors. Whether to use it depends
  on the individual style of the user.
  If there's a need for specific error handling, then regular methods like `YieldRequest` should be used.

## Complete example

exampleSpider.go

```go
package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mockServer"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
)

type ExtraOk struct {
	Count int
}

type DataOk struct {
	Count int
}

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.UnmarshalExtra(&extra); err != nil {
		s.logger.Error(err)
		return
	}

	if err = s.YieldItem(ctx, items.NewItemNone().
		SetData(&DataOk{
			Count: extra.Count,
		})); err != nil {
		s.logger.Error(err)
		return
	}

	if extra.Count > 0 {
		s.logger.Info("manual stop")
		return
	}

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseOk)); err != nil {
		s.logger.Error(err)
	}
	return
}

func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk)); err != nil {
		s.logger.Error(err)
	}
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("example"),
		pkg.WithHost("https://localhost:8081"),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run(pkg.WithMockServerRoute(mockServer.NewRouteOk))
}

```

or

```go
package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/mockServer"
	"github.com/lizongying/go-crawler/pkg/request"
)

type ExtraOk struct {
	Count int
}

type DataOk struct {
	Count int
}

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	response.MustUnmarshalExtra(&extra)

	s.MustYieldItem(ctx, items.NewItemNone().
		SetData(&DataOk{
			Count: extra.Count,
		}))

	if extra.Count > 0 {
		s.logger.Info("manual stop")
		return
	}

	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseOk))
	return
}

func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("example"),
		pkg.WithHost("https://localhost:8081"),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run(pkg.WithMockServerRoute(mockServer.NewRouteOk))
}

```

```shell
go run exampleSpider.go -c example.yml -n example -f TestOk -m once

```

For more examples, you can refer to the following project.

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

```shell
git clone github.com/lizongying/go-crawler-example
```

## TODO

* AutoThrottle
* monitor
* statistics

```shell
go get -u github.com/lizongying/go-css@latest
go get -u github.com/lizongying/go-xpath@latest
go get -u github.com/lizongying/go-re@latest
go get -u github.com/lizongying/go-json@latest

```

## Docker build

```shell
docker build -f ./cmd/testSpider/Dockerfile -t go-crawler/test-spider:latest . 
```

```shell
docker run -d go-crawler/test-spider:latest spider -c example.yml -f TestRedirect -m once
```

## Feature Support Summary

* Parsing supports CSS, XPath, Regex, and JSON.
* Output supports JSON, CSV, MongoDB, MySQL, and Kafka.
* Supports Chinese decoding for gb2312, gb18030, gbk, big5 character encodings.
* Supports distributed processing.
* Supports Redis and Kafka as message queues.
* Supports automatic handling of cookies and redirects.
* Browser simulation is supported.
* Mock server is supported.
* Priority queue is supported.
* Supports scheduled tasks, recurring tasks, and one-time tasks.


1. Parsing:
    - CSS: Supported
    - XPath: Supported
    - Regex: Supported
    - JSON: Supported

2. Output:
    - JSON: Supported
    - CSV: Supported
    - MongoDB: Supported
    - MySQL: Supported
    - Kafka: Supported

3. Chinese Decoding:
    - gb2312: Supported
    - gb18030: Supported
    - gbk: Supported
    - big5: Supported

4. Distributed Processing: Supported

5. Message Queues:
    - Redis: Supported
    - Kafka: Supported

6. Automatic Handling:
    - Cookies: Supported
    - Redirects: Supported

7. Mock Server:
8. DNS Cache

## Generate Certificate

```shell
# CA-Signed Certificate
./releases/tls

# Self-Signed Certificate
./releases/tls -s

```

## MITM

```shell
./releases/mitm
```
# go-crawler

真正的爬虫框架，编写简单，性能强劲。丰富的内置中间件。支持多种解析、保存方式。灵感来源于scrapy。

[go-crawler](https://github.com/lizongying/go-crawler)
[document](https://pkg.go.dev/github.com/lizongying/go-crawler)

## Feature

* 编写简单，性能强劲。
* 内置devServer，方便调试开发。
* 丰富的配置项，自由性更高。
* 丰富的内置中间件，减少额外工作。
* 自定义中间件简单方便。
* 支持多种解析方式，解析页面更简单。
* 支持多种保存方式，按需使用。

## Usage

* 为了方便开发调试，增加了本地devServer，在`-m dev`模式下会默认启用。可以自定义route，仅需要实现`pkg.Route`
  ，然后在spider中通过`AddDevServerRoutes(...pkg.Route)`注册到devServer即可。
    * 支持http和https，可以设置`dev_server`。如http`http://localhost:8081`，https`https://localhost:8081`。
    * 默认显示ja3指纹。
    * 可以通过tls工具`tls`生成服务器的私钥和证书。
* 基本架构
    * spider-baseSpider-crawler
    * spider主要是发起请求和回调解析方法。需要给spider起个名字，`spider.SetName(name)`
    * baseSpider，实现spider的公共方法，不必在spider重复编写。如GetName和setName
    * crawler集成spider、downloader、exporter、scheduler等，是爬虫的处理逻辑中心
    * 由于方法继承，实际上spider可以直接调用baseSpider和crawler里的方法。
* crawler选项
    * WithMode 设置Mode，会执行SetMode
    * WithPlatforms 设置Platforms，会执行SetPlatforms
    * WithBrowsers 设置Browsers，会执行SetBrowsers
    * WithLogger 设置Logger，会执行SetLogger
    * WithFilter 设置Filter，会执行SetFilter
    * WithDownloader 设置Downloader，会执行SetDownloader
    * WithExporter 设置Exporter，会执行SetExporter
    * WithMiddleware 设置Middleware，会执行SetMiddleware
    * WithPipeline 设置Pipeline，会执行SetPipeline
    * WithRetryMaxTimes 设置请求最大重试此时，会执行SetRetryMaxTimes
    * WithTimeout 设置请求超时时间，会执行SetTimeout
    * WithInterval 设置请求间隔，会执行SetInterval
    * WithOkHttpCodes 设置正常的HttpCodes，会执行SetOkHttpCodes
* 需要返回的item需要实现Item（可以组合ItemUnimplemented）
    * `GetReferer()` 可以获取到referer。
    * UniqueKey 唯一键，不会保存到数据库。可以用作过滤等其他用途。
    * Id 保存主键
    * Data 完整数据
    * 内置item：ItemCsv、ItemJsonl、ItemMongo、ItemMysql、ItemKafka，需要开启相应pipeline，进行保存
* middleware包括框架内置、公共自定义（internal/middlewares，internal/pipelines）和爬虫内自定义（和爬虫同module）。
* middleware/pipeline的order不能重复。相同order，后面的middleware/pipeline会替换前面的middleware/pipeline
* 框架内置middleware，自定义middleware请参照以下order进行配置。内置中间件order为10的倍数，自定义中间件请避开。
    * stats:10
        * 数据统计
        * 配置 enable_stats_middleware: true 是否开启统计，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.StatsMiddleware), 10)`
    * dump:20
        * 控制台打印item.data
        * 配置 enable_dump_middleware: true 是否开启打印request/response，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.DumpMiddleware), 20)`
    * filter:30
        * 过滤重复请求。默认支持的是item保存成功后才会进入去重队列，防止出现请求失败后再次请求却被过滤的问题。
          当请求速度大于保存速度的时候可能会有请求不被过滤的情况。
        * 配置 enable_filter_middleware: true 是否开启过滤，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.FilterMiddleware), 30)`
    * image:40
        * 自动添加图片的宽高等信息
        * 配置 enable_image_middleware: false 是否开启图片处理，默认未开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.ImageMiddleware), 40)`
    * http:50
        * 创建request
        * 配置 enable_http_middleware: true 是否开启创建http request，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.HttpMiddleware), 50)`
    * retry:60
        * 如果请求出错，会进行重试。
        * `RetryMaxTimes=10`
        * 配置 enable_retry_middleware: true 是否开启重试，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.RetryMiddleware), 60)`
    * url:70
        * 限制url的长度
        * 配置 enable_url_middleware: true 是否开启url长度限制，默认开启
        * 配置 url_length_limit: 2083 url的最长长度默认为2083
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.UrlMiddleware), 70)`
    * referer:80
        * 通过设置`referrer_policy`采用不同的referer策略，
        * DefaultReferrerPolicy。默认会加入请求来源
        * NoReferrerPolicy。不加入请求来源
        * 配置 enable_referer_middleware: true 是否开启自动添加referer，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.RefererMiddleware), 80)`
    * cookie:90
        * 如果之前请求返回cookie，会自动加到后面的请求里
        * 配置 enable_cookie_middleware: true 是否开启cookie支持，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.CookieMiddleware), 90)`
    * redirect:100
        * 网址重定向，默认支持301、302
        * 配置 enable_redirect_middleware: true 是否开启重定向，默认开启
        * 配置 redirect_max_times: 1 重定向最大次数
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.RedirectMiddleware), 100)`
    * chrome:110
        * 模拟chrome
        * 配置 enable_chrome_middleware: true 模拟chrome，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.ChromeMiddleware), 110)`
    * httpAuth:120
        * 通过`username`、`password`添加httpAuth认证。需要设置`SetUsername`和`SetPassword`
        * 配置 enable_http_auth_middleware: true 是否开启httpAuth，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.HttpAuthMiddleware), 120)`
    * compress:130
        * 支持 gzip/deflate解压缩
        * 配置 enable_compress_middleware: true 是否开启gzip/deflate解压缩，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.CompressMiddleware), 130)`
    * decode:140
        * 支持gbk、gb2310、big5中文解码
        * 配置 enable_decode_middleware: true 是否开启中文解码，默认开启
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.DecodeMiddleware), 140)`
    * device:150
        * 修改request设备信息。修改header和tls信息，暂时只支持user-agent随机切换。需要设置`SetPlatforms`和`SetBrowsers`
          限定设备范围。默认不启用。
        * Platforms: Windows/Mac/Android/Iphone/Ipad/Linux
        * Browsers: Chrome/Edge/Safari/FireFox
        * 配置 enable_device_middleware: false 随机模拟设备，默认关闭
        * 启用方法：在NewApp中加入crawler选项`WithMiddleware(new(middlewares.DeviceMiddleware), 150)`
* pipeline用于处理item。
    * dump:10
        * 控制台打印item详细
        * 配置 enable_dump_pipeline: true 是否开启打印item详细，默认启用
    * filter:100
        * filter可能有不同的实现
        * 用户过滤重复请求，需要middleware同时开启filter。默认支持的是item保存成功后才会进入去重队列。
          当请求速度大于保存速度的时候可能会有请求不被过滤的情况，如果不需要判断item，可以在middleware中进行进入去重队列操作
        * 配置 enable_filter_pipeline: true 是否开启过滤，默认启用
    * csv
        * 保存结果到csv文件。
        * 需在在ItemCsv中设置`FileName`，保存的文件名称，不包含.csv
        * 启用方法：在NewApp中加入crawler选项`WithPipeline(pipelines.NewCsvPipeline, 11)`
    * jsonLines
        * 保存结果到jsonlines文件。
        * 需在在ItemJsonl中设置`FileName`，保存的文件名称，不包含.jsonl
        * 启用方法：在NewApp中加入crawler选项`WithPipeline(pipelines.NewJsonLinesPipeline, 12)`
    * mongo
        * 保存结果到mongo。
        * 需在在ItemMongo中设置`Collection`，保存的collection
        * 启用方法：在NewApp中加入crawler选项`WithPipeline(pipelines.NewMongoPipeline, 13)`
    * mysql
        * 保存结果到mysql。
        * 需在在ItemMysql中设置`table`，保存的table
        * 启用方法：在NewApp中加入crawler选项`WithPipeline(pipelines.NewMysqlPipeline, 14)`
    * kafka
        * 保存结果到kafka。
        * 需在在ItemKafka中设置`Topic`，保存的topic
        * 启用方法：在NewApp中加入crawler选项`WithPipeline(pipelines.NewKafkaPipeline, 15)`
* 在配置文件中可以配置全局request参数，在具体request中可以覆盖此配置
* 解析模块
    * query选择器 [go-query](https://github.com/lizongying/go-query)
        * ```response.Query()```
    * xpath选择器 [go-xpath](https://github.com/lizongying/go-xpath)
        * ```response.Xpath()```
    * gjson
        * ```response.Json()```
    * re选择器 [go-re](https://github.com/lizongying/go-re)
        * ```response.Re()```
* 代理
    * 可以自行搭建隧道代理 [go-proxy](https://github.com/lizongying/go-proxy)
      。这是一个随机切换的隧道代理，调用方无感知，方便使用。后期会加入一些其他的调用方式，比如维持原来的代理地址。
* 增加爬虫性能
    * 在不影响功能的情况下，可以考虑关闭一些用不到的中间件或pipeline。可以在配置文件中修改，或者爬虫入口中修改
    * 配置文件:
        * enable_stats_middleware: false
        * enable_dump_middleware: false
        * enable_filter_middleware: false
        * enable_image_middleware: false
        * enable_http_middleware: false
        * enable_retry_middleware: false
        * enable_referer_middleware: false
        * enable_http_auth_middleware: false
        * enable_cookie_middleware: false
        * enable_url_middleware: false
        * enable_compress_middleware: false
        * enable_decode_middleware: false
        * enable_redirect_middleware: false
        * enable_chrome_middleware: false
        * enable_device_middleware: false
        * enable_dump_pipeline: false
        * enable_filter_pipeline: false
* 爬虫结构
    * 建议按照每个网站（子网站）或者每个业务为一个spider。不必分的太细，也不必把所有的网站和业务都写在一个spider里

### args

* -c config file. must set it. 配置文件，必须配置。
* -f start func. default `Test`. 入口方法，默认`Test`。
* -a args. json string. 额外的参数，用于入口方法调用，非必须项。
* -m mode. default `test`. 启动模式，如`dev`,`prod`等，默认`test`

### config

* mongo.example.uri: mongo uri
* mongo.example.database: mongo database
* log.filename: Log file path. You can replace {name} with -ldflags.
* log.long_file: If set to true, the full file path is logged.
* log.level: DEBUG/INFO/WARN/ERROR
* proxy.example: proxy
* request.concurrency: Number of request concurrency
* request.interval: Request interval(Millisecond). If set to 0, it is the default interval(1000). If set to a negative
  number,
  it is 0.
* request.timeout: Request timeout(seconds)
* request.ok_http_codes: Request ok httpcodes
* request.retry_max_times: Request retry max times，默认10
* request.http_proto: Request http proto
* dev_server: devServer。如http`http://localhost:8081`，https`https://localhost:8081`。
* enable_ja3: false devServer是否显示ja3指纹，默认关闭
* enable_stats_middleware: true 是否开启统计，默认开启
* enable_dump_middleware: true 是否开启打印request/response middleware，默认开启
* enable_filter_middleware: true 是否开启过滤middleware，默认开启
* enable_image_middleware: true 是否开启image，默认开启
* enable_http_middleware: true 是否开启http，默认开启
* enable_retry_middleware: true 是否开启重试，默认开启
* enable_referer_middleware: true 是否开启referer，默认开启
* referrer_policy: DefaultReferrerPolicy 来源政策，默认DefaultReferrerPolicy，可选DefaultReferrerPolicy、NoReferrerPolicy
* enable_http_auth_middleware: true 是否开启httpAuth，默认开启
* enable_cookie_middleware: true 是否开启cookie，默认开启
* enable_url_middleware: true 是否开启url长度限制，默认开启
* url_length_limit: 2083 url长度限制，默认2083
* enable_compress_middleware: true 是否开启gzip/deflate解压缩，默认开启
* enable_decode_middleware: true 是否开启中文解码，默认开启
* enable_redirect_middleware: true 是否开启重定向，默认开启
* redirect_max_times: 1 重定向最大次数，默认1
* enable_chrome_middleware: true 模拟chrome，默认开启
* enable_device_middleware: false 随机模拟设备，默认关闭
* enable_dump_pipeline: true 是否开启打印item pipeline，默认开启
* enable_filter_pipeline: true 是否开启过滤pipeline，默认开启

## Example

可以按照以下示例进行开发

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

```shell
git clone github.com/lizongying/go-crawler-example
```

## Test

```shell
go run cmd/testSpider/*.go -c dev.yml -f TestOk -m dev

```

## TODO

* middlewares
    * robots
    * file
    * media
    * proxy
    * random
    * downloadtimeout

* cron
* max request limit?
* multi-spider


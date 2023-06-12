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
* 需要返回的item需要实现Item（可以组合ItemUnimplemented）
    * `GetReferer()` 可以获取到referer。
    * UniqueKey 唯一键，不会保存到数据库。可以用作过滤等其他用途。
    * Id 保存主键
    * Data 完整数据
    * 内置item：ItemCsv、ItemJsonl、ItemMongo、ItemMysql、ItemKafka
* 中间件的order不能重复。编写的时候不要忘记`nextRequest()`/`nextResponse()`/`nextItem()`
* 本框架舍弃了pipeline概念，功能合并到middleware。在很多情况下，功能会有交叉，合并后会更方便，同时编写也更简单。
* middleware包括框架内置、自定义公共（internal/middlewares）和自定义爬虫内（和爬虫同module）。
* 框架内置middleware，自定义middleware请参照以下order进行配置。内置中间件order为10的倍数，自定义中间件请避开。
    * stats:10
        * 数据统计
        * 配置 enable_stats: true 是否开启统计，默认开启
    * device:101
        * 修改request设备信息。修改header和tls信息，暂时只支持user-agent随机切换。需要设置`SetPlatforms`和`SetBrowsers`
          限定设备范围。默认不启用。
        * 启用方法。`spider.SetMiddleware(middlewares.NewDeviceMiddleware, 101)`
        * Platforms: Windows/Mac/Android/Iphone/Ipad/Linux
        * Browsers: Chrome/Edge/Safari/FireFox
    * filter:20
        * 过滤重复请求。默认支持的是item保存成功后才会进入去重队列，防止出现请求失败后再次请求却被过滤的问题。所以当请求速度大于保存速度的时候可能会有请求不被过滤的情况。
        * 配置 enable_filter: true 是否开启过滤，默认开启
    * retry:30
        * 如果请求出错，会进行重试。
        * `RetryMaxTimes=10`
        * 配置 enable_retry: true 是否开启重试，默认开启
    * url:40
        * 限制url的长度
        * 配置 enable_url: true 是否开启url长度限制，默认开启
        * 配置 url_length_limit: 2083 url的最长长度默认为2083
    * referer:50
        * 通过设置`referrer_policy`采用不同的referer策略，
        * DefaultReferrerPolicy。会加入请求来源，默认
        * NoReferrerPolicy。不加入请求来源
    * cookie:60
        * 如果之前请求返回cookie，会自动加到后面的请求里
        * 配置 enable_cookie: true 是否开启cookie支持，默认开启
    * http:70
        * 创建request
    * chrome:80
        * 模拟chrome
        * 配置 enable_chrome: true 模拟chrome，默认开启
    * httpAuth:90
        * 通过设置`username`、`password`添加httpAuth认证，
        * 配置 enable_http_auth: false 是否开启httpAuth，默认关闭
    * compress:100
        * 支持 gzip/deflate解压缩
        * 配置 enable_compress: true 是否开启gzip/deflate解压缩，默认开启
    * decode:110
        * 支持gbk、gb2310、big5中文解码
        * 配置 enable_decode: true 是否开启中文解码，默认开启
    * redirect:120
        * 网址重定向，默认支持301、302
        * 配置 enable_redirect: true 是否开启重定向，默认开启
        * 配置 redirect_max_times: 1 重定向最大次数
    * dump:80
        * 控制台打印item.data
        * 配置 enable_dump: true 是否开启打印item，默认开启
    * csv
        * 保存结果到csv文件。
        * 需在在ItemCsv中设置`FileName`，保存的文件名称，不包含.csv
        * 启用方法。`spider.SetMiddleware(middlewares.NewCsvMiddleware, 181)`
    * jsonLines
        * 保存结果到jsonlines文件。
        * 需在在ItemJsonl中设置`FileName`，保存的文件名称，不包含.jsonl
        * 启用方法。`spider.SetMiddleware(middlewares.NewJsonLinesMiddleware, 182)`
    * mongo
        * 保存结果到mongo。
        * 需在在ItemMongo中设置`Collection`，保存的collection
        * 启用方法。`spider.SetMiddleware(middlewares.NewMongoMiddleware, 183)`
    * mysql
        * 保存结果到mysql。
        * 需在在ItemMysql中设置`table`，保存的table
        * 启用方法。`spider.SetMiddleware(middlewares.NewMysqlMiddleware, 184)`
    * kafka
        * 保存结果到kafka。
        * 需在在ItemKafka中设置`Topic`，保存的topic
        * 启用方法。`spider.SetMiddleware(middlewares.NewKafkaMiddleware, 185)`
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
    * 在不影响功能的情况下，可以考虑关闭一些用不到的中间件。可以在配置文件中修改，或者爬虫入口中修改
    * 配置文件:
        * enable_retry: false
        * enable_stats: true
        * enable_filter: true
        * enable_referer: true
        * enable_http_auth: false
        * enable_cookie: true
        * enable_dump: true
        * enable_url: true
        * enable_compress: true
        * enable_decode: true
        * enable_redirect: true
* 爬虫结构
    * 建议按照每个网站（子网站）或者每个业务为一个spider。不必分的太细，也不必把所有的网站和业务都写在一个spider里

### args

* -c config file. must set it. 配置文件，必须配置。
* -f start func. default `Test`. 入口方法，默认`Test`。
* -a args. json string. 额外的参数，用于入口方法调用，非必须项。
* -m mode. default `test`. 启动模式，如`dev`,`prod`等，默认`test`

### config

* log.filename: Log file path. You can replace {name} with -ldflags.
* log.long_file: If set to true, the full file path is logged.
* log.level: DEBUG/INFO/WARN/ERROR
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
* enable_retry: true 是否开启重试，默认开启
* enable_stats: true 是否开启统计，默认开启
* enable_filter: true 是否开启过滤，默认开启
* enable_referer: true 是否开启referer，默认开启
* referrer_policy: DefaultReferrerPolicy 来源政策，默认DefaultReferrerPolicy，可选DefaultReferrerPolicy、NoReferrerPolicy
* enable_http_auth: true 是否开启httpAuth，默认开启
* enable_cookie: true 是否开启cookie，默认开启
* enable_dump: true 是否开启打印item，默认开启
* enable_url: true 是否开启url长度限制，默认开启
* url_length_limit: 2083 url长度限制，默认2083
* enable_compress: true 是否开启gzip/deflate解压缩，默认开启
* enable_decode: true 是否开启中文解码，默认开启
* enable_redirect: true 是否开启重定向，默认开启
* redirect_max_times: 1 重定向最大次数，默认1
* enable_chrome: true 模拟chrome，默认开启

## Example

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


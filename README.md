# go-crawler

[go-crawler](https://github.com/lizongying/go-crawler)
[document](https://pkg.go.dev/github.com/lizongying/go-crawler)

## Feature

* 内置devServer，方便调试开发。
* 完善的配置项，自由性更高。
* 丰富的内置中间件，同时方便自定义中间件。
* 多种解析方式，解析页面更简单。
* 支持多种结果保存方式，按需使用。

## Usage

* 为了方便开发调试，增加了本地devServer，在`-m dev`模式下会默认启用。可以自定义route，仅需要实现`pkg.Route`
  ，然后在spider中通过`AddDevServerRoutes(...pkg.Route)`注册到devServer即可。
* 中间件的order不能重复。编写的时候不要忘记`nextRequest()`/`nextResponse()`/`nextItem()`
* 本框架舍弃了pipeline概念，功能合并到middleware。在很多情况下，功能会有交叉，合并后会更方便，同时编写也更简单。
* middleware包括框架内置、自定义公共（internal/middlewares）和自定义爬虫内（和爬虫同module）。
* 框架内置middleware和默认order，建议自定义ProcessItem的中间件order大于140
    * stats:90
    * device:100
        * 修改request设备信息。修改header和tls信息，暂时只支持user-agent随机切换。需要设置`SetPlatforms`和`SetBrowsers`
          限定设备范围。默认不启用。
        * 启用方法。`spider.SetMiddleware(middlewares.NewDeviceMiddleware, 100)`
    * filter:110
        * 过滤重复请求。默认支持的是item保存成功后才会进入去重队列，防止出现请求失败后再次请求却被过滤的问题。所以当请求速度大于保存速度的时候可能会有请求不被过滤的情况。
    * retry:120
        * 如果请求出错，会进行重试。默认启用，`RetryMaxTimes=3`
    * http:130
    * dump:140 在debug模式下打印item.data
    * csv
        * 保存结果到csv文件。
    * jsonlines
        * 保存结果到jsonlines文件。
    * mongo
        * 保存结果到mongo。
    * mysql
        * 保存结果到mysql。
    * kafka
        * 保存结果到kafka。
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

### args

* -c config file. must set it.
* -f start func. default Test.
* -a args. json string.
* -m mode. default test. prod? dev? or another something.

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
* request.retry_max_times: Request retry max times
* request.http_proto: Request http proto
* dev_addr: dev httpserver addr

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

* once
* cron
* max request limit?


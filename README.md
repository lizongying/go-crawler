# go-crawler

[go-crawler](https://github.com/lizongying/go-crawler)

## Feature

* 为了方便开发调试，增加了本地httpserver，可以通过<code>-m dev</code>启用。可以自定义handle，仅需要实现<code>
  pkg.Route</code>，然后注册即可。
* 编写中间件的时候需要注意，name不能重复。注册的时候order不能重复

## Usage

* log.filename: Log file path. You can replace {name} with -ldflags.
* log.long_file: If set to true, the full file path is logged.
* log.level: DEBUG/INFO/WARN/ERROR
* request.concurrency: Number of request concurrency
* request.interval: Request interval(seconds). If set to 0, it is the default interval(1). If set to a negative number,
  it is 0.
* request.timeout: Request timeout(seconds)
* request.retry_max_times: Request retry max times
* dev_addr: dev httpserver addr

## Example

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

```shell
git clone github.com/lizongying/go-crawler-example
```

## Test

```shell
go run example/testNoLimitSpider/*.go -c dev.yml -f TestNoLimit -m dev

```


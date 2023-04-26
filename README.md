# go-crawler

[go-crawler](https://github.com/lizongying/go-crawler)

## Usage

* log.filename: Log file path. You can replace {name} with -ldflags.
* log.long_file: If set to true, the full file path is logged.
* log.level: DEBUG/INFO/WARN/ERROR
* request.concurrency: Number of request concurrency
* request.interval: Request interval(seconds). If set to 0, it is the default interval. If set to a negative number, it
  is 0.
* request.timeout: Request timeout(seconds)

## Example

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

```shell
git clone github.com/lizongying/go-crawler-example
```



# go-crawler

[go-crawler](https://github.com/lizongying/go-crawler)

## Usage

* log.filename: log file path. you can replace {name} by ldflags.
* log.long_file: if set true, will log full file path.
* log.level: DEBUG/INFO/WARN/ERROR

example

```shell
git clone github.com/lizongying/go-crawler-example
```

build

```shell
make
```

run

```shell
./releases/youtubeSpider -c example.yml
```

---
weight: 2
bookFlatSection: true
title: Usage
---

## Run

```shell
git clone git@github.com:lizongying/go-crawler-example.git my-crawler
cd my-crawler
go run cmd/multi_spider/*.go -c example.yml -n test1 -m once
```

### Build

```shell
make
```

Currently, the framework is updated frequently. Recommended to stay attentive. It's advisable to use the
latest version.

```shell
go get -u github.com/lizongying/go-crawler

# Latest Released Version.
go get -u github.com/lizongying/go-crawler@latest

# Latest Submission (Recommended).
go get -u github.com/lizongying/go-crawler@6f52307

```

### Docker build

```shell
# cross platform
docker buildx create --use

# for linux
docker buildx build --platform linux/amd64 -f ./cmd/test_spider/Dockerfile -t lizongying/go-crawler/test-spider:amd64 . --load

# for mac m1
docker buildx build --platform linux/arm64 -f ./cmd/test_spider/Dockerfile -t lizongying/go-crawler/test-spider:arm64 . --load
```

```shell
docker run -p 8090:8090 -d lizongying/go-crawler/test-spider:arm64 -c example.yml -f TestRedirect -m once
```
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
docker build -f ./cmd/testSpider/Dockerfile -t go-crawler/test-spider:latest . 
```

```shell
docker run -d go-crawler/test-spider:latest spider -c example.yml -f TestRedirect -m once
```
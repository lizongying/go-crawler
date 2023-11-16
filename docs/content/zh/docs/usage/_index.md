---
weight: 2
bookFlatSection: true
title: 使用
---

## 运行

```shell
git clone git@github.com:lizongying/go-crawler-example.git my-crawler
cd my-crawler
go run cmd/multi_spider/*.go -c example.yml -n test1 -m once
```

### 构建

```shell
make

```

目前框架更新较为频繁, 建议保持关注，使用最新版本。
可以在项目里执行更新命令，如:

```shell
go get -u github.com/lizongying/go-crawler

# 最新发布版本
go get -u github.com/lizongying/go-crawler@latest

# 最新提交（推荐）
go get -u github.com/lizongying/go-crawler@6f52307

```

### 容器构建

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
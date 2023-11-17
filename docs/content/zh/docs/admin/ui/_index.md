---
weight: 2
title: 界面
---

## 界面

你可以直接使用https://lizongying.github.io/go-crawler/

[ca](https://github.com/lizongying/go-crawler/blob/main/static/tls/ca.crt)

开发

```shell
npm run dev --prefix ./web/ui

# docs
hugo server --source docs --noBuildLock
```

文档开发

```shell
# docs
hugo server --source docs --noBuildLock

```

构建

web_server 非必须项，你可以直接使用nginx等网络服务

```shell
# ui
make web_ui

# server
make web_server

```

运行

```shell
./releases/web_server
```
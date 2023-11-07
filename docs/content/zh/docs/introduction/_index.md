---
weight: 1
bookFlatSection: true
title: 介绍
---

基于golang实现的爬虫框架，编写简单，性能强劲。内置了丰富的实用中间件，支持多种解析、保存方式，支持分布式部署。

## 运行

```shell
git clone git@github.com:lizongying/go-crawler-example.git my-crawler
cd my-crawler
go run cmd/multi_spider/*.go -c example.yml -n test1 -m once
```

## 功能

* 编写简单，性能强劲。
* 内置多种实用中间件，开发起来更轻松。
* 支持多种解析方式，解析页面更简单。
* 支持多种保存方式，数据存储更灵活。
* 提供了更多的配置选项，配置更丰富。
* 组件支持自定义，功能拓展更自由。
* 内置模拟服务，调试开发更方便。
* 支持分布式部署

### 支持情况

* 解析支持CSS、XPath、Regex、Json
* 支持Json、Csv、Mongo、Mysql、Sqlite、Kafka输出
* 支持gb2312、gb18030、gbk、big5中文解码
* 支持gzip、deflate、brotli解压缩
* 支持分布式
* 支持Redis、Kafka作为消息队列
* 支持自动Cookie、重定向
* 支持BaseAuth认证
* 支持请求重试
* 支持请求过滤
* 支持图片文件下载
* 支持图片处理
* 支持对象存储
* 支持ssl指纹修改
* 支持http2
* 支持随机请求头
* 支持模拟浏览器
* 支持浏览器ajax请求
* 支持模拟服务
* 支持优先级队列
* 支持定时任务、循环任务、单次任务
* 支持基于字段标签的解析
* 支持dns缓存
* 支持中间人代理
* 支持错误记录
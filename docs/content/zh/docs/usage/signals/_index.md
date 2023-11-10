---
weight: 8
title: 信号
---

## 信号

通过信号可以获取爬虫事件。

* `CrawlerChanged`: 程序状态已变化。通过`RegisterCrawlerChanged(FnCrawlerChanged)`注册。
* `SpiderChanged`: 爬虫状态已变化。通过`RegisterSpiderChanged(FnSpiderChanged)`注册。
* `JobChanged`: 计划任务状态已变化。通过`RegisterJobChanged(FnJobChanged)`注册。
* `TaskChanged`: 任务状态已变化。通过`RegisterTaskChanged(FnTaskChanged)`注册。
* `RequestChanged`: 请求状态已变化。通过`RegisterRequestChanged(FnRequestChanged)`注册。
* `ItemChanged`: 数据状态已变化。通过`RegisterItemChanged(FnItemChanged)`注册。
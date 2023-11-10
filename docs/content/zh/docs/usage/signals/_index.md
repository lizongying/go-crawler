---
weight: 8
title: 信号
---

## 信号

通过信号可以获取爬虫事件。

* `CrawlerChanged`: 程序状态已变化。通过`RegisterCrawlerChanged(FnCrawlerChanged)`注册。
* `SpiderChanged`: 爬虫状态已变化。通过`RegisterSpiderChanged(FnSpiderChanged)`注册。
* `JobChanged`: 计划任务状态已变化。通过`RegisterJobChanged(FnJobChanged)`注册。
* `TaskStarted`: 任务已启动。通过`RegisterTaskStarted(FnTaskStarted)`注册。
* `TaskStopped`: 任务已停止。通过`RegisterTaskClosed(FnTaskStopped)`注册。
* `ItemSaved`: 数据已保存。通过`RegisterItemSaved(FnItemSaved)`注册。
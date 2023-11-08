---
weight: 8
title: 信号
---

## 信号

通过信号可以获取爬虫事件。

* `CrawlerStarted`: 程序已启动。通过`RegisterCrawlerStarted(FnCrawlerStarted)`注册。
* `CrawlerStopped`: 程序已停止。通过`RegisterCrawlerClosed(FnCrawlerStopped)`注册。
* `SpiderStarting`: 爬虫启动中。通过`RegisterSpiderStarting(FnSpiderStarting)`注册。
* `SpiderStarted`: 爬虫已启动。通过`RegisterSpiderStarted(FnSpiderStarted)`注册。
* `SpiderStopping`: 爬虫停止中。通过`RegisterSpiderStopping(FnSpiderStopping)`注册。
* `SpiderStopped`: 爬虫已停止。通过`RegisterSpiderClosed(FnSpiderStopped)`注册。
* `JobStarted`: 计划任务已启动。通过`RegisterJobStarted(FnJobStarted)`注册。
* `JobStopped`: 计划任务已停止。通过`RegisterJobClosed(FnJobStopped)`注册。
* `TaskStarted`: 任务已启动。通过`RegisterTaskStarted(FnTaskStarted)`注册。
* `TaskStopped`: 任务已停止。通过`RegisterTaskClosed(FnTaskStopped)`注册。
* `ItemSaved`: 数据已保存。通过`RegisterItemSaved(FnItemSaved)`注册。
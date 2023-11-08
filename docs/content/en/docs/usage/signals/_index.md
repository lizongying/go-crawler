---
weight: 8
title: Signals
---

## Signals

By using signals, it's possible to capture crawler events and perform certain actions.

* `CrawlerStarted`: Indicates the started of the crawler. You can register it
  using `RegisterCrawlerStarted(FnCrawlerStarted)`.
* `CrawlerStopped`: Indicates the stopped of the crawler. You can register it
  using `RegisterCrawlerClosed(FnCrawlerStopped)`.
* `SpiderStarting`: Indicates the starting of the spider. You can register it
  using `RegisterSpiderStarting(FnSpiderStarting)`.
* `SpiderStarted`: Indicates the started of the spider. You can register it
  using `RegisterSpiderStarted(FnSpiderStarted)`.
* `SpiderStopping`: Indicates the stopping of the spider.You can register it
  using `RegisterSpiderStopping(FnSpiderStopping)`.
* `SpiderStopped`: Indicates the stopped of the spider.You can register it
  using `RegisterSpiderClosed(FnSpiderStopped)`.
* `JobStarted`: Indicates the started of the job. You can register it
  using `RegisterJobStarted(FnScheduleStarted)`.
* `JobStopped`: Indicates the stopped of the job. You can register it
  using `RegisterJobClosed(FnJobStopped)`.
* `TaskStarted`: Indicates the started of the task. You can register it using `RegisterTaskStarted(FnTaskStarted)`.
* `TaskStopped`: Indicates the stopped of the task. You can register it using `RegisterTaskClosed(FnTaskStopped)`.
* `ItemSaved`: Indicates the saved of the item. You can register it using `RegisterItemSaved(FnItemSaved)`.
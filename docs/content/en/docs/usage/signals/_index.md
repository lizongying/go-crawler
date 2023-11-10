---
weight: 8
title: Signals
---

## Signals

By using signals, it's possible to capture crawler events and perform certain actions.

* `CrawlerChanged`: Indicates the status changed of the crawler. You can register it
  using `RegisterCrawlerChanged(FnCrawlerChanged)`.
* `SpiderChanged`: Indicates the status changed of the spider.You can register it
  using `RegisterSpiderChanged(FnSpiderChanged)`.
* `JobChanged`: Indicates the status changed of the job. You can register it
  using `RegisterJobChanged(FnJobChanged)`.
* `TaskChanged`: Indicates the status changed of the task. You can register it using `RegisterTaskChanged(FnTaskChanged)`.
* `RequestChanged`: Indicates the status changed of the request. You can register it using `RegisterRequestChanged(FnRequestChanged)`.
* `ItemChanged`: Indicates the status changed of the item. You can register it using `RegisterItemChanged(FnItemChanged)`.
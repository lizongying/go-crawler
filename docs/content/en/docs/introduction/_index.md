---
weight: 1
bookFlatSection: true
title: Introduction
---

# Go Crawler Documentation

## About Go Crawler

A web crawling framework implemented in Golang, it is simple to write and delivers powerful performance. It comes with a
wide range of practical middleware and supports various parsing and storage methods. Additionally, it supports
distributed deployment.

## Run

```shell
git clone git@github.com:lizongying/go-crawler-example.git my-crawler
cd my-crawler
go run cmd/multi_spider/*.go -c example.yml -n test1 -m once
```

## Feature

* Simple to write, yet powerful in performance.
* Built-in various practical middleware for easier development.
* Supports multiple parsing methods for simpler page parsing.
* Supports multiple storage methods for more flexible data storage.
* Provides numerous configuration options for richer customization.
* Allows customizations for components, providing more freedom for feature extensions.
* Includes a built-in mock Server for convenient debugging and development.
* It supports distributed deployment.

### Support Summary

* Parsing supports CSS, XPath, Regex, and JSON.
* Output supports JSON, CSV, MongoDB, MySQL, Sqlite, and Kafka.
* Supports Chinese decoding for gb2312, gb18030, gbk, big5 character encodings.
* Supports gzip, deflate, and brotli decompression.
* Supports distributed processing.
* Supports Redis and Kafka as message queues.
* Supports automatic handling of cookies and redirects.
* Supports BaseAuth authentication.
* Supports request retry.
* Supports request filtering.
* Supports image file downloads.
* Supports image processing.
* Supports object storage.
* Supports SSL fingerprint modification.
* Supports HTTP/2.
* Supports random request headers.
* Browser simulation is supported.
* Supports browser AJAX requests.
* Mock server is supported.
* Priority queue is supported.
* Supports scheduled tasks, recurring tasks, and one-time tasks.
* Supports parsing based on field labels.
* Supports DNS Cache.
* Supports MITM
* Supports error logging
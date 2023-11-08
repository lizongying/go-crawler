---
weight: 14
title: 基于字段标签的网页解析
---

## 基于字段标签的网页解析

在本框架里，返回的数据是个结构体。我们仅需在字段上加上解析规则的标签，框架会自动进行网页解析，看起来非常简洁。
这对于一些简单的爬虫来说，更加便捷高效。特别是需要添加大量的通用爬虫时，仅需要配置这些标签就可以直接解析。
比如：

```go
type DataRanks struct {
Data []struct {
Name           string  `_json:"name"`
FullName       string  `_json:"fullname"`
Code           string  `_json:"code"`
MarketBalue    int     `_json:"market_value"`
MarketValueUsd int     `_json:"market_value_usd"`
Marketcap      int     `_json:"marketcap"`
Turnoverrate   float32 `_json:"turnoverrate"`
} `_json:"data"`
}

```

data可以设置根解析`_json:"data"`， 也就是里面的字段都是在根解析下，仅需要写属性就可以了。`_json:"name"`

根标签和子标签可以混用，比如根标签用xpth，子标签用json

可以使用如下标签：

* `_json:""` gjson 格式
* `_xpath:""` xpath 格式
* `_css:""` css 格式
* `_re:""` re 格式
---
weight: 14
title: Web Page Parsing Based on Field Tags
---

## Web Page Parsing Based on Field Tags

In this framework, the returned data is a struct. We only need to add parsing rule tags to the fields, and the framework
will automatically perform web page parsing, making it appear very clean and concise.

For some simple web scraping tasks, this approach is more convenient and efficient. Especially when you need to create a
large number of generic web scrapers, you can directly configure these tags for parsing.

For example:

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

You can set the root parsing for `data` as `_json:"data"`, meaning that the fields inside it are all parsed under the
root. For example, `_json:"name"`.

You can mix and match root and sub-tags, for instance, use XPath for the root and JSON for the sub-tags.

You can use the following tags:

* `_json:""` for gjson format
* `_xpath:""` for XPath format
* `_css:""` for CSS format
* `_re:""` for regular expression (regex) format
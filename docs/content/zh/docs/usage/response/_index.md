---
weight: 7
title: 响应
---

## 响应

框架内置了多个解析模块。您可以根据具体的爬虫需求，选择适合您的解析方式。

* `Xpath() (*xpath.Selector, error)` `MustXpath() *xpath.Selector`

  返回Xpath选择器，具体语法请参考 [go-xpath](https://github.com/lizongying/go-xpath)

* `Css() (*css.Selector, error)` `MustCss() *css.Selector`

  返回CSS选择器，具体语法请参考 [go-query](https://github.com/lizongying/go-css)

* `Json() (*gson.Selector, error)` `MustJson() gjson.Result`

  返回gjson选择器，具体语法请参考 [go-json](https://github.com/lizongying/go-json)

* `Re() (*re.Selector, error)` `MustRe() *re.Selector`

  返回正则选择器，具体语法请参考 [go-re](https://github.com/lizongying/go-re)

* `AllLink() []*url.URL`

  可以获取response中的所有链接。

* `BodyText() string`

  可以获取清理过html标签的正文，处理比较粗糙。

* `AbsoluteURL(relativeUrl string) (absoluteURL *url.URL, err error)`

  可以获取url绝对地址
---
weight: 7
title: Response
---

## Response

The framework comes with several built-in parsing modules. You can choose the one that suits your specific spider's
needs.

* `Xpath() (*xpath.Selector, error)` `MustXpath() *xpath.Selector`

  Returns an XPath selector, for specific syntax, please refer
  to [go-xpath](https://github.com/lizongying/go-xpath).

* `Css() (*css.Selector, error)` `MustCss() *css.Selector`

  Returns a CSS selector, for specific syntax, please refer to [go-css](https://github.com/lizongying/go-css).

* `Json() (*gson.Selector, error)` `MustJson() gjson.Result`

  Returns a gjson selector, for specific syntax, please refer to [go-json](https://github.com/lizongying/go-json).

* `Re() (*re.Selector, error)` `MustRe() *re.Selector`

  Returns a regular expression selector, for specific syntax, please refer
  to [go-re](https://github.com/lizongying/go-re).

* `AllLink() []*url.URL`

  Retrieves all links from the response.

* `BodyText() string`

  Retrieves the cleaned text content without HTML tags, the handling may be rough.

* `AbsoluteURL(relativeUrl string) (absoluteURL *url.URL, err error)`

  Retrieves the absolute URL for a given relative URL.
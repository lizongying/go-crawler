---
weight: 6
title: 请求
---

## 请求

创建一个请求
 
```go
  // 创建一个请求
req := request.NewRequest()

// 设置URL
req.SetUrl("")

// 设置请求方法
req.SetMethod(http.MethodGet)

// 设置一个请求头
req.SetHeader("name", "value")

// 一次设置所有请求头
req.SetHeaders(map[string]string{"name1": "value1", "name2": "value2"})

// 设置字符串请求体
req.SetBodyStr(``)

// 设置字节数组请求体
req.SetBodyBytes([]byte(``))

// 设置解析方法
var parse func (ctx pkg.Context, response pkg.Response) (err error)
req.SetCallBack(parse)

// 返回请求
s.MustYieldRequest(ctx, req)

// 建议这么写，更简单
s.MustYieldRequest(ctx, request.NewRequest().
SetUrl("").
SetBodyStr(``).
SetExtra(&Extra{}).
SetCallBack(s.Parse))
```

创建请求的简单方法

```go
_ = request.Get()
_ = request.Post()
_ = request.Head()
_ = request.Options()
_ = request.Delete()
_ = request.Put()
_ = request.Patch()
_ = request.Trace()
```

* `SetFingerprint(string) Request`

  现在很多网站都对ssl指纹进行了风控，通过设置此参数，可以进行伪装。

  如果fingerprint是`pkg.Browser`,框架会自动选择此浏览器合适的指纹。

  如果fingerprint是ja3格式指纹，框架会应用此ssl指纹。

  如果fingerprint为空，框架会根据user-agent进行选择。

  注意框架仅会在`enable_ja3 = true` 的情况下进行修改，默认使用golang自身的ssl配置。

* `SetClient(Client) Request`

  一些网站可能会识别浏览器指纹，这种情况下建议使用模拟浏览器。

  设置Client为`pkg.Browser`后，框架会自动启用模拟浏览器。

* `SetAjax(bool) Request` 如果需要使用无头浏览器，并且请求是ajax，请设置此选项为true，框架会进行xhr请求。可能需要设置referrer。
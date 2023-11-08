---
weight: 6
title: Request
---

## Request

Build a request.

```go
  // Build a request.
req := request.NewRequest()

// Set the URL
req.SetUrl("")

// Set the request method.
req.SetMethod(http.MethodGet)

// Set the request header.
req.SetHeader("name", "value")

// Set all request headers at once.
req.SetHeaders(map[string]string{"name1": "value1", "name2": "value2"})

// Set the request content string.
req.SetBodyStr(``)

// Set the request content bytes.
req.SetBodyBytes([]byte(``))

// Set the parsing method
var parse func (ctx pkg.Context, response pkg.Response) (err error)
req.SetCallBack(parse)

// Send the request
s.MustYieldRequest(ctx, req)

// Suggest writing it this way, simpler.
s.MustYieldRequest(ctx, request.NewRequest().
SetUrl("").
SetBodyStr(``).
SetExtra(&Extra{}).
SetCallBack(s.Parse))

 ```

Create a request using a simple method.

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

  Many websites nowadays implement security measures based on SSL fingerprints. By
  setting this parameter, you can perform disguising. If the fingerprint is `pkg.Browser`, the framework will
  automatically select a suitable fingerprint for this browser. If the fingerprint is in the ja3 format, the
  framework will apply this SSL fingerprint. If the fingerprint is empty, the framework will choose based on the
  user-agent. Note that the framework will only make modifications when `enable_ja3 = true`, and it uses the default
  SSL configuration of the Go programming language.

* `SetClient(Client) Request`

  Some websites may detect browser fingerprints. In such cases, it is recommended to use browser simulation.

  After setting the client to `pkg.Browser`, the framework will automatically enable browser simulation.

* `SetAjax(bool) Request`

  If you need to use a headless browser and the request is an AJAX request, please set
  this option to true. The framework will handle the request as an XHR (XMLHttpRequest) request. You may also
  need to set the referrer.
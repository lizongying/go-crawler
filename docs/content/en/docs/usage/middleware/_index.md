---
weight: 4
title: Middleware
---

## Middleware

Middleware and Pipeline include built-in ones, commonly used custom ones (internal/middlewares, internal/pipelines),
and custom ones defined within the spider's module.
Please make sure that the order values for different middleware and pipelines are not duplicated.If there are
duplicate order values, the later middleware or pipeline will replace the earlier ones.

In the framework, built-in middleware has pre-defined `order` values that are multiples of 10, such as 10, 20, 30, and
so on.To avoid conflicts with the `order` values of built-in middleware, it is recommended to choose
different `order` values when defining custom middleware.

When customizing middleware, arrange them in the expected execution order based on their functionalities and
requirements.Make sure that middleware with lower `order` values is executed first, followed by middleware with
higher `order` values.

Built-in middleware and custom middleware can use the default `order` values.If you need to change the
default `order` value, `spider.WithOptions(pkg.WithMiddleware(new(middleware), order)` to
enable the middleware with the specified `order` value.

The following are the built-in middleware with their respective `order` values:

* custom: 10
    * Custom middleware.
    * `spider.WithOptions(pkg.WithCustomMiddleware(new(CustomMiddleware))`
* retry: 20
    * Request retry middleware used for retrying requests when they fail.
    * The default maximum number of retries is 10. You can control whether to enable this middleware by configuring
      the `enable_retry_middleware` option, which is enabled by default.
    * `spider.WithOptions(pkg.WithRetryMiddleware()`
* dump: 30
    * Console dump middleware used for printing detailed information of item.data, including request and response
      details.
    * You can control whether to enable this middleware by configuring the `enable_dump_middleware` option, which is
      enabled by default.
    * `spider.WithOptions(pkg.WithDumpMiddleware()`
* proxy: 40
    * Proxy switch middleware used for switching proxies for requests.
    * You can control whether to enable this middleware by configuring the `enable_proxy_middleware` option, which
      is enabled by default.
    * `spider.WithOptions(pkg.WithProxyMiddleware()`
* robotsTxt: 50
    * Robots.txt support middleware for handling robots.txt files of websites.
    * You can control whether to enable this middleware by configuring the `enable_robots_txt_middleware` option,
      which is disabled by default.
    * `spider.WithOptions(pkg.WithRobotsTxtMiddleware()`
* filter: 60
    * Request deduplication middleware used for filtering duplicate requests.By default, items are added to the
      deduplication queue only after being successfully saved.
    * You can control whether to enable this middleware by configuring the `enable_filter_middleware` option, which
      is enabled by default.
    * `spider.WithOptions(pkg.WithFilterMiddleware()`
* file: 70
    * Automatic file information addition middleware used for automatically adding file information to requests.
    * You can control whether to enable this middleware by configuring the `enable_file_middleware` option, which is
      disabled by default.
    * `spider.WithOptions(pkg.WithFileMiddleware()`
* image: 80
    * Automatic image information addition middleware used for automatically adding image information to requests.
    * You can control whether to enable this middleware by configuring the `enable_image_middleware` option, which
      is disabled by default.
    * `spider.WithOptions(pkg.WithImageMiddleware()`
* url: 90
    * URL length limiting middleware used for limiting the length of requests' URLs.
    * You can control whether to enable this middleware and set the maximum URL length by configuring
      the `enable_url_middleware` and `url_length_limit` options, respectively.Both options are enabled and set to
      2083 by default.
    * `spider.WithOptions(pkg.WithUrlMiddleware()`
* referrer: 100
    * Automatic referrer addition middleware used for automatically adding the referrer to requests.
    * You can choose different referrer policies based on the `referrer_policy` configuration
      option.`DefaultReferrerPolicy` includes the request source, while `NoReferrerPolicy` does not include the
      request source.
    * You can control whether to enable this middleware by configuring the `enable_referrer_middleware` option,
      which is enabled by default.
    * `spider.WithOptions(pkg.WithReferrerMiddleware()`
* cookie: 110
    * Automatic cookie addition middleware used for automatically adding cookies returned from previous requests to
      subsequent requests.
    * You can control whether to enable this middleware by configuring the `enable_cookie_middleware` option, which
      is enabled by default.
    * `spider.WithOptions(pkg.WithCookieMiddleware()`
* redirect: 120
    * Website redirection middleware used for handling URL redirection.By default, it supports 301 and 302
      redirects.
    * You can control whether to enable this middleware and set the maximum number of redirections by
      configuring
      the `enable_redirect_middleware` and `redirect_max_times` options, respectively.Both options are enabled
      and set
      to 1 by default.
    * `spider.WithOptions(pkg.WithRedirectMiddleware()`
* chrome: 130
    * Chrome simulation middleware used for simulating a Chrome browser.
    * You can control whether to enable this middleware by configuring the `enable_chrome_middleware` option,
      which is
      enabled by default.
    * `spider.WithOptions(pkg.WithChromeMiddleware()`
* httpAuth: 140
    * HTTP authentication middleware used for performing HTTP authentication by providing a username and
      password.
    * You need to set the username and password in the specific request.You can control whether to enable this
      middleware by configuring the `enable_http_auth_middleware` option, which is disabled by default.
    * `spider.WithOptions(pkg.WithHttpAuthMiddleware()`
* compress: 150
    * Gzip/deflate/br decompression middleware used for handling response compression encoding.
    * You can control whether to enable this middleware by configuring the `enable_compress_middleware` option,
      which is
      enabled by default.
    * `spider.WithOptions(pkg.WithCompressMiddleware()`
* decode: 160
    * Chinese decoding middleware used for decoding responses with GBK, GB2312, GB18030,and Big5 encodings.
    * You can control whether to enable this middleware by configuring the `enable_decode_middleware` option,
      which is
      enabled by default.
    * `spider.WithOptions(pkg.WithDecodeMiddleware()`
* device: 170
    * Modify request device information middleware used for modifying the device information of requests,
      including
      request headers and TLS information.Currently, only User-Agent random switching is supported.
    * You need to set the device range (Platforms) and browser range (Browsers).
    * Platforms: Windows/Mac/Android/Iphone/Ipad/Linux
    * Browsers: Chrome/Edge/Safari/FireFox
    * You can control whether to enable this middleware by configuring the `enable_device_middleware` option,
      which is
      disabled by default.
    * `spider.WithOptions(pkg.WithDeviceMiddleware()`
* http: 200
    * Create request middleware used for creating HTTP requests.
    * You can control whether to enable this middleware by configuring the `enable_http_middleware` option,
      which is
      enabled by default.
    * `spider.WithOptions(pkg.WithHttpMiddleware()`
* stats: 210
    * Data statistics middleware used for collecting statistics on requests, responses, and processing in the
      spider.
    * You can control whether to enable this middleware by configuring the `enable_stats_middleware` option,
      which is
      enabled by default.
    * `spider.WithOptions(pkg.WithStatsMiddleware()`
* recordError: 220
    * Error recording middleware used to log requests and errors occurring during request processing.
    * It can be enabled or disabled using the configuration option `enable_record_error_middleware`, disabled by
      default.
    * `spider.WithOptions(pkg.WithRecordErrorMiddleware()`
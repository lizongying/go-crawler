---
weight: 4
title: 中间件
---

### 中间件

middleware/pipeline包括框架内置、公共自定义（internal/middlewares，internal/pipelines）和爬虫内自定义（和爬虫同module）。
请确保不同中间件和Pipeline的order值不重复。如果有重复的order值，后面的中间件或Pipeline将替换前面的中间件或Pipeline。

在框架中，内置的中间件具有预定义的order值，这些order值是10的倍数，例如10、20、30等。
为了避免与内置中间件的order冲突，建议自定义中间件时选择不同的order值。
当您自定义中间件时，请选择避开内置中间件的order值。
根据中间件的功能和需求，按照预期的执行顺序进行配置。确保较低order值的中间件先执行，然后依次执行较高order值的中间件。
内置的中间件和自定义中间件使用默认的order值即可。
如果需要改变默认的order值，需要`spider.WithOptions(pkg.WithMiddleware(new(middleware), order)`启用该中间件并应用该order值。

* custom: 10
    * 自定义中间件
    * `spider.WithOptions(pkg.WithCustomMiddleware(new(CustomMiddleware))`
* retry: 20
    * 请求重试中间件，用于在请求失败时进行重试。
    * 默认最大重试次数为10。可以通过配置项enable_retry_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithRetryMiddleware()`
* dump: 30
    * 控制台打印item.data中间件，用于打印请求和响应的详细信息。
    * 可以通过配置项enable_dump_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithDumpMiddleware()`
* proxy: 40
    * 用于切换请求使用的代理。
    * 可以通过配置项enable_proxy_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithProxyMiddleware()`
* robotsTxt: 50
    * robots.txt支持中间件，用于支持爬取网站的robots.txt文件。
    * 可以通过配置项enable_robots_txt_middleware来启用或禁用，默认禁用。
    * `spider.WithOptions(pkg.WithRobotsTxtMiddleware()`
* filter: 60
    * 过滤重复请求中间件，用于过滤重复的请求。默认只有在Item保存成功后才会进入去重队列。
    * 可以通过配置项enable_filter_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithFilterMiddleware()`
* file: 70
    * 自动添加文件信息中间件，用于自动添加文件信息到请求中。
    * 可以通过配置项enable_file_middleware来启用或禁用，默认禁用。
    * `spider.WithOptions(pkg.WithFileMiddleware()`
* image: 80
    * 自动添加图片的宽高等信息中间件
    * 用于自动添加图片信息到请求中。可以通过配置项enable_image_middleware来启用或禁用，默认禁用。
    * `spider.WithOptions(pkg.WithImageMiddleware()`
* url: 90
    * 限制URL长度中间件，用于限制请求的URL长度。
    * 可以通过配置项enable_url_middleware和url_length_limit来启用和设置最长URL长度，默认启用和最长长度为2083。
    * `spider.WithOptions(pkg.WithUrlMiddleware()`
* referrer: 100
    * 自动添加Referrer中间件，用于自动添加Referrer到请求中。
    * 可以根据referrer_policy配置项选择不同的Referrer策略，DefaultReferrerPolicy会加入请求来源，NoReferrerPolicy不加入请求来源
    * 配置 enable_referrer_middleware: true 是否开启自动添加referrer，默认启用。
    * `spider.WithOptions(pkg.WithReferrerMiddleware()`
* cookie: 110
    * 自动添加Cookie中间件，用于自动添加之前请求返回的Cookie到后续请求中。
    * 可以通过配置项enable_cookie_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithCookieMiddleware()`
* redirect: 120
    * 网址重定向中间件，用于处理网址重定向，默认支持301和302重定向。
    * 可以通过配置项enable_redirect_middleware和redirect_max_times来启用和设置最大重定向次数，默认启用和最大次数为1。
    * `spider.WithOptions(pkg.WithRedirectMiddleware()`
* chrome: 130
    * 模拟Chrome中间件，用于模拟Chrome浏览器。
    * 可以通过配置项enable_chrome_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithChromeMiddleware()`
* httpAuth: 140
    * HTTP认证中间件，通过提供用户名（username）和密码（password）进行HTTP认证。
    * 需要在具体的请求中设置用户名和密码。可以通过配置项enable_http_auth_middleware来启用或禁用，默认禁用。
    * `spider.WithOptions(pkg.WithHttpAuthMiddleware()`
* compress: 150
    * 支持gzip/deflate/br解压缩中间件，用于处理响应的压缩编码。
    * 可以通过配置项enable_compress_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithCompressMiddleware()`
* decode: 160
    * 中文解码中间件，支持对响应中的GBK、GB2312、GB18030和Big5编码进行解码。
    * 可以通过配置项enable_decode_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithDecodeMiddleware()`
* device: 170
    * 修改请求设备信息中间件，用于修改请求的设备信息，包括请求头（header）和TLS信息。目前只支持User-Agent随机切换。
    * 需要设置设备范围（Platforms）和浏览器范围（Browsers）。
    * Platforms: Windows/Mac/Android/Iphone/Ipad/Linux
    * Browsers: Chrome/Edge/Safari/FireFox
    * 可以通过配置项enable_device_middleware来启用或禁用，默认禁用。
    * `spider.WithOptions(pkg.WithDeviceMiddleware()`
* http: 200
    * 创建请求中间件，用于创建HTTP请求。
    * 可以通过配置项enable_http_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithHttpMiddleware()`
* stats: 210
    * 数据统计中间件，用于统计爬虫的请求、响应和处理情况。
    * 可以通过配置项enable_stats_middleware来启用或禁用，默认启用。
    * `spider.WithOptions(pkg.WithStatsMiddleware()`
* recordError: 220
    * 错误记录中间件，用于记录请求，以及请求和解析中出现的错误。
    * 可以通过配置项enable_record_error_middleware来启用或禁用，默认禁用。
    * `spider.WithOptions(pkg.WithRecordErrorMiddleware())`
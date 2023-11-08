---
weight: 11
title: 模拟服务
---

## 模拟服务

为了方便开发和调试，框架内置了本地mockServer，在`mock_server.enable: true`配置下会启用。
通过使用本地mockServer，您可以在开发和调试过程中更方便地模拟和观察网络请求和响应，以及处理自定义路由逻辑。
这为开发者提供了一个便捷的工具，有助于快速定位和解决问题。
您可以自定义路由（route），只需要实现`pkg.Route` 接口，并通过在Spider中调用`AddMockServerRoutes(...pkg.Route)`
方法将其注册到mockServer中。

* 支持http和https，您可以通过设置`mock_server`选项来指定mockServer的URL。
  `http://localhost:8081`表示使用HTTP协议，`https://localhost:8081`表示使用HTTPS协议。
* 默认显示JA3指纹。JA3是一种用于TLS客户端指纹识别的算法，它可以显示与服务器建立连接时客户端使用的TLS版本和加密套件等信息。
* 您可以使用tls工具来生成服务器的私钥和证书，以便在mockServer中使用HTTPS。tls工具可以帮助您生成自签名的证书，用于本地开发和测试环境。
* mockServer内置了多种Route，这些Route提供了丰富的功能，可以模拟各种网络情景，帮助进行开发和调试。
  您可以根据需要选择合适的route，并将其配置到mockServer中，以模拟特定的网络响应和行为。
    * BadGatewayRoute 模拟返回502状态码
    * Big5Route 模拟使用big5编码
    * BrotliRoute 模拟使用brotli压缩
    * CookieRoute 模拟返回cookie
    * DeflateRoute 模拟使用Deflate压缩
    * FileRoute 模拟输出文件
    * Gb2312Route 模拟使用gb2312编码
    * Gb18030Route 模拟使用gb18030编码
    * GbkRoute 模拟使用gbk编码
    * GzipRoute 模拟使用gzip压缩
    * HelloRoute 打印请求的header和body信息
    * HtmlRoute 模拟返回html静态文件，可以把html文件放在/static/html/目录内，用于网页解析测试，不用重复请求
    * HttpAuthRoute 模拟http-auth认证
    * InternalServerErrorRoute 模拟返回500状态码
    * OkRoute 模拟正常输出，返回200状态码
    * RateLimiterRoute 模拟速率限制，目前基于全部请求，不区分用户。可与HttpAuthRoute配合使用。
    * RedirectRoute 模拟302临时跳转，需要同时启用OkRoute
    * RobotsTxtRoute 返回robots.txt文件
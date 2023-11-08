---
weight: 11
title: Mock Server
---

## Mock Server

To facilitate development and debugging, the framework comes with a built-in local MockServer that can be enabled by
setting `mock_server.enable: true` in the configuration. By using the local MockServer, you can more easily simulate
and
observe network requests and responses, as well as handle custom route logic. This provides developers with a
convenient tool to quickly locate and resolve issues.

You can customize routes by implementing the `pkg.Route` interface and registering them with the MockServer in the
spider by calling `AddMockServerRoutes(...pkg.Route)`.

* The MockServer supports both HTTP and HTTPS, and you can specify the MockServer's URL by setting the `mock_server`
  option. `http://localhost:8081` represents using HTTP protocol, and `https://localhost:8081` represents using
  HTTPS protocol.
* By default, the MockServer displays JA3 fingerprints. JA3 is an algorithm used for TLS client fingerprinting, and
  it shows information about the TLS version and cipher suites used by the client when establishing a connection
  with the server.
* You can use the tls tool to generate the server's private key and certificate for use with HTTPS in the
  MockServer.
  The tls tool can help you generate self-signed certificates for local development and testing environments.
* The MockServer includes multiple built-in routes that provide rich functionalities to simulate various network
  scenarios and assist in development and debugging. You can choose the appropriate route based on your needs and
  configure it in the MockServer to simulate specific network responses and behaviors.

    * BadGatewayRoute: Simulates returning a 502 status code.
    * Big5Route: Simulates using the big5 encoding.
    * BrotliRoute: Simulates using brotli compression.
    * CookieRoute: Simulates returning cookies.
    * DeflateRoute: Simulates using Deflate compression.
    * FileRoute: Simulates outputting files.
    * Gb2312Route: Simulates using the gb2312 encoding.
    * Gb18030Route: Simulates using the gb18030 encoding.
    * GbkRoute: Simulates using the gbk encoding.
    * GzipRoute: Simulates using gzip compression.
    * HelloRoute: Prints the header and body information of the request.
    * HtmlRoute: simulates the return of HTML static files. You can place HTML files inside the `/static/html/`
      directory for web parsing testing purposes, eliminating the need for redundant requests.
    * HttpAuthRoute: Simulates http-auth authentication.
    * InternalServerErrorRoute: Simulates returning a 500 status code.
    * OkRoute: Simulates normal output, returning a 200 status code.
    * RateLimiterRoute: Simulates rate limiting, currently based on all requests and not differentiated by users.
      Can be used in conjunction with HttpAuthRoute.
    * RedirectRoute: Simulates a 302 temporary redirect, requires enabling OkRoute simultaneously.
    * RobotsTxtRoute: Returns the robots.txt file.
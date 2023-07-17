# go-crawler

基于golang实现的爬虫框架，编写简单，性能强劲，内置了丰富的实用中间件，支持多种解析、保存方式。

[go-crawler](https://github.com/lizongying/go-crawler)
[document](https://pkg.go.dev/github.com/lizongying/go-crawler)

## Feature

* 编写简单，性能强劲。
* 内置多种实用中间件，开发起来更轻松。
* 支持多种解析方式，解析页面更简单。
* 支持多种保存方式，数据存储更灵活。
* 提供了丰富的配置选项，配置更自由。
* 组件支持自定义，扩展功能更简单。
* 内置开发服务，调试开发更方便。

## Usage

* 基本架构
    * Spider：在Spider里可以发起请求和解析内容。您需要使用`spider.SetName(name)`方法为每个Spider设置一个唯一名称。
    * BaseSpider：BaseSpider实现了Spider的公共方法，避免了在每个Spider中重复编写相同的代码。
    * Crawler：Crawler集成了Spider、Downloader（下载器）、Exporter（导出器）、Scheduler（调度器）等组件，是爬虫的逻辑处理中心。
* crawler选项。
    * WithMode 设置模式（Mode）。
    * WithPlatforms 设置浏览器平台（Platforms）。
    * WithBrowsers 设置浏览器（Browsers）。
    * WithLogger 设置日志（Logger）。
    * WithFilter 设置过滤器（Filter）。
    * WithDownloader 设置下载器（Downloader）。
    * WithExporter 设置导出器（Exporter）。
    * WithMiddleware 设置中间件（Middleware）。
    * WithStatsMiddleware 设置统计中间件，用于记录和统计爬虫的性能和运行情况。
    * WithDumpMiddleware 设置打印中间件。
    * WithProxyMiddleware 设置代理中间件，用于使用代理服务器进行爬取。
    * WithRobotsTxtMiddleware 设置开启robots.txt支持中间件，用于遵守网站的 robots.txt 规则。
    * WithFilterMiddleware 设置过滤器中间件，用于过滤已处理的请求。
    * WithFileMiddleware 设置文件中间件，用于处理文件下载请求。
    * WithImageMiddleware 设置图像中间件，用于处理图像下载请求。
    * WithHttpMiddleware 设置 HTTP 中间件。
    * WithRetryMiddleware 设置重试中间件，用于在请求失败时进行自动重试。
    * WithUrlMiddleware 设置 URL 中间件。
    * WithReferrerMiddleware 设置 Referrer 中间件，用于自动设置请求的 Referrer 头。
    * WithCookieMiddleware 设置 Cookie 中间件，用于处理请求和响应中的 Cookie，自动在接下来的请求设置之前的 Cookie。
    * WithRedirectMiddleware 设置重定向中间件，用于自动处理请求的重定向，跟随重定向链接并获取最终响应。
    * WithChromeMiddleware 设置 Chrome 中间件，用于模拟 Chrome 浏览器。
    * WithHttpAuthMiddleware 设置开启HTTP认证中间件，用于处理需要认证的网站。
    * WithCompressMiddleware 设置压缩中间件，用于处理请求和响应的压缩。当爬虫发送请求或接收响应时，该中间件可以自动处理压缩算法，解压缩请求或响应的内容。
    * WithDecodeMiddleware 设置解码中间件，用于处理请求和响应的解码操作。该中间件可以处理请求或响应中的编码内容。
    * WithDeviceMiddleware 设置开启设备模拟中间件
    * WithCustomMiddleware 设置自定义中间件，允许用户定义自己的中间件组件。
    * WithPipeline 设置Pipeline，用于处理爬取的数据并进行后续操作。
    * WithDumpPipeline 设置打印管道。
    * WithFilePipeline 设置文件管道，用于处理爬取的文件数据，将文件保存到指定位置。
    * WithImagePipeline 设置图像管道，用于处理爬取的图像数据，将保存图像到指定位置。
    * WithFilterPipeline 设置过滤器管道，用于过滤爬取过的数据。
    * WithCsvPipeline 设置 CSV 数据处理管道，将爬取的数据保存为 CSV 格式。
    * WithJsonLinesPipeline 设置 JSON Lines 数据处理管道，将爬取的数据保存为 JSON Lines 格式。
    * WithMongoPipeline 设置 MongoDB 数据处理管道，将爬取的数据保存到 MongoDB 数据库。
    * WithMysqlPipeline 设置 MySQL 数据处理管道，将爬取的数据保存到 MySQL 数据库。
    * WithKafkaPipeline 设置 Kafka 数据处理管道，将爬取的数据发送到 Kafka 消息队列。
    * WithCustomPipeline 设置自定义数据处理管道。
    * WithRetryMaxTimes 设置请求的最大重试次数（RetryMaxTimes）。
    * WithTimeout 设置请求的超时时间（Timeout）。
    * WithInterval 设置请求的间隔时间（Interval）。
    * WithOkHttpCodes 设置正常的HTTP状态码（OkHttpCodes）。
* Item类需要实现Item接口（可以组合ItemUnimplemented）
    * `GetReferrer()` 可以获取到referrer。
    * UniqueKey属性作为唯一键用于过滤和其他用途
    * Id属性用于保存主键
    * Data属性用于保存完整数据（必须是指针类型）
    * 内置Item实现：框架提供了一些内置的Item实现，如ItemNone、ItemCsv、ItemJsonl、ItemMongo、ItemMysql、ItemKafka等。
      您可以根据需要开启相应的Pipeline，以实现数据的保存功能。
* middleware/pipeline包括框架内置、公共自定义（internal/middlewares，internal/pipelines）和爬虫内自定义（和爬虫同module）。
* 请确保不同中间件和Pipeline的order值不重复。如果有重复的order值，后面的中间件或Pipeline将替换前面的中间件或Pipeline。
* 在框架中，内置的中间件具有预定义的order值，这些order值是10的倍数，例如10、20、30等。
  为了避免与内置中间件的order冲突，建议自定义中间件时选择不同的order值。
  当您定义自己的中间件时，请选择避开内置中间件的order值。例如，您可以选择使用11、12、13等不同的order值来定义自定义中间件。
  根据中间件的功能和需求，按照预期的执行顺序进行配置。确保较低order值的中间件先执行，然后依次执行较高order值的中间件。
  内置的中间件和自定义中间件使用默认的order值即可。
  如果需要改变默认的order值，需要在NewApp中加入crawler选项`pkg.WithMiddleware(new(middleware), order)`启用该中间件并应用该order值。
    * stats: 10
        * 数据统计中间件，用于统计爬虫的请求、响应和处理情况。
        * 可以通过配置项enable_stats_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithStatsMiddleware()`
    * dump: 20
        * 控制台打印item.data中间件，用于打印请求和响应的详细信息。
        * 可以通过配置项enable_dump_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithDumpMiddleware()`
    * proxy: 30
        * 用于切换请求使用的代理。
        * 可以通过配置项enable_proxy_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithProxyMiddleware()`
    * robotsTxt: 40
        * robots.txt支持中间件，用于支持爬取网站的robots.txt文件。
        * 可以通过配置项enable_robots_txt_middleware来启用或禁用，默认禁用。
        * 在NewApp中加入crawler选项`pkg.WithRobotsTxtMiddleware()`
    * filter: 50
        * 过滤重复请求中间件，用于过滤重复的请求。默认只有在Item保存成功后才会进入去重队列。
        * 可以通过配置项enable_filter_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithFilterMiddleware()`
    * file: 60
        * 自动添加文件信息中间件，用于自动添加文件信息到请求中。
        * 可以通过配置项enable_file_middleware来启用或禁用，默认禁用。
        * 在NewApp中加入crawler选项`pkg.WithFileMiddleware()`
    * image: 70
        * 自动添加图片的宽高等信息中间件
        * 用于自动添加图片信息到请求中。可以通过配置项enable_image_middleware来启用或禁用，默认禁用。
        * 在NewApp中加入crawler选项`pkg.WithImageMiddleware()`
    * http: 80
        * 创建请求中间件，用于创建HTTP请求。
        * 可以通过配置项enable_http_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithHttpMiddleware()`
    * retry: 90
        * 请求重试中间件，用于在请求失败时进行重试。
        * 默认最大重试次数为10。可以通过配置项enable_retry_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithRetryMiddleware()`
    * url: 100
        * 限制URL长度中间件，用于限制请求的URL长度。
        * 可以通过配置项enable_url_middleware和url_length_limit来启用和设置最长URL长度，默认启用和最长长度为2083。
        * 在NewApp中加入crawler选项`pkg.WithUrlMiddleware()`
    * referrer: 110
        * 自动添加Referrer中间件，用于自动添加Referrer到请求中。
        * 可以根据referrer_policy配置项选择不同的Referrer策略，DefaultReferrerPolicy会加入请求来源，NoReferrerPolicy不加入请求来源
        * 配置 enable_referrer_middleware: true 是否开启自动添加referrer，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithReferrerMiddleware()`
    * cookie: 120
        * 自动添加Cookie中间件，用于自动添加之前请求返回的Cookie到后续请求中。
        * 可以通过配置项enable_cookie_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithCookieMiddleware()`
    * redirect: 130
        * 网址重定向中间件，用于处理网址重定向，默认支持301和302重定向。
        * 可以通过配置项enable_redirect_middleware和redirect_max_times来启用和设置最大重定向次数，默认启用和最大次数为1。
        * 在NewApp中加入crawler选项`pkg.WithRedirectMiddleware()`
    * chrome: 140
        * 模拟Chrome中间件，用于模拟Chrome浏览器。
        * 可以通过配置项enable_chrome_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithChromeMiddleware()`
    * httpAuth: 150
        * HTTP认证中间件，通过提供用户名（username）和密码（password）进行HTTP认证。
        * 需要在具体的请求中设置用户名和密码。可以通过配置项enable_http_auth_middleware来启用或禁用，默认禁用。
        * 在NewApp中加入crawler选项`pkg.WithHttpAuthMiddleware()`
    * compress: 160
        * 支持gzip/deflate解压缩中间件，用于处理响应的压缩编码。
        * 可以通过配置项enable_compress_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithCompressMiddleware()`
    * decode: 170
        * 中文解码中间件，支持对响应中的GBK、GB2312和Big5编码进行解码。
        * 可以通过配置项enable_decode_middleware来启用或禁用，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithDecodeMiddleware()`
    * device: 180
        * 修改请求设备信息中间件，用于修改请求的设备信息，包括请求头（header）和TLS信息。目前只支持User-Agent随机切换。
        * 需要设置设备范围（Platforms）和浏览器范围（Browsers）。
        * Platforms: Windows/Mac/Android/Iphone/Ipad/Linux
        * Browsers: Chrome/Edge/Safari/FireFox
        * 可以通过配置项enable_device_middleware来启用或禁用，默认禁用。
        * 在NewApp中加入crawler选项`pkg.WithDeviceMiddleware()`启用该中间件。
    * custom: 11
        * 自定义中间件
        * 在NewApp中加入crawler选项`pkg.WithCustomMiddleware(new(CustomMiddleware))`启用该中间件。
* Pipeline用于流式处理Item，如数据过滤、数据存储等。
  通过配置不同的Pipeline，您可以方便地处理Item并将结果保存到不同的目标，如控制台、文件、数据库或消息队列中。
  内置的Pipeline和自定义Pipeline使用默认的order值即可。
  如果需要改变默认的order值，需要在NewApp中加入crawler选项`pkg.WithPipeline(new(pipeline), order)`启用该Pipeline并应用该order值。
    * dump: 10
        * 用于在控制台打印Item的详细信息。
        * 您可以通过配置enable_dump_pipeline来控制是否启用该Pipeline，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithDumpPipeline()`启用该Pipeline。
    * file: 20
        * 用于下载文件并保存到Item中。
        * 您可以通过配置enable_file_pipeline来控制是否启用该Pipeline，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithFilePipeline()`启用该Pipeline。
    * image: 30
        * 用于下载图片并保存到Item中。
        * 您可以通过配置enable_image_pipeline来控制是否启用该Pipeline，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithImagePipeline()`启用该Pipeline。
    * filter: 200
        * 用于对Item进行过滤。
        * 它可用于去重请求，需要在中间件同时启用filter。
        * 默认情况下，Item只有在成功保存后才会进入去重队列。
        * 您可以通过配置enable_filter_pipeline来控制是否启用该Pipeline，默认启用。
        * 在NewApp中加入crawler选项`pkg.WithFilterPipeline()`启用该Pipeline。
    * csv: 101
        * 用于将结果保存到CSV文件中。
        * 需要在ItemCsv中设置`FileName`，指定保存的文件名称（不包含.csv扩展名）。
        * 您可以使用tag `column:""`来定义CSV文件的列名。
        * 您可以通过配置enable_csv_pipeline来控制是否启用该Pipeline，默认关闭。
        * 在NewApp中加入crawler选项`pkg.WithCsvPipeline()`启用该Pipeline。
    * jsonLines: 102
        * 用于将结果保存到JSON Lines文件中。
        * 需要在ItemJsonl中设置`FileName`，指定保存的文件名称（不包含.jsonl扩展名）。
        * 您可以使用tag `json:""`来定义JSON Lines文件的字段。
        * 您可以通过配置enable_json_lines_pipeline来控制是否启用该Pipeline，默认关闭。
        * 在NewApp中加入crawler选项`pkg.WithJsonLinesPipeline()`启用该Pipeline。
    * mongo: 103
        * 用于将结果保存到MongoDB中。
        * 需要在ItemMongo中设置`Collection`，指定保存的collection名称。
        * 您可以使用tag `bson:""`来定义MongoDB文档的字段。
        * 您可以通过配置enable_mongo_pipeline来控制是否启用该Pipeline，默认关闭。
        * 在NewApp中加入crawler选项`pkg.WithMongoPipeline()`启用该Pipeline。
    * mysql: 104
        * 用于将结果保存到MySQL中。
        * 需要在ItemMysql中设置`Table`，指定保存的表名。
        * 您可以使用tag `column:""`来定义MySQL表的列名。
        * 您可以通过配置enable_mysql_pipeline来控制是否启用该Pipeline，默认关闭。
        * 在NewApp中加入crawler选项`pkg.WithMysqlPipeline()`启用该Pipeline。
    * kafka: 105
        * 用于将结果保存到Kafka中。
        * 需要在ItemKafka中设置`Topic`，指定保存的主题名。
        * 您可以使用tag `json:""`来定义Kafka消息的字段。
        * 您可以通过配置enable_kafka_pipeline来控制是否启用该Pipeline，默认关闭。
        * 在NewApp中加入crawler选项`pkg.WithKafkaPipeline()`启用该Pipeline。
    * custom: 110
        * 自定义pipeline
        * 在NewApp中加入crawler选项`pkg.WithCustomPipeline(new(CustomPipeline))`启用该Pipeline。
* 信号（Signal）是一种机制，用于在运行时处理外部发出的操作指令。通过捕获和处理信号，您可以实现对爬虫的控制和管理
* 在配置文件中配置全局的请求参数，并在具体的请求中可以覆盖这些全局配置，可以提供更灵活和细粒度的请求定制
* 框架内置了多个解析模块。这些解析模块提供了不同的选择器和语法，以适应不同的数据提取需求。您可以根据具体的爬虫任务和数据结构，选择适合您的解析模块和语法，从网页响应中准确地提取所需的数据。
    * query选择器 go-query是一个处理query选择器的库 [go-query](https://github.com/lizongying/go-query)
        * 通过调用`response.Query()`方法，您可以使用query选择器语法来从HTML或XML响应中提取数据。
    * xpath选择器 go-xpath是一个可用于XPath选择的库 [go-xpath](https://github.com/lizongying/go-xpath)
        * 通过调用`response.Xpath()`方法，您可以使用XPath表达式来从HTML或XML响应中提取数据。
    * gjson gjson是一个用于处理JSON的库
        * 通过调用`response.Json()`方法，您可以使用gjson语法从JSON响应中提取数据。
    * re选择器 go-re是一个处理正则的库 [go-re](https://github.com/lizongying/go-re)
        * 通过调用`response.Re()`方法，您可以使用正则表达式从响应中提取数据。
* 代理。它可以帮助爬虫在请求网站时隐藏真实IP地址。
    * 自行搭建隧道代理：您可以使用 [go-proxy](https://github.com/lizongying/go-proxy)
      等工具来搭建隧道代理。这些代理工具可以提供随机切换的代理功能，对调用方无感知，方便使用。
      您可以在爬虫框架中集成这些代理工具，以便在爬虫请求时自动切换代理。
      这是一个随机切换的隧道代理，调用方无感知，方便使用。后期会加入一些其他的调用方式，比如维持原来的代理地址。
    * 其他调用方式：除了随机切换的代理方式，后期可以考虑加入其他的调用方式。
      例如，保持原来的代理地址不变，或者使用其他代理池工具进行代理IP的管理和调度。这样可以提供更多灵活性和选择性，以满足不同的代理需求。
* 要提高爬虫的性能，您可以考虑关闭一些未使用的中间件或Pipeline，以减少不必要的处理和资源消耗。以下是一些建议：
    * 检查中间件：审查已配置的中间件，并根据需要禁用不使用的中间件。您可以在配置文件中进行修改，或者在爬虫的入口方法中进行相应的配置更改。
    * 禁用不需要的Pipeline：检查已配置的Pipeline，并禁用不需要的Pipeline。
      例如，如果您不需要保存结果到MongoDB，可以禁用MongoPipeline。
    * 评估性能影响：在禁用中间件或Pipeline之前，请评估其对爬虫性能的实际影响。确保禁用的部分不会对功能产生负面影响。
    * 可以禁用的配置:
        * enable_stats_middleware: false
        * enable_dump_middleware: false
        * enable_filter_middleware: false
        * enable_file_middleware: false
        * enable_image_middleware: false
        * enable_http_middleware: false
        * enable_retry_middleware: false
        * enable_referrer_middleware: false
        * enable_http_auth_middleware: false
        * enable_cookie_middleware: false
        * enable_url_middleware: false
        * enable_compress_middleware: false
        * enable_decode_middleware: false
        * enable_redirect_middleware: false
        * enable_chrome_middleware: false
        * enable_device_middleware: false
        * enable_proxy_middleware: false
        * enable_robots_txt_middleware: false
        * enable_dump_pipeline: false
        * enable_file_pipeline: false
        * enable_image_pipeline: false
        * enable_filter_pipeline: false
        * enable_csv_pipeline: false
        * enable_json_lines_pipeline: false
        * enable_mongo_pipeline: false
        * enable_mysql_pipeline: false
        * enable_kafka_pipeline: false
* 文件下载
    * 如果您希望将文件保存到S3等对象存储中，需要进行相应的配置
    * Files下载
        * 在Item中设置Files请求：在Item中，您需要设置Files请求，即包含要下载的文件的请求列表。
          可以使用`item.SetFilesRequest([]pkg.Request{...})`
          方法设置请求列表。
        * Item.Data：您的Item.Data字段需要实现pkg.File的切片，用于保存下载文件的结果。
          该字段的名称必须是Files，如`type DataFile struct {Files []*media.File}`。
    * Images下载
        * 在Item中设置Images请求：在Item中，您需要设置Images请求，即包含要下载的图片的请求列表。
          可以使用item.SetImagesRequest([]pkg.Request{...})方法设置请求列表。
        * Item.Data：您的Item.Data字段需要实现pkg.Image的切片，用于保存下载图片的结果。
          该字段的名称必须是Images，如`type DataImage struct {Images []*media.Image}`。
* 爬虫结构
    * 建议按照每个网站（子网站）或者每个业务为一个spider。不必分的太细，也不必把所有的网站和业务都写在一个spider里
* 为了方便开发和调试，框架内置了本地devServer，在`-m dev`模式下会默认启用。
  通过使用本地devServer，您可以在开发和调试过程中更方便地模拟和观察网络请求和响应，以及处理自定义路由逻辑。
  这为开发者提供了一个便捷的工具，有助于快速定位和解决问题。
  您可以自定义路由（route），只需要实现`pkg.Route` 接口，并通过在Spider中调用`AddDevServerRoutes(...pkg.Route)`
  方法将其注册到devServer中。
    * 支持http和https，您可以通过设置`dev_server`选项来指定devServer的URL。
      `http://localhost:8081`表示使用HTTP协议，`https://localhost:8081`表示使用HTTPS协议。
    * 默认显示JA3指纹。JA3是一种用于TLS客户端指纹识别的算法，它可以显示与服务器建立连接时客户端使用的TLS版本和加密套件等信息。
    * 您可以使用tls工具来生成服务器的私钥和证书，以便在devServer中使用HTTPS。tls工具可以帮助您生成自签名的证书，用于本地开发和测试环境。
    * devServer内置了多种handler，这些handler提供了丰富的功能，可以模拟各种网络情景，帮助进行开发和调试。
      您可以根据需要选择合适的handler，并将其配置到devServer中，以模拟特定的网络响应和行为。
        * BadGatewayHandler 模拟返回502状态码
        * Big5Handler 模拟使用big5编码
        * CookieHandler 模拟返回cookie
        * DeflateHandler 模拟使用Deflate压缩
        * FileHandler 模拟输出文件
        * Gb2312Handler 模拟使用gb2312编码
        * Gb18030Handler 模拟使用gb18030编码
        * GbkHandler 模拟使用gbk编码
        * GzipHandler 模拟使用gzip压缩
        * HelloHandler 打印请求的header和body信息
        * HttpAuthHandler 模拟http-auth认证
        * InternalServerErrorHandler 模拟返回500状态码
        * OkHandler 模拟正常输出，返回200状态码
        * RateLimiterHandler 模拟速率限制，目前基于全部请求，不区分用户。可与HttpAuthHandler配合使用。
        * RedirectHandler 模拟302临时跳转，需要同时启用OkHandler
        * RobotsTxtHandler 返回robots.txt文件

### args

通过配置环境变量和启动参数，您可以更灵活地配置和控制爬虫的行为，包括选择配置文件、指定入口方法、传递额外参数以及设定启动模式。这样的设计可以提高爬虫的可配置性和可扩展性，使得爬虫框架更适应各种不同的应用场景。

* CRAWLER_CONFIG_FILE -c 配置文件路径，必须进行配置。
* CRAWLER_START_FUNC -f 入口方法名称，默认Test。
* CRAWLER_ARGS -a 额外的参数，该参数是非必须项。建议使用JSON字符串。参数会被入口方法调用。
* CRAWLER_MODE -m 启动模式，默认test。您可以根据需要使用不同的模式，如dev、prod等，以区分开发和生产环境。

### config

数据库配置：

* mongo_enable: 是否启用MongoDB。
* mongo.example.uri: MongoDB的URI。
* mongo.example.database: MongoDB的数据库名称。
* mysql_enable: 是否启用MySQL。
* mysql.example.uri: MySQL的URI。
* mysql.example.database: MySQL的数据库名称。
* redis_enable: 是否启用Redis。
* redis.example.addr: Redis的地址。
* redis.example.password: Redis的密码。
* redis.example.db: Redis的数据库。
* s3_enable: 是否启用S3对象存储（如COS、OSS、MinIO等）
* s3.example.endpoint: S3的地址
* s3.example.region: S3的区域。
* s3.example.id: S3的ID。
* s3.example.key: S3的密钥。
* s3.example.bucket: S3的桶名称。
* kafka_enable: 是否启用Kafka。
* kafka.example.uri: Kafka的URI。

日志配置：

* log.filename: 日志文件路径。可以使用"{name}"的方式替换成-ldflags的参数。
* log.long_file: 如果设置为true（默认），则记录完整文件路径。
* log.level: 日志级别，可选DEBUG/INFO/WARN/ERROR。

其他配置：

* proxy.example: 代理。
* request.concurrency: 请求并发数。
* request.interval: 请求间隔时间（毫秒）。默认1000毫秒（1秒）。
* request.timeout: 请求超时时间（秒）。默认60秒（1分钟）。
* request.ok_http_codes: 请求正常的HTTP状态码。
* request.retry_max_times: 请求重试的最大次数，默认10。
* request.http_proto: 请求的HTTP协议。默认`2.0`
* dev_server: 开发服务器的地址。默认`https://localhost:8081`
* enable_ja3: 是否打印JA3指纹。默认关闭。
* scheduler: 调度方式，默认memory（内存调度），可选值memory、redis、kafka。选择redis或kafka后可以实现集群调度。
* filter: 过滤方式，默认memory（内存过滤），可选值memory、redis。选择redis后可以实现集群过滤。

中间件和Pipeline配置

* enable_stats_middleware: 是否开启统计中间件，默认启用。
* enable_dump_middleware: 是否开启打印请求和响应中间件，默认启用。
* enable_filter_middleware: 是否开启过滤中间件，默认启用。
* enable_file_middleware: 是否开启文件处理中间件，默认启用。
* enable_image_middleware: 是否开启图片处理中间件，默认启用。
* enable_http_middleware: 是否开启HTTP请求中间件，默认启用。
* enable_retry_middleware: 是否开启请求重试中间件，默认启用。
* enable_referrer_middleware: 是否开启Referrer中间件，默认启用。
* referrer_policy: 设置Referrer策略，可选值为DefaultReferrerPolicy（默认）和NoReferrerPolicy。
* enable_http_auth_middleware: 是否开启HTTP认证中间件，默认关闭。
* enable_cookie_middleware:  是否开启Cookie中间件，默认启用。
* enable_url_middleware: 是否开启URL长度限制中间件，默认启用。
* url_length_limit: URL的最大长度限制，默认2083。
* enable_compress_middleware: 是否开启响应解压缩中间件（gzip、deflate），默认启用。
* enable_decode_middleware: 是否开启中文解码中间件（GBK、GB2312、Big5编码），默认启用。
* enable_redirect_middleware: 是否开启重定向中间件，默认启用。
* redirect_max_times: 重定向的最大次数，默认1。
* enable_chrome_middleware: 是否开启Chrome模拟中间件，默认启用。
* enable_device_middleware: 是否开启设备模拟中间件，默认关闭。
* enable_proxy_middleware: 是否开启代理中间件，默认启用。
* enable_robots_txt_middleware: 是否开启robots.txt支持中间件，默认关闭。
* enable_dump_pipeline: 是否开启打印Item Pipeline，默认启用。
* enable_file_pipeline: 是否开启文件下载Pipeline，默认启用。
* enable_image_pipeline: 是否开启图片下载Pipeline，默认启用。
* enable_filter_pipeline: 是否开启过滤Pipeline，默认启用。
* enable_csv_pipeline: 是否开启csv Pipeline，默认关闭。
* enable_json_lines_pipeline: 是否开启json lines Pipeline，默认关闭。
* enable_mongo_pipeline: 是否开启mongo Pipeline，默认关闭。
* enable_mysql_pipeline: 是否开启mysql Pipeline，默认关闭。
* enable_kafka_pipeline: 是否开启kafka Pipeline，默认关闭。

## Example

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/request"
)

type ExtraOk struct {
	Count int
}

type DataOk struct {
	Count int
}

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	
	item := pkg.ItemNone{
		ItemUnimplemented: pkg.ItemUnimplemented{
			Data: &DataOk{
				Count: extra.Count,
			},
		},
	}
	err = s.YieldItem(ctx, &item)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	if extra.Count > 0 {
		return
	}

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func (s *Spider) TestOk(ctx context.Context, _ string) (err error) {
	// mock server
	s.AddDevServerRoutes(devServer.NewOkHandler(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.SetName("test-ok")
	host, _ := spider.GetConfig().GetDevServer()
	spider.SetHost(host.String())

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}

```

### Test

```shell
go run cmd/testOkSpider/*.go -c example.yml -f TestOk -m dev

```

更多示例可以按照以下项目

[go-crawler-example](https://github.com/lizongying/go-crawler-example)

```shell
git clone github.com/lizongying/go-crawler-example
```

## TODO

* downloadtimeout
* AutoThrottle
* cron
* max request limit?
* multi-spider
* devServer独立拆分


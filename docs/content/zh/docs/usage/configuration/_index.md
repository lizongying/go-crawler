---
weight: 12
title: 配置
---

### 配置

在配置文件中的是全局配置，部分配置可以在爬虫中或者具体的请求中进行修改覆盖。
配置文件需要在启动时通过环境变量或参数指定，以下是配置参数：

* `env: dev` 环境。dev环境下不会写入数据库。
* `bot_name: crawler` 项目名

数据库配置：

* `mongo_enable:` 是否启用MongoDB。
* `mongo.example.uri:` MongoDB的URI。
* `mongo.example.database:` MongoDB的数据库名称。
* `mysql_enable:` 是否启用MySQL。
* `mysql.example.uri:` MySQL的URI。
* `mysql.example.database:` MySQL的数据库名称。
* `redis_enable:` 是否启用Redis。
* `redis.example.addr:` Redis的地址。
* `redis.example.password:` Redis的密码。
* `redis.example.db:` Redis的数据库。
* `sqlite.0.name:` sqlite名称，自定义
* `sqlite.0.path:` sqlite文件地址
* `store.0.name:` 存储名称，自定义
* `store.0.type:` 存储方式（如s3、cos、oss、minio、file等）
* `store.0.endpoint:` 对象存储的地址或者本地文件存储地址如“file://tmp/”
* `store.0.region:` 对象存储的区域。
* `store.0.id:` 对象存储的ID。
* `store.0.key:` 对象存储的密钥。
* `store.0.bucket:` 对象存储的桶名称。
* `kafka_enable:` 是否启用Kafka。
* `kafka.example.uri:` Kafka的URI。

日志配置：

* `log.filename:` 日志文件路径。可以使用"{name}"的方式替换成`-ldflags`的参数。
* `log.long_file:` 如果设置为true（默认），则记录完整文件路径。
* `log.level:` 日志级别，可选DEBUG/INFO/WARN/ERROR。

* `mock_server`: 模拟服务
    * `enable: false` 是否启用模拟服务
    * `host: https://localhost:8081` 模拟服务的地址。
    * `client_auth: 0` 客户端验证类型，0 不验证。

中间件和Pipeline配置

* `enable_stats_middleware:` 是否开启统计中间件，默认启用。
* `enable_dump_middleware:` 是否开启打印请求和响应中间件，默认启用。
* `enable_filter_middleware:` 是否开启过滤中间件，默认启用。
* `enable_file_middleware:` 是否开启文件处理中间件，默认启用。
* `enable_image_middleware:` 是否开启图片处理中间件，默认启用。
* `enable_http_middleware:` 是否开启HTTP请求中间件，默认启用。
* `enable_retry_middleware:` 是否开启请求重试中间件，默认启用。
* `enable_referrer_middleware:` 是否开启Referrer中间件，默认启用。
* `referrer_policy:` 设置Referrer策略，可选值为DefaultReferrerPolicy（默认）和NoReferrerPolicy。
* `enable_http_auth_middleware:` 是否开启HTTP认证中间件，默认关闭。
* `enable_cookie_middleware:`  是否开启Cookie中间件，默认启用。
* `enable_url_middleware:` 是否开启URL长度限制中间件，默认启用。
* `url_length_limit:` URL的最大长度限制，默认2083。
* `enable_compress_middleware:` 是否开启响应解压缩中间件（gzip、deflate），默认启用。
* `enable_decode_middleware:` 是否开启中文解码中间件（GBK、GB2312、Big5编码），默认启用。
* `enable_redirect_middleware:` 是否开启重定向中间件，默认启用。
* `redirect_max_times:` 重定向的最大次数，默认10。
* `enable_chrome_middleware:` 是否开启Chrome模拟中间件，默认启用。
* `enable_device_middleware:` 是否开启设备模拟中间件，默认关闭。
* `enable_proxy_middleware:` 是否开启代理中间件，默认启用。
* `enable_robots_txt_middleware:` 是否开启robots.txt支持中间件，默认关闭。
* `enable_record_error_middleware:` 是否开启record_error支持中间件，默认关闭。
* `enable_dump_pipeline:` 是否开启打印Item Pipeline，默认启用。
* `enable_none_pipeline:` 是否开启none Pipeline，默认关闭。
* `enable_file_pipeline:` 是否开启文件下载Pipeline，默认启用。
* `enable_image_pipeline:` 是否开启图片下载Pipeline，默认启用。
* `enable_filter_pipeline:` 是否开启过滤Pipeline，默认启用。
* `enable_csv_pipeline:` 是否开启csv Pipeline，默认关闭。
* `enable_json_lines_pipeline:` 是否开启json lines Pipeline，默认关闭。
* `enable_mongo_pipeline:` 是否开启mongo Pipeline，默认关闭。
* `enable_sqlite_pipeline:` 是否开启sqlite Pipeline，默认关闭。
* `enable_mysql_pipeline:` 是否开启mysql Pipeline，默认关闭。
* `enable_kafka_pipeline:` 是否开启kafka Pipeline，默认关闭。
* `enable_priority_queue:` 是否开启优先级队列，默认开启，目前只支持redis。

其他配置：

* proxy.example: 代理。
* request.concurrency: 请求并发数。
* request.interval: 请求间隔时间（毫秒）。默认1000毫秒（1秒）。
* request.timeout: 请求超时时间（秒）。默认60秒（1分钟）。
* request.ok_http_codes: 请求正常的HTTP状态码。
* request.retry_max_times: 请求重试的最大次数，默认10。
* request.http_proto: 请求的HTTP协议。默认`2.0`
* enable_ja3: 是否修改/打印JA3指纹。默认关闭。
* scheduler: 调度方式，默认memory（内存调度），可选值memory、redis、kafka。选择redis或kafka后可以实现集群调度。
* filter: 过滤方式，默认memory（内存过滤），可选值memory、redis。选择redis后可以实现集群过滤。
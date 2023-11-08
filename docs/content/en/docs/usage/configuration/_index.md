---
weight: 12
title: Configuration
---

## Configuration

In the configuration file, you can set global configurations that apply to all spiders. However, some configurations can
be modified and overridden in individual spiders or specific requests.
The configuration file needs to be specified at startup using environment variables or parameters. Here are the
configuration parameters:

* `env: dev`. In the `dev` environment, data will not be written to the database.
* `bot_name: crawler` Project Name

Database Configuration:

* `mongo_enable:` Whether to enable MongoDB.
* `mongo.example.uri:` MongoDB URI.
* `mongo.example.database:` MongoDB database name.
* `mysql_enable:` Whether to enable MySQL.
* `mysql.example.uri:` MySQL URI.
* `mysql.example.database:` MySQL database name.
* `redis_enable:` Whether to enable Redis.
* `redis.example.addr:` Redis address.
* `redis.example.password:` Redis password.
* `redis.example.db:` Redis database number.
* `sqlite.0.name:` sqlite name.
* `sqlite.0.path` sqlite file path.
* `store.0.name:` storage name.
* `store.0.type:` storage type (e.g., s3, cos, oss, minio, file, etc.).
* `store.0.endpoint:` S3 endpoint or file path like "file://tmp/".
* `store.0.region:` S3 region.
* `store.0.id:` S3 access ID.
* `store.0.key:` S3 access key.
* `store.0.bucket:` S3 bucket name.
* `kafka_enable:` Whether to enable Kafka.
* `kafka.example.uri:` Kafka URI.

Log Configuration:

* `log.filename:` Log file path. You can use "{name}" to replace it with a parameter from `-ldflags`.
* `log.long_file:` If set to true (default), it logs the full file path.
* `log.level:` Log level, options are DEBUG/INFO/WARN/ERROR.

* `mock_server`: Mock Server
    * `enable: false`: Whether to enable the mock Server.
    * `host: https://localhost:8081`: The address of the mock Server.
    * `client_auth: 0` Client authentication type, 0 means no authentication.

Middleware and Pipeline Configuration:

* `enable_stats_middleware:` Whether to enable the statistics middleware, enabled by default.
* `enable_dump_middleware:` Whether to enable the dump middleware for printing requests and responses, enabled by
  default.
* `enable_filter_middleware:` Whether to enable the filter middleware, enabled by default.
* `enable_file_middleware:` Whether to enable the file handling middleware, enabled by default.
* `enable_image_middleware:` Whether to enable the image handling middleware, enabled by default.
* `enable_http_middleware:` Whether to enable the HTTP request middleware, enabled by default.
* `enable_retry_middleware:` Whether to enable the request retry middleware, enabled by default.
* `enable_referrer_middleware:` Whether to enable the Referrer middleware, enabled by default.
* `referrer_policy:` Set the Referrer policy, options are DefaultReferrerPolicy (default) and NoReferrerPolicy.
* `enable_http_auth_middleware:` Whether to enable the HTTP authentication middleware, disabled by default.
* `enable_cookie_middleware:` Whether to enable the Cookie middleware, enabled by default.
* `enable_url_middleware:` Whether to enable the URL length limiting middleware, enabled by default.
* `url_length_limit:` Maximum length limit for URLs, default is 2083.
* `enable_compress_middleware:` Whether to enable the response decompression middleware (gzip, deflate), enabled by
  default.
* `enable_decode_middleware:` Whether to enable the Chinese decoding middleware (GBK, GB2312, Big5 encodings), enabled
  by default.
* `enable_redirect_middleware:` Whether to enable the redirect middleware, enabled by default.
* `redirect_max_times:` Maximum number of times to follow redirects, default is 10.
* `enable_chrome_middleware:` Whether to enable the Chrome simulation middleware, enabled by default.
* `enable_device_middleware:` Whether to enable the device simulation middleware, disabled by default.
* `enable_proxy_middleware:` Whether to enable the proxy middleware, enabled by default.
* `enable_robots_txt_middleware:` Whether to enable the robots.txt support middleware, disabled by default.
* `enable_record_error_middleware:` Whether to enable the record error support middleware, disabled by default.
* `enable_dump_pipeline:` Whether to enable the print item pipeline, enabled by default.
* `enable_none_pipeline:` Whether to enable the None pipeline, disabled by default。
* `enable_file_pipeline:` Whether to enable the file download pipeline, enabled by default.
* `enable_image_pipeline:` Whether to enable the image download pipeline, enabled by default.
* `enable_filter_pipeline:` Whether to enable the filter pipeline, enabled by default.
* `enable_csv_pipeline:` Whether to enable the CSV pipeline, disabled by default.
* `enable_json_lines_pipeline:` Whether to enable the JSON Lines pipeline, disabled by default.
* `enable_mongo_pipeline:` Whether to enable the MongoDB pipeline, disabled by default.
* `enable_sqlite_pipeline:` Whether to enable the Sqlite pipeline, disabled by default.
* `enable_mysql_pipeline:` Whether to enable the MySQL pipeline, disabled by default.
* `enable_kafka_pipeline:` Whether to enable the Kafka pipeline, disabled by default.
* `enable_priority_queue:` Whether to enable the priority queue, enabled by default, currently only supports Redis.

Other Configurations:

* `proxy.example`: Proxy configuration.
* `request.concurrency`: Number of concurrent requests.
* `request.interval`: Request interval time in milliseconds. Default is 1000 milliseconds (1 second).
* `request.timeout`: Request timeout in seconds. Default is 60 seconds (1 minute).
* `request.ok_http_codes`: Normal HTTP status codes for requests.
* `request.retry_max_times`: Maximum number of retries for requests. Default is 10.
* `request.http_proto`: HTTP protocol for requests. Default is `2.0`.
* `enable_ja3`: Whether to modify/print JA3 fingerprints. Default is disabled.
* `scheduler`: Scheduler method. Default is `memory` (memory-based scheduling). Options are `memory`, `redis`, `kafka`.
  Selecting `redis` or `kafka` enables cluster scheduling.
* `filter`: Filter method. Default is `memory` (memory-based filtering). Options are `memory`, `redis`.
  Selecting `redis` enables cluster filtering.
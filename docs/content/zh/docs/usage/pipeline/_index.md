---
weight: 5
title: 数据管道
---

## 数据管道

用于流式处理Item，如数据过滤、数据存储等。
通过配置不同的Pipeline，您可以方便地处理Item并将结果保存到不同的目标，如控制台、文件、数据库或消息队列中。
内置的Pipeline和自定义Pipeline使用默认的order值即可。
如果需要改变默认的order值，需要`spider.WithOptions(pkg.WithPipeline(new(pipeline), order)`启用该Pipeline并应用该order值。

* dump: 10
    * 用于在控制台打印Item的详细信息。
    * 您可以通过配置enable_dump_pipeline来控制是否启用该Pipeline，默认启用。
    * `spider.WithOptions(pkg.WithDumpPipeline()`
* file: 20
    * 用于下载文件并保存到Item中。
    * 您可以通过配置enable_file_pipeline来控制是否启用该Pipeline，默认启用。
    * `spider.WithOptions(pkg.WithFilePipeline()`
* image: 30
    * 用于下载图片并保存到Item中。
    * 您可以通过配置enable_image_pipeline来控制是否启用该Pipeline，默认启用。
    * `spider.WithOptions(pkg.WithImagePipeline()`
* filter: 200
    * 用于对Item进行过滤。
    * 它可用于去重请求，需要在中间件同时启用filter。
    * 默认情况下，Item只有在成功保存后才会进入去重队列。
    * 您可以通过配置enable_filter_pipeline来控制是否启用该Pipeline，默认启用。
    * `spider.WithOptions(pkg.WithFilterPipeline()`
* none: 101
    * item不做任何处理，但会认为结果已保存。
    * 您可以通过配置enable_none_pipeline来控制是否启用该Pipeline，默认关闭。
    * `spider.WithOptions(pkg.WithNonePipeline()`
* csv: 102
    * 用于将结果保存到CSV文件中。
    * 需要在ItemCsv中设置`FileName`，指定保存的文件名称（不包含.csv扩展名）。
    * 您可以使用tag `column:""`来定义CSV文件的列名。
    * 您可以通过配置enable_csv_pipeline来控制是否启用该Pipeline，默认关闭。
    * `spider.WithOptions(pkg.WithCsvPipeline()`
* jsonLines: 103
    * 用于将结果保存到JSON Lines文件中。
    * 需要在ItemJsonl中设置`FileName`，指定保存的文件名称（不包含.jsonl扩展名）。
    * 您可以使用tag `json:""`来定义JSON Lines文件的字段。
    * 您可以通过配置enable_json_lines_pipeline来控制是否启用该Pipeline，默认关闭。
    * `spider.WithOptions(pkg.WithJsonLinesPipeline()`
* mongo: 104
    * 用于将结果保存到MongoDB中。
    * 需要在ItemMongo中设置`Collection`，指定保存的collection名称。
    * 您可以使用tag `bson:""`来定义MongoDB文档的字段。
    * 您可以通过配置enable_mongo_pipeline来控制是否启用该Pipeline，默认关闭。
    * `spider.WithOptions(pkg.WithMongoPipeline()`
* sqlite: 105
    * 用于将结果保存到Sqlite中。
    * 需要在ItemSqlite中设置`Table`，指定保存的表名。
    * 您可以使用tag `column:""`来定义Sqlite表的列名。
    * 您可以通过配置enable_sqlite_pipeline来控制是否启用该Pipeline，默认关闭。
    * `spider.WithOptions(pkg.WithSqlitePipeline()`
* mysql: 106
    * 用于将结果保存到MySQL中。
    * 需要在ItemMysql中设置`Table`，指定保存的表名。
    * 您可以使用tag `column:""`来定义MySQL表的列名。
    * 您可以通过配置enable_mysql_pipeline来控制是否启用该Pipeline，默认关闭。
    * `spider.WithOptions(pkg.WithMysqlPipeline()`
* kafka: 107
    * 用于将结果保存到Kafka中。
    * 需要在ItemKafka中设置`Topic`，指定保存的主题名。
    * 您可以使用tag `json:""`来定义Kafka消息的字段。
    * 您可以通过配置enable_kafka_pipeline来控制是否启用该Pipeline，默认关闭。
    * `spider.WithOptions(pkg.WithKafkaPipeline()`
* custom: 110
    * 自定义pipeline
    * `spider.WithOptions(pkg.WithCustomPipeline(new(CustomPipeline))`
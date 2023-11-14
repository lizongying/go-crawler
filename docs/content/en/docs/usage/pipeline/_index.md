---
weight: 5
title: Pipeline
---

## Pipeline

Pipelines are used for stream processing of items, such as data filtering and data storage.By configuring different
pipelines, you can conveniently process items and save the results to different targets, such as the console, files,
databases, or message queues.

Built-in pipelines and custom pipelines use the default `order` value.If you need to change the default `order`
value, `spider.WithOptions(pkg.WithPipeline(new(pipeline), order)` to enable the
pipeline with the specified `order` value.

The following are the built-in pipelines with their respective `order` values:

* dump: 10
    * Used to print detailed information of items to the console.
    * You can control whether to enable this pipeline by configuring `enable_dump_pipeline`, which is enabled by
      default.
    * `spider.WithOptions(pkg.WithDumpPipeline()`
* file: 20
    * Used to download files and save them to items.
    * You can control whether to enable this pipeline by configuring `enable_file_pipeline`, which is enabled by
      default.
    * `spider.WithOptions(pkg.WithFilePipeline()`
* image: 30
    * Used to download images and save them to items.
    * You can control whether to enable this pipeline by configuring `enable_image_pipeline`, which is enabled by
      default.
    * `spider.WithOptions(pkg.WithImagePipeline()`
* filter: 200
    * Used for item filtering.
    * It can be used for deduplicating requests when filter middleware is enabled.
    * By default, items are only added to the deduplication queue after they are successfully saved.
    * You can control whether to enable this pipeline by configuring `enable_filter_pipeline`, which is enabled by
      default.
    * `spider.WithOptions(pkg.WithFilterPipeline()`
* none: 101
    * item is not processed in any way, but it is assumed that the result has been saved.
    * You can control whether to enable this pipeline by configuring `enable_none_pipeline`, which is enabled by
      default.
    * `spider.WithOptions(pkg.WithNonePipeline()`
* csv: 102
    * Used to save results to CSV files.
    * You need to set the `FileName` in the `ItemCsv`, which specifies the name of the file to be saved (without the
      .csv extension).
    * You can use the tag `column:""` to define the column names of the CSV file.
    * You can control whether to enable this pipeline by configuring `enable_csv_pipeline`, which is disabled by
      default.
    * `spider.WithOptions(pkg.WithCsvPipeline()`
* jsonLines: 103
    * Used to save results to JSON Lines files.
    * You need to set the `FileName` in the `ItemJsonl`, which specifies the name of the file to be saved (without
      the.jsonl extension).
    * You can use the tag `json:""` to define the fields of the JSON Lines file.
    * You can control whether to enable this pipeline by configuring `enable_json_lines_pipeline`, which is disabled
      by default.
    * `spider.WithOptions(pkg.WithJsonLinesPipeline()`
* mongo: 104
    * Used to save results to MongoDB.
    * You need to set the `Collection` in the `ItemMongo`, which specifies the name of the collection to be saved.
    * You can use the tag `bson:""` to define the fields of the MongoDB document.
    * You can control whether to enable this pipeline by configuring `enable_mongo_pipeline`, which is disabled by
      default.
    * `spider.WithOptions(pkg.WithMongoPipeline()`
* sqlite: 105
    * Used to save results to Sqlite.
    * You need to set the `Table` in the `ItemSqlite`, which specifies the name of the table to be saved.
    * You can use the tag `column:""` to define the column names of the Sqlite table.
    * You can control whether to enable this pipeline by configuring `enable_sqlite_pipeline`, which is disabled by
      default.
    * `spider.WithOptions(pkg.WithSqlitePipeline()`
* mysql: 106
    * Used to save results to MySQL.
    * You need to set the `Table` in the `ItemMysql`, which specifies the name of the table to be saved.
    * You can use the tag `column:""` to define the column names of the MySQL table.
    * You can control whether to enable this pipeline by configuring `enable_mysql_pipeline`, which is disabled by
      default.
    * `spider.WithOptions(pkg.WithMysqlPipeline()`
* kafka: 107
    * Used to save results to Kafka.
    * You need to set the `Topic` in the `ItemKafka`, which specifies the name of the topic to be saved.
    * You can use the tag `json:""` to define the fields of the Kafka message.
    * You can control whether to enable this pipeline by configuring `enable_kafka_pipeline`, which is disabled by
      default.
    * `spider.WithOptions(pkg.WithKafkaPipeline()`
* custom: 110
    * Custom pipeline.
    * `spider.WithOptions(pkg.WithCustomPipeline(new(CustomPipeline))`
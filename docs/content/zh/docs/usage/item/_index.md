---
weight: 3
title: 存储
---

### 存储

Item用于存储需要导出的数据和一些其他辅助信息。
框架里内置的Item涵盖了主要文件、数据库、消息队列等存储方式。
pkg.Item是一个接口，不能直接使用。pkg.ItemUnimplemented实现了pkg.Item的所有方法。
如果Item需要实现pkg.Item，可以组合pkg.ItemUnimplemented。 如：

```go
type ItemNone struct {
pkg.ItemUnimplemented
}

```

Item有一些通用方法：

* `Name() pkg.ItemName`
  获取Item的具体类型，如pkg.ItemNone、pkg.ItemCsv、pkg.ItemJsonl、pkg.ItemMongo、pkg.ItemSqlite、pkg.ItemMysql、pkg.ItemKafka等，用于Item反序列化到具体Item实现。
* `SetReferrer(string)` 设置referrer，可以用于记录请求的来源，一般不需要自己设置，由ReferrerMiddleware自动设置。
* `Referrer() string` 获取referrer。
* `SetUniqueKey(string)` 设置uniqueKey，可以用于过滤和其他唯一用途。
* `UniqueKey() string` 获取uniqueKey。
* `SetId(any)` 设置id，主要用于保存数据时的主键，和uniqueKey的一个区别是，id可能是在Response中产生，请求时不一定能获得。
* `Id() any` 获取id。
* `SetData(any)` 设置data，这是要存储的完整数据。为了规范化，强制要求指针类型。存储到不同的目标时，data需要设置不同的格式。
* `Data() any` 获取data。
* `DataJson() string` 获取data json字符串。
* `SetFilesRequest([]pkg.Request)` 设置文件的请求。这是一个slice，可以下载多个文件。
* `FilesRequest() []pkg.Request` 获取文件的请求。
* `SetFiles([]pkg.File)` 设置文件。下载后的文件通过这个方法设置到Item中。
* `Files() []pkg.File` 获取文件。
* `SetImagesRequest([]pkg.Request)` 设置图片的请求。这是一个slice，可以下载多个图片。
* `ImagesRequest() []pkg.Request` 获取图片的请求。
* `SetImages([]pkg.Image)` 设置图片。下载后的图片通过这个方法设置到Item中。
* `Images() []pkg.Image` 获取图片。

* 内置Item实现：框架提供了一些内置的Item实现，如pkg.ItemNone、pkg.ItemCsv、pkg.ItemJsonl、pkg.ItemMongo、pkg.ItemSqlite、
  pkg.ItemMysql、pkg.ItemKafka等。
  您可以根据需要，返回Item，并开启相应的Pipeline。如：

    ```go
    err = s.YieldItem(ctx, items.NewItemMongo(s.collection, true).
    SetUniqueKey(extra.Keyword).
    SetId(extra.Keyword).
    SetData(&data))

    ```

    ```go
    spider.WithOptions(pkg.WithMongoPipeline())
    ```

    * pkg.ItemNone 这个Item没有实现任何其他方法，主要用于调试。
        * `items.NewItemNone()`
    * pkg.ItemCsv 保存到csv中。
        * `items.NewItemCsv(filename string)`
        * filename：存储的文件名，不包括拓展名
    * pkg.ItemJsonl 保存到jsonl中。
        * `items.NewItemJsonl(filename string)`
        * filename：存储的文件名，不包括拓展名
    * pkg.ItemMongo 保存到mongo中。
        * `items.NewItemMongo(collection string, update bool)`
        * collection：mongo collection
        * update：如果数据已存在mongo中，是否更新
    * pkg.ItemSqlite 保存到Sqlite中。
        * `items.NewItemSqlite(table string, update bool)`
        * table：sqlite table
        * update：如果数据已存在mongo中，是否更新
    * pkg.ItemMysql 保存到mysql中。
        * `items.NewItemMysql(table string, update bool)`
        * table：mysql table
        * update：如果数据已存在mongo中，是否更新
    * pkg.ItemKafka 保存到kafka中。
        * `items.NewItemKafka(topic string)`
        * topic：kafka topic
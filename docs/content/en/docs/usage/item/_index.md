---
weight: 3
title: Item
---

## Item

The Item is used to store data that needs to be exported and some other auxiliary information.
The built-in Items in the framework cover major storage methods such as files, databases, and message queues.
`pkg.Item` is an interface and cannot be used directly. `pkg.ItemUnimplemented` implements all methods of `pkg.Item`.
If a custom Item needs to implement `pkg.Item`, it can be composed with `pkg.ItemUnimplemented`. For example:

```go
type ItemNone struct {
pkg.ItemUnimplemented
}
```

Item has some common methods:

* `Name() pkg.ItemName`: Get the specific type of the Item, such
  as `pkg.ItemNone`, `pkg.ItemCsv`, `pkg.ItemJsonl`, `pkg.ItemMongo`, `pkg.ItemMysql`, `pkg.ItemSqlite`, `pkg.ItemKafka`,
  etc.,
  which is used for deserializing the Item to the specific Item implementation.
* `SetReferrer(string)`: Set the referrer, which can be used to record the source of the request. Generally,
  there is no need to set it manually as it is automatically set by the `ReferrerMiddleware`.
* `Referrer() string`: Get the referrer.
* `SetUniqueKey(string)`: Set the unique key, which can be used for filtering and other unique purposes.
* `UniqueKey() string`: Get the unique key.
* `SetId(any)`: Set the ID, mainly used as the primary key when saving data. One difference from `UniqueKey` is
  that `Id` may be generated in the Response and may not be obtained when making the request.
* `Id() any`: Get the ID.
* `SetData(any)`: Set the data, which is the complete data to be stored. For standardization, it is required to
  be a pointer type. When storing data in different destinations, the data needs to be set in different formats.
* `Data() any`: Get the data.
* `DataJson() string`: Get the data json.
* `SetFilesRequest([]pkg.Request)`: Set the requests for downloading files. This is a slice and can be used to
  download multiple files.
* `FilesRequest() []pkg.Request`: Get the requests for downloading files.
* `SetFiles([]pkg.File)`: Set the downloaded files using this method.
* `Files() []pkg.File`: Get the downloaded files.
* `SetImagesRequest([]pkg.Request)`: Set the requests for downloading images. This is a slice and can be used to
  download multiple images.
* `ImagesRequest() []pkg.Request`: Get the requests for downloading images.
* `SetImages([]pkg.Image)`: Set the downloaded images using this method.
* `Images() []pkg.Image`: Get the downloaded images.
* Built-in Item Implementations: The framework provides some built-in Item implementations, such
  as `pkg.ItemNone`, `pkg.ItemCsv`, `pkg.ItemJsonl`, `pkg.ItemMongo`, `pkg.ItemMysql`, `pkg.ItemSqlite`, `pkg.ItemKafka`,
  etc.
  You can return an Item as needed and enable the corresponding Pipeline. For example:

    ```go
    err = s.YieldItem(ctx, items.NewItemMongo(s.collection, true).
    SetUniqueKey(extra.Keyword).
    SetId(extra.Keyword).
    SetData(&data))
    
    ```go
    spider.WithOptions(pkg.WithMongoPipeline())
    ```

    * pkg.ItemNone: This Item does not implement any other methods and is mainly used for debugging.
        * `items.NewItemNone()`
    * pkg.ItemCsv: Saves data to a CSV file.
        * `items.NewItemCsv(filename string)`
        * filename is the name of the file to be saved, without the extension.
    * pkg.ItemJsonl: Saves data to a JSONL file.
        * `items.NewItemJsonl(filename string)`
        * filename is the name of the file to be saved, without the extension.
    * pkg.ItemMongo: Saves data to MongoDB.
        * `items.NewItemMongo(collection string, update bool)`
        * collection is the MongoDB collection
        * update: whether to update the data if it already exists in MongoDB.
    * pkg.ItemSqlite: Saves data to Sqlite.
        * `items.NewItemSqlite(table string, update bool)`
        * table: the Sqlite table
        * update: whether to update the data if it already exists in Sqlite.
    * pkg.ItemMysql: Saves data to MySQL.
        * `items.NewItemMysql(table string, update bool)`
        * table: the MySQL table
        * update: whether to update the data if it already exists in MySQL.
    * pkg.ItemKafka: Sends data to Kafka.
        * `items.NewItemKafka(topic string)`
        * topic: the Kafka topic.
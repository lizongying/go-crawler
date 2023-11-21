package items

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
)

type ItemNone struct {
	pkg.ItemUnimplemented
}

func (i *ItemNone) MetaJson() string {
	return ""
}
func NewItemNone() pkg.Item {
	item := &ItemNone{}
	item.SetName(pkg.ItemNone)
	return item
}

type ItemMongoMeta struct {
	Collection string `json:"collection,omitempty"`
	Update     bool   `json:"update,omitempty"`
}
type ItemMongo struct {
	pkg.ItemUnimplemented
	ItemMongoMeta
}

func (i *ItemMongo) GetCollection() string {
	return i.Collection
}
func (i *ItemMongo) GetUpdate() bool {
	return i.Update
}
func (i *ItemMongo) MetaJson() string {
	bytes, err := json.Marshal(i.ItemMongoMeta)
	if err != nil {
		return ""
	}
	return string(bytes)
}
func NewItemMongo(collection string, update bool) pkg.Item {
	item := &ItemMongo{
		ItemMongoMeta: ItemMongoMeta{
			Collection: collection,
			Update:     update,
		},
	}
	item.SetName(pkg.ItemMongo)
	item.SetItem(item)
	return item
}

type ItemMysqlMeta struct {
	Table  string `json:"table,omitempty"`
	Update bool   `json:"update,omitempty"`
}
type ItemMysql struct {
	pkg.ItemUnimplemented
	ItemMysqlMeta
}

func (i *ItemMysql) GetTable() string {
	return i.Table
}
func (i *ItemMysql) GetUpdate() bool {
	return i.Update
}
func (i *ItemMysql) MetaJson() string {
	bytes, err := json.Marshal(i.ItemMysqlMeta)
	if err != nil {
		return ""
	}
	return string(bytes)
}
func NewItemMysql(table string, update bool) pkg.Item {
	item := &ItemMysql{
		ItemMysqlMeta: ItemMysqlMeta{
			Table:  table,
			Update: update,
		},
	}
	item.SetName(pkg.ItemMysql)
	item.SetItem(item)
	return item
}

type ItemSqliteMeta struct {
	Table  string `json:"table,omitempty"`
	Update bool   `json:"update,omitempty"`
}
type ItemSqlite struct {
	pkg.ItemUnimplemented
	ItemSqliteMeta
}

func (i *ItemSqlite) GetTable() string {
	return i.Table
}
func (i *ItemSqlite) GetUpdate() bool {
	return i.Update
}
func (i *ItemSqlite) MetaJson() string {
	bytes, err := json.Marshal(i.ItemSqliteMeta)
	if err != nil {
		return ""
	}
	return string(bytes)
}
func NewItemSqlite(table string, update bool) pkg.Item {
	item := &ItemSqlite{
		ItemSqliteMeta: ItemSqliteMeta{
			Table:  table,
			Update: update,
		},
	}
	item.SetName(pkg.ItemSqlite)
	item.SetItem(item)
	return item
}

type ItemKafkaMeta struct {
	Topic string `json:"topic,omitempty"`
}
type ItemKafka struct {
	pkg.ItemUnimplemented
	ItemKafkaMeta
}

func (i *ItemKafka) GetTopic() string {
	return i.Topic
}
func (i *ItemKafka) MetaJson() string {
	bytes, err := json.Marshal(i.ItemKafkaMeta)
	if err != nil {
		return ""
	}
	return string(bytes)
}
func NewItemKafka(topic string) pkg.Item {
	item := &ItemKafka{
		ItemKafkaMeta: ItemKafkaMeta{
			Topic: topic,
		},
	}
	item.SetName(pkg.ItemKafka)
	item.SetItem(item)
	return item
}

type ItemCsvMeta struct {
	FileName string `json:"file_name,omitempty"`
}
type ItemCsv struct {
	pkg.ItemUnimplemented
	ItemCsvMeta
}

func (i *ItemCsv) GetFileName() string {
	return i.FileName
}
func (i *ItemCsv) MetaJson() string {
	bytes, err := json.Marshal(i.ItemCsvMeta)
	if err != nil {
		return ""
	}
	return string(bytes)
}
func NewItemCsv(fileName string) pkg.Item {
	item := &ItemCsv{
		ItemCsvMeta: ItemCsvMeta{
			FileName: fileName,
		},
	}
	item.SetName(pkg.ItemCsv)
	item.SetItem(item)
	return item
}

type ItemJsonlMeta struct {
	FileName string `json:"file_name,omitempty"`
}
type ItemJsonl struct {
	pkg.ItemUnimplemented
	ItemJsonlMeta
}

func (i *ItemJsonl) GetFileName() string {
	return i.FileName
}
func (i *ItemJsonl) MetaJson() string {
	bytes, err := json.Marshal(i.ItemJsonlMeta)
	if err != nil {
		return ""
	}
	return string(bytes)
}
func NewItemJsonl(fileName string) pkg.Item {
	item := &ItemJsonl{
		ItemJsonlMeta: ItemJsonlMeta{
			FileName: fileName,
		},
	}
	item.SetName(pkg.ItemJsonl)
	item.SetItem(item)
	return item
}

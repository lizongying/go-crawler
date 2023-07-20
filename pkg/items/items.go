package items

import (
	"github.com/lizongying/go-crawler/pkg"
)

type ItemNone struct {
	pkg.ItemUnimplemented
}

func NewItemNone() pkg.Item {
	item := &ItemNone{}
	item.SetName(pkg.ItemNone)
	return item
}

type ItemMongo struct {
	pkg.ItemUnimplemented
	collection string
	update     bool
}

func (i *ItemMongo) GetCollection() string {
	return i.collection
}
func (i *ItemMongo) GetUpdate() bool {
	return i.update
}

func NewItemMongo(collection string, update bool) pkg.Item {
	item := &ItemMongo{
		collection: collection,
		update:     update,
	}
	item.SetName(pkg.ItemMongo)
	item.SetItem(item)
	return item
}

type ItemMysql struct {
	pkg.ItemUnimplemented
	table  string
	update bool
}

func (i *ItemMysql) GetTable() string {
	return i.table
}
func (i *ItemMysql) GetUpdate() bool {
	return i.update
}

func NewItemMysql(table string, update bool) pkg.Item {
	item := &ItemMysql{
		table:  table,
		update: update,
	}
	item.SetName(pkg.ItemMysql)
	item.SetItem(item)
	return item
}

type ItemKafka struct {
	pkg.ItemUnimplemented
	topic string
}

func (i *ItemKafka) GetTopic() string {
	return i.topic
}

func NewItemKafka(topic string) pkg.Item {
	item := &ItemKafka{
		topic: topic,
	}
	item.SetName(pkg.ItemKafka)
	item.SetItem(item)
	return item
}

type ItemCsv struct {
	pkg.ItemUnimplemented
	fileName string
}

func (i *ItemCsv) GetFileName() string {
	return i.fileName
}

func NewItemCsv(fileName string) pkg.Item {
	item := &ItemCsv{
		fileName: fileName,
	}
	item.SetName(pkg.ItemCsv)
	item.SetItem(item)
	return item
}

type ItemJsonl struct {
	pkg.ItemUnimplemented
	fileName string
}

func (i *ItemJsonl) GetFileName() string {
	return i.fileName
}

func NewItemJsonl(fileName string) pkg.Item {
	item := &ItemJsonl{
		fileName: fileName,
	}
	item.SetName(pkg.ItemJsonl)
	item.SetItem(item)
	return item
}

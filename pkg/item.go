package pkg

type Item interface {
	GetUniqueKey() string
	GetId() any
	GetData() any
}

type ItemUnimplemented struct {
	UniqueKey string
	Id        any
	Data      any
}

func (i *ItemUnimplemented) GetUniqueKey() string {
	return i.UniqueKey
}

func (i *ItemUnimplemented) GetId() any {
	return i.Id
}

func (i *ItemUnimplemented) GetData() any {
	return i.Data
}

type ItemMongo struct {
	ItemUnimplemented
	Update     bool
	Collection string
}

type ItemMysql struct {
	ItemUnimplemented
	Update bool
	Table  string
}

type ItemKafka struct {
	ItemUnimplemented
	Topic string
}

type ItemCsv struct {
	ItemUnimplemented
	FileName string
}

type ItemJsonl struct {
	ItemUnimplemented
	FileName string
}

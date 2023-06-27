package pkg

type Item interface {
	GetUniqueKey() string
	GetId() any
	GetData() any
	SetReferer(string)
	GetReferer() string
}

type ItemUnimplemented struct {
	referer   string
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
func (i *ItemUnimplemented) SetReferer(referer string) {
	i.referer = referer
}
func (i *ItemUnimplemented) GetReferer() string {
	return i.referer
}

type ItemNone struct {
	ItemUnimplemented
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

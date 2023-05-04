package pkg

type Item interface {
	GetUniqueKey() string
	GetId() any
	GetData() any
}

type ItemUnimplemented struct {
}

func (i *ItemUnimplemented) GetUniqueKey() (uniqueKey string) {
	return
}

func (i *ItemUnimplemented) GetId() (id any) {
	return
}

func (i *ItemUnimplemented) GetData() (data any) {
	return
}

type ItemMongo struct {
	ItemUnimplemented
	UniqueKey  string
	Id         any
	Update     bool
	Data       any
	Collection string
}

func (i *ItemMongo) GetUniqueKey() string {
	return i.UniqueKey
}

func (i *ItemMongo) GetId() any {
	return i.Id
}

func (i *ItemMongo) GetData() any {
	return i.Data
}

type ItemCsv struct {
	ItemUnimplemented
	UniqueKey string
	Id        any
	Data      any
	FileName  string
}

func (i *ItemCsv) GetUniqueKey() string {
	return i.UniqueKey
}

func (i *ItemCsv) GetId() any {
	return i.Id
}

func (i *ItemCsv) GetData() any {
	return i.Data
}

type ItemJsonl struct {
	ItemUnimplemented
	UniqueKey string
	Id        any
	Data      any
	FileName  string
}

func (i *ItemJsonl) GetUniqueKey() string {
	return i.UniqueKey
}

func (i *ItemJsonl) GetId() any {
	return i.Id
}

func (i *ItemJsonl) GetData() any {
	return i.Data
}

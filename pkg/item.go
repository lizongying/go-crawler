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

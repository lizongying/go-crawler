package items

import (
	"github.com/lizongying/go-crawler/pkg"
)

const Custom pkg.ItemName = "custom"

type ItemCustom struct {
	pkg.ItemUnimplemented
}

func NewItemCustom() pkg.Item {
	item := &ItemCustom{}
	item.SetName(Custom)
	return item
}

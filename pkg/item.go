package pkg

import (
	"reflect"
	"strings"
)

type Item interface {
	GetUniqueKey() string
	GetId() any
	GetData() any
	SetReferer(string)
	GetReferer() string
	SetImagesRequest([]*Request)
	GetImagesRequest() []*Request
	SetImages([]Image)
	GetImages() []Image
}

type ItemUnimplemented struct {
	images    []*Request
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
func (i *ItemUnimplemented) SetImagesRequest(images []*Request) {
	for _, v := range images {
		v.SetImage(true)
		i.images = append(i.images, v)
	}
}
func (i *ItemUnimplemented) GetImagesRequest() []*Request {
	return i.images
}

func (i *ItemUnimplemented) SetImages(images []Image) {
	if len(images) == 0 {
		return
	}

	t := reflect.TypeOf(i.Data).Elem()
	v := reflect.ValueOf(i.Data).Elem()
	l := t.NumField()

	for idx := 0; idx < l; idx++ {
		if t.Field(idx).Tag.Get("images") != "" {
			names := strings.Split(t.Field(idx).Tag.Get("images"), ",")
			if t.Field(idx).Type.Kind() != reflect.Slice {
				continue
			}

			elemType := t.Field(idx).Type.Elem()

			for _, image := range images {
				vv := reflect.New(elemType.Elem())
				for _, name := range names {
					f := vv.Elem().FieldByName(name)
					if !f.IsValid() || !f.CanSet() {
						continue
					}
					switch name {
					case "Name":
						f.SetString(image.GetName())
					case "Extension":
						f.SetString(image.GetExtension())
					case "Width":
						f.SetInt(int64(image.GetWidth()))
					case "Height":
						f.SetInt(int64(image.GetHeight()))
					}
				}
				v.Field(idx).Set(reflect.Append(v.Field(idx), vv))
			}
			break
		}
	}
}
func (i *ItemUnimplemented) GetImages() []Image {
	t := reflect.TypeOf(i.Data).Elem()
	v := reflect.ValueOf(i.Data).Elem()
	l := t.NumField()
	for idx := 0; idx < l; idx++ {
		if t.Field(idx).Tag.Get("images") != "" {
			return v.Field(idx).Interface().([]Image)
		}
	}
	return nil
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

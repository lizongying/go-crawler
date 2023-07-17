package pkg

import (
	"reflect"
)

type ItemName string

const (
	ItemNone  ItemName = ""
	ItemMongo ItemName = "mongo"
	ItemKafka ItemName = "kafka"
	ItemMysql ItemName = "mysql"
	ItemCsv   ItemName = "csv"
	ItemJsonl ItemName = "jsonl"
)

type Item interface {
	GetName() ItemName
	SetUniqueKey(string) Item
	GetUniqueKey() string
	SetId(any) Item
	GetId() any
	SetData(any) Item
	GetData() any
	SetReferrer(string) Item
	GetReferrer() string
	SetFilesRequest([]Request) Item
	GetFilesRequest() []Request
	SetFiles([]File) Item
	GetFiles() []File
	SetImagesRequest([]Request) Item
	GetImagesRequest() []Request
	SetImages([]Image) Item
	GetImages() []Image
}

type ItemUnimplemented struct {
	name      ItemName
	files     []Request
	images    []Request
	referrer  string
	uniqueKey string
	id        any
	data      any
}

func (i *ItemUnimplemented) SetName(name ItemName) Item {
	i.name = name
	return i
}
func (i *ItemUnimplemented) GetName() ItemName {
	return i.name
}
func (i *ItemUnimplemented) SetUniqueKey(uniqueKey string) Item {
	i.uniqueKey = uniqueKey
	return i
}
func (i *ItemUnimplemented) GetUniqueKey() string {
	return i.uniqueKey
}
func (i *ItemUnimplemented) SetId(id any) Item {
	i.id = id
	return i
}
func (i *ItemUnimplemented) GetId() any {
	return i.id
}
func (i *ItemUnimplemented) SetData(data any) Item {
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Ptr || dataValue.IsNil() {
		return i
	}
	i.data = data
	return i
}
func (i *ItemUnimplemented) GetData() any {
	return i.data
}
func (i *ItemUnimplemented) SetReferrer(referrer string) Item {
	i.referrer = referrer
	return i
}
func (i *ItemUnimplemented) GetReferrer() string {
	return i.referrer
}
func (i *ItemUnimplemented) SetFilesRequest(files []Request) Item {
	for _, v := range files {
		v.SetFile(true)
		i.files = append(i.files, v)
	}
	return i
}
func (i *ItemUnimplemented) GetFilesRequest() []Request {
	return i.files
}
func (i *ItemUnimplemented) SetImagesRequest(images []Request) Item {
	for _, v := range images {
		v.SetImage(true)
		i.images = append(i.images, v)
	}
	return i
}
func (i *ItemUnimplemented) GetImagesRequest() []Request {
	return i.images
}
func (i *ItemUnimplemented) SetFiles(files []File) Item {
	if len(files) == 0 {
		return i
	}

	f := reflect.ValueOf(i.data).Elem().FieldByName("Files")
	if f.IsValid() && f.Type().Kind() == reflect.Slice {
		for _, file := range files {
			f.Set(reflect.Append(f, reflect.ValueOf(file)))
		}
	}
	return i
}
func (i *ItemUnimplemented) GetFiles() []File {
	f := reflect.ValueOf(i.data).Elem().FieldByName("Files")
	if f.IsValid() && f.Type().Kind() == reflect.Slice {
		return f.Interface().([]File)
	}

	return nil
}
func (i *ItemUnimplemented) SetImages(images []Image) Item {
	if len(images) == 0 {
		return i
	}

	img := reflect.ValueOf(i.data).Elem().FieldByName("Images")
	if img.IsValid() && img.Type().Kind() == reflect.Slice {
		for _, image := range images {
			img.Set(reflect.Append(img, reflect.ValueOf(image)))
		}
	}
	return i
}
func (i *ItemUnimplemented) GetImages() []Image {
	img := reflect.ValueOf(i.data).Elem().FieldByName("Images")
	if img.IsValid() && img.Type().Kind() == reflect.Slice {
		return img.Interface().([]Image)
	}

	return nil
}

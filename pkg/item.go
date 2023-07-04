package pkg

import (
	"reflect"
)

type Item interface {
	GetUniqueKey() string
	GetId() any
	GetData() any
	SetReferer(string)
	GetReferer() string
	SetFilesRequest([]*Request)
	GetFilesRequest() []*Request
	SetFiles([]File)
	GetFiles() []File
	SetImagesRequest([]*Request)
	GetImagesRequest() []*Request
	SetImages([]Image)
	GetImages() []Image
}

type ItemUnimplemented struct {
	files     []*Request
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
func (i *ItemUnimplemented) SetFilesRequest(files []*Request) {
	for _, v := range files {
		v.SetFile(true)
		i.files = append(i.files, v)
	}
}
func (i *ItemUnimplemented) GetFilesRequest() []*Request {
	return i.files
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
func (i *ItemUnimplemented) SetFiles(files []File) {
	if len(files) == 0 {
		return
	}

	f := reflect.ValueOf(i.Data).Elem().FieldByName("Files")
	if f.IsValid() && f.Type().Kind() == reflect.Slice {
		for _, file := range files {
			f.Set(reflect.Append(f, reflect.ValueOf(file)))
		}
	}
}
func (i *ItemUnimplemented) GetFiles() []File {
	f := reflect.ValueOf(i.Data).Elem().FieldByName("Files")
	if f.IsValid() && f.Type().Kind() == reflect.Slice {
		return f.Interface().([]File)
	}

	return nil
}
func (i *ItemUnimplemented) SetImages(images []Image) {
	if len(images) == 0 {
		return
	}

	img := reflect.ValueOf(i.Data).Elem().FieldByName("Images")
	if img.IsValid() && img.Type().Kind() == reflect.Slice {
		for _, image := range images {
			img.Set(reflect.Append(img, reflect.ValueOf(image)))
		}
	}
}
func (i *ItemUnimplemented) GetImages() []Image {
	img := reflect.ValueOf(i.Data).Elem().FieldByName("Images")
	if img.IsValid() && img.Type().Kind() == reflect.Slice {
		return img.Interface().([]Image)
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

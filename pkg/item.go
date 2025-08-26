package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ItemName string

const (
	ItemNone   ItemName = ""
	ItemMongo  ItemName = "mongo"
	ItemKafka  ItemName = "kafka"
	ItemMysql  ItemName = "mysql"
	ItemCsv    ItemName = "csv"
	ItemJsonl  ItemName = "jsonl"
	ItemSqlite ItemName = "sqlite"
)

type Item interface {
	GetContext() Context
	WithContext(Context) Item
	GetItem() any
	Name() ItemName
	SetUniqueKey(string) Item
	UniqueKey() string
	SetId(any) Item
	Id() any
	SetData(any) Item
	Data() any
	DataJson() string
	MetaJson() string
	SetReferrer(string) Item
	Referrer() string
	SetFilesRequest([]Request) Item
	FilesRequest() []Request
	SetFiles([]File) Item
	Files() []File
	SetImagesRequest([]Request) Item
	ImagesRequest() []Request
	SetImages([]Image) Item
	Images() []Image
	Yield() error
	MustYield()
	UnsafeYield()
}

type ItemUnimplemented struct {
	context   Context
	item      any
	name      ItemName
	files     []Request
	images    []Request
	referrer  string
	uniqueKey string
	id        any
	data      any
}

func (i *ItemUnimplemented) GetContext() Context {
	return i.context
}
func (i *ItemUnimplemented) WithContext(ctx Context) Item {
	i.context = ctx
	return i
}
func (i *ItemUnimplemented) SetItem(item any) Item {
	i.item = item
	return i
}
func (i *ItemUnimplemented) GetItem() any {
	return i.item
}
func (i *ItemUnimplemented) SetName(name ItemName) Item {
	i.name = name
	return i
}
func (i *ItemUnimplemented) Name() ItemName {
	return i.name
}
func (i *ItemUnimplemented) SetUniqueKey(uniqueKey string) Item {
	i.uniqueKey = uniqueKey
	return i
}
func (i *ItemUnimplemented) UniqueKey() string {
	return i.uniqueKey
}
func (i *ItemUnimplemented) SetId(id any) Item {
	i.id = id
	return i
}
func (i *ItemUnimplemented) Id() any {
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
func (i *ItemUnimplemented) Data() any {
	return i.data
}
func (i *ItemUnimplemented) DataJson() string {
	bytes, err := json.Marshal(i.data)
	if err != nil {
		return ""
	}
	return string(bytes)
}
func (i *ItemUnimplemented) MetaJson() string {
	return ""
}
func (i *ItemUnimplemented) SetReferrer(referrer string) Item {
	i.referrer = referrer
	return i
}
func (i *ItemUnimplemented) Referrer() string {
	return i.referrer
}
func (i *ItemUnimplemented) SetFilesRequest(files []Request) Item {
	for _, v := range files {
		v.AsFile(true)
		i.files = append(i.files, v)
	}
	return i
}
func (i *ItemUnimplemented) FilesRequest() []Request {
	return i.files
}
func (i *ItemUnimplemented) SetImagesRequest(images []Request) Item {
	for _, v := range images {
		v.AsImage(true)
		i.images = append(i.images, v)
	}
	return i
}
func (i *ItemUnimplemented) ImagesRequest() []Request {
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
func (i *ItemUnimplemented) Files() []File {
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
func (i *ItemUnimplemented) Images() []Image {
	img := reflect.ValueOf(i.data).Elem().FieldByName("Images")
	if img.IsValid() && img.Type().Kind() == reflect.Slice {
		return img.Interface().([]Image)
	}

	return nil
}
func (i *ItemUnimplemented) Yield() error {
	s := i.context.GetSpider().GetSpider()
	return s.YieldItem(i.context, i)
}
func (i *ItemUnimplemented) MustYield() {
	s := i.context.GetSpider().GetSpider()
	if err := s.YieldItem(i.context, i); err != nil {
		panic(fmt.Errorf("%w: %v", ErrYieldItemFailed, err))
	}
}
func (i *ItemUnimplemented) UnsafeYield() {
	s := i.context.GetSpider().GetSpider()
	if err := s.YieldItem(i.context, i); err != nil {
		s.Logger().Error(err)
	}
}

type ItemStatus uint8

const (
	ItemStatusUnknown = iota
	ItemStatusPending
	ItemStatusRunning
	ItemStatusSuccess
	ItemStatusFailure
)

func (s ItemStatus) String() string {
	switch s {
	case ItemStatusPending:
		return "pending"
	case ItemStatusRunning:
		return "running"
	case ItemStatusSuccess:
		return "success"
	case ItemStatusFailure:
		return "failure"
	default:
		return "unknown"
	}
}

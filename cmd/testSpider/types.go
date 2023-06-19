package main

import (
	"github.com/lizongying/go-crawler/pkg/image"
)

type ExtraNoLimit struct {
	Count int
}

type DataNoLimit struct {
	Count int
}

type ExtraTest struct {
	*image.Image
}

type DataTest struct {
	*image.Image
}

type ExtraOk struct {
	Count int
}

type DataOk struct {
	Id    string `bson:"_id" json:"id"`
	Count int    `bson:"count" json:"count"`
}
type DataMysql struct {
	Id int    `column:"id"`
	A  uint   `column:"a"`
	B  uint32 `column:"b"`
	C  string `column:"c"`
	D  string `column:"d"`
}

type ExtraHttpAuth struct {
	Count int
}

type ExtraCookie struct {
	Count int
}

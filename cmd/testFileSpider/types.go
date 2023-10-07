package main

import "github.com/lizongying/go-crawler/pkg/media"

type DataImage struct {
	Images []*media.Image `json:"images"`
	DataOk
}

type DataOk struct {
	Id    string `bson:"_id" json:"id"`
	Count int    `bson:"count" json:"count"`
}

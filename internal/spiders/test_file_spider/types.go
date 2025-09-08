package test_file_spider

import "github.com/lizongying/go-crawler/pkg/media"

type DataImage struct {
	Images []*media.Image `json:"images1" image:"url,name,ext,width,height"`
}

type DataFile struct {
	Files []*media.File `json:"files1" file:"url,name,ext"`
}

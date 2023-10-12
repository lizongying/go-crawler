package test_file_spider

import "github.com/lizongying/go-crawler/pkg/media"

type DataImage struct {
	Images []*media.Image `json:"images" field:"url,name,ext,width,height"`
}

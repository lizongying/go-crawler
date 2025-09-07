package middlewares

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/media"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type FileMiddleware struct {
	pkg.UnimplementedMiddleware
	logger         pkg.Logger
	store          pkg.Store
	contentTypeMap map[string][]string
}

func (m *FileMiddleware) ProcessResponse(c pkg.Context, response pkg.Response) (err error) {
	if len(response.BodyBytes()) == 0 {
		m.logger.Debug("BodyBytes empty")
		return
	}

	isFile := response.IsFile()
	if isFile {
		options := response.FileOptions()
		i := new(media.File)
		i.SetUrl(response.Url())
		name := utils.StrMd5(response.Url())
		if options.Name {
			i.SetName(name)
		}
		ext := ""
		if e, ok := m.contentTypeMap[response.GetHeader("Content-Type")]; ok {
			ext = e[0]
			if options.Ext {
				i.SetExt(ext)
			}
		}

		key := name
		if ext != "" {
			key = fmt.Sprintf("%s.%s", name, ext)
		}
		storePath := ""
		storePath, err = m.store.Save("", key, response.BodyBytes())
		if err != nil {
			m.logger.Error(err)
			return
		}
		i.SetStorePath(storePath)

		response.SetFiles(append(response.Files(), i))

		stats, ok := c.GetTask().GetStats().(pkg.StatsWithFile)
		if ok {
			stats.IncFileTotal()
		}
	}

	return
}

func (m *FileMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(FileMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	crawler := spider.GetCrawler()
	m.logger = spider.GetLogger()
	m.store, _ = crawler.GetStore(spider.GetConfig().GetStorage())

	// https://developer.mozilla.org/zh-CN/docs/Web/Media/Formats/Image_types
	// https://developer.mozilla.org/zh-CN/docs/Web/Media/Formats/Containers
	m.contentTypeMap = map[string][]string{
		"image/apng":      {"apng"},
		"image/avif":      {"avif"},
		"image/bmp":       {"bmp"},
		"image/gif":       {"gif"},
		"image/x-icon":    {"icon", "cur"},
		"image/jpeg":      {"jpg", "jpeg", "jfif", "pjpeg", "pjp"},
		"image/png":       {"png"},
		"image/svg+xml":   {"svg"},
		"image/tiff":      {"tif", "tiff"},
		"image/webp":      {"webp"},
		"audio/3gpp":      {"3gp"},
		"video/3gpp":      {"3gp"},
		"audio/aac":       {"aac"},
		"audio/flac":      {"flac"},
		"audio/mpeg":      {"mpg", "mpeg"},
		"video/mpeg":      {"mpg", "mpeg"},
		"audio/mp3":       {"mp3"},
		"audio/mp4":       {"mp4", "m4a"},
		"video/mp4":       {"mp4", "m4v", "m4p"},
		"audio/ogg":       {"oga", "ogg"},
		"video/ogg":       {"ogv", "ogg"},
		"video/quicktime": {"mov"},
		"audio/wav":       {"wav"},
		"audio/webm":      {"webm"},
		"video/webm":      {"webm"},
	}

	return m
}

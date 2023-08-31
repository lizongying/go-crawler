package middlewares

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/media"
	"github.com/lizongying/go-crawler/pkg/utils"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type FileMiddleware struct {
	pkg.UnimplementedMiddleware
	logger         pkg.Logger
	s3             *s3.Client
	bucketName     string
	key            string
	ContentTypeMap map[string]string
}

func (m *FileMiddleware) ProcessResponse(ctx pkg.Context, response pkg.Response) (err error) {
	spider := m.GetSpider()
	if len(response.BodyBytes()) == 0 {
		err = errors.New("BodyBytes empty")
		m.logger.Error(err)
		return
	}

	isFile := response.File()
	if isFile {
		i := new(media.File)
		i.SetName(utils.StrMd5(response.GetUrl()))
		ext := ""
		if e, ok := m.ContentTypeMap[response.GetHeader("Content-Type")]; ok {
			ext = e
		}

		if m.s3 != nil {
			key := fmt.Sprintf("%s.%s", utils.StrMd5(response.GetUrl()), ext)
			storePath := fmt.Sprintf("s3://%s/%s", m.bucketName, key)
			uploadParams := &s3.PutObjectInput{
				Bucket: &m.bucketName,
				Key:    &key,
				Body:   bytes.NewReader(response.BodyBytes()),
			}

			// Upload the file
			_, e := m.s3.PutObject(context.TODO(), uploadParams)
			if e != nil {
				err = e
				m.logger.Error(err)
				return
			}

			i.SetStorePath(storePath)
		}

		response.SetFiles(append(response.Files(), i))

		stats, ok := spider.GetStats().(pkg.StatsWithFile)
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
	m.logger = spider.GetLogger()
	crawler := spider.GetCrawler()
	m.s3 = crawler.GetS3()
	m.bucketName = crawler.GetConfig().GetBotName()
	m.ContentTypeMap = map[string]string{
		"image/jpeg": "jpeg",
		"image/png":  "png",
		"image/gif":  "gif",
	}
	return m
}

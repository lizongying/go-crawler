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
	logger     pkg.Logger
	s3         *s3.Client
	stats      pkg.StatsWithImage
	bucketName string
	key        string
}

func (m *FileMiddleware) ProcessResponse(ctx context.Context, response *pkg.Response) (err error) {
	if len(response.BodyBytes) == 0 {
		err = errors.New("BodyBytes empty")
		m.logger.Error(err)
		return
	}

	isImage := response.Request.GetImage()
	if isImage {
		i := new(media.File)
		i.SetName(utils.StrMd5(response.Request.URL.String()))

		if m.s3 != nil {
			key := fmt.Sprintf("%s.%s", utils.StrMd5(response.Request.URL.String()), "")
			storePath := fmt.Sprintf("s3://%s/%s", m.bucketName, key)
			uploadParams := &s3.PutObjectInput{
				Bucket: &m.bucketName,
				Key:    &key,
				Body:   bytes.NewReader(response.BodyBytes),
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

		response.Files = append(response.Files, i)
		if m.stats != nil {
			m.stats.IncImageTotal()
		}
	}

	return
}

func (m *FileMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(FileMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	m.s3 = crawler.GetS3()
	m.stats, _ = crawler.GetStats().(pkg.StatsWithImage)
	m.bucketName = "crawler"
	return m
}

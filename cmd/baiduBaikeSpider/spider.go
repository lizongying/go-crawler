package main

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/url"
	"regexp"
)

type Spider struct {
	pkg.Spider
	logger               pkg.Logger
	collectionBaiduBaike string
	reItem               *regexp.Regexp
}

func (s *Spider) ParseDetail(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraDetail
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))

	content := response.BodyText()
	if content == "" {
		return
	}
	data := DataWord{
		Id:      extra.Keyword,
		Keyword: extra.Keyword,
		Content: content,
	}
	err = s.YieldItem(ctx, items.NewItemMongo(s.collectionBaiduBaike, true).
		SetUniqueKey(extra.Keyword).
		SetId(extra.Keyword).
		SetData(&data))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) ParseIndex(ctx context.Context, response pkg.Response) (err error) {
	links := response.AllLink()
	if err != nil {
		s.logger.Error(err)
		return
	}

	for _, v := range links {
		r := s.reItem.FindStringSubmatch(v.Path)
		if len(r) == 2 {
			decodedString, e := url.QueryUnescape(r[1])
			if e != nil {
				continue
			}
			err = s.YieldRequest(ctx, request.NewRequest().
				SetExtra(&ExtraDetail{
					Keyword: decodedString,
				}).
				SetCallBack(s.ParseDetail))
			if err != nil {
				s.logger.Error(err)
				continue
			}
		} else {
			err = s.YieldRequest(ctx, request.NewRequest().
				SetUrl(v.String()).
				SetCallBack(s.ParseIndex))
			if err != nil {
				s.logger.Error(err)
				continue
			}
		}
	}

	return
}

// Test go run cmd/baiduBaikeSpider/* -c dev.yml -m prod
func (s *Spider) Test(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetExtra(&ExtraDetail{
			Keyword: "周口店遗址",
		}).
		SetCallBack(s.ParseDetail))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestIndex go run cmd/baiduBaikeSpider/* -c dev.yml -m prod -f TestIndex
func (s *Spider) TestIndex(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://baike.baidu.com/").
		SetCallBack(s.ParseIndex))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

	spider = &Spider{
		Spider:               baseSpider,
		logger:               baseSpider.GetLogger(),
		collectionBaiduBaike: "baidu_baike",
		reItem:               regexp.MustCompile(`/item/([^/]+)`),
	}
	spider.SetName("baidu-baike")

	return
}

func main() {
	app.NewApp(NewSpider,
		pkg.WithCustomMiddleware(new(Middleware)),
		pkg.WithMongoPipeline(),
	).Run()
}

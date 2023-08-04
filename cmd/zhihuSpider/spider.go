package main

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"regexp"
	"strconv"
)

type Spider struct {
	pkg.Spider
	logger          pkg.Logger
	collectionZhihu string
	reItem          *regexp.Regexp
}

func (s *Spider) ParseDetail(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraDetail
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))

	content := response.BodyText()
	s.logger.Info(content)
	if content == "" {
		return
	}
	data := DataWord{
		Id:      extra.Id,
		Content: content,
	}
	err = s.YieldItem(ctx, items.NewItemMongo(s.collectionZhihu, true).
		SetUniqueKey(strconv.Itoa(extra.Id)).
		SetId(extra.Id).
		SetData(&data))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) ParseIndex(ctx pkg.Context, response pkg.Response) (err error) {
	return
}

// Test go run cmd/zhihuSpider/*.go -c dev.yml -n zhihu -m prod
func (s *Spider) Test(ctx pkg.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetExtra(&ExtraDetail{
			Id: 615389425,
		}).
		SetCallBack(s.ParseDetail))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestQuestion go run cmd/zhihuSpider/*.go -c dev.yml -n zhihu -f TestQuestion -m prod
func (s *Spider) TestQuestion(ctx pkg.Context, _ string) (err error) {
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
		Spider:          baseSpider,
		logger:          baseSpider.GetLogger(),
		collectionZhihu: "zhihu",
		reItem:          regexp.MustCompile(`/item/([^/]+)`),
	}
	spider.WithOptions(
		pkg.WithName("zhihu"),
		pkg.WithCustomMiddleware(new(Middleware)),
		pkg.WithMongoPipeline(),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}

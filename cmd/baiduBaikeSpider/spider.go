package main

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	pkg.Spider
	logger               pkg.Logger
	collectionBaiduBaike string
}

func (s *Spider) ParseDetail(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraDetail
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("ExtraDetail", utils.JsonStr(extra))
	s.logger.Info("BodyBytes", string(response.GetBodyBytes()))

	x, err := response.Xpath()
	if err != nil {
		s.logger.Error(err)
		return
	}

	content := x.FindNodeOne("//div[contains(@class, 'J-content')]").FindStrOne("string(.)")
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

// Test go run cmd/baiduBaikeSpider/* -c dev.yml -m prod
func (s *Spider) Test(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetExtra(&ExtraDetail{
			Keyword: "动物传染病",
		}).
		SetCallBack(s.ParseDetail))
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
	}
	spider.SetName("baidu-baike")

	return
}

func main() {
	app.NewApp(NewSpider,
		pkg.WithCustomMiddleware(new(Middleware)),
		//pkg.WithMongoPipeline(),
	).Run()
}

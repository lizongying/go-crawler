package main

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	pkg.Spider
	logger               pkg.Logger
	collectionBaiduBaike string
}

func (s *Spider) ParseDetail(ctx context.Context, response *pkg.Response) (err error) {
	//s.logger.Info(response.Request.Request.Header)
	var extra ExtraDetail
	err = response.Request.GetExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("Detail", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	x, err := response.Xpath()
	if err != nil {
		s.logger.Error(err)
		return
	}

	content := x.FindNodeOne("//div[contains(@class, 'J-content')]").FindStrOne("string(.)")
	//s.logger.Info(string(response.BodyBytes))
	data := DataWord{
		Id:      extra.Keyword,
		Keyword: extra.Keyword,
		Content: content,
	}
	item := pkg.ItemMongo{
		Update:     true,
		Collection: s.collectionBaiduBaike,
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: extra.Keyword,
			Id:        extra.Keyword,
			Data:      &data,
		},
	}
	err = s.YieldItem(ctx, &item)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return
}

// Test go run cmd/baiduBaikeSpider/* -c dev.yml -m prod
func (s *Spider) Test(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, &pkg.Request{
		Extra: &ExtraDetail{
			Keyword: "动物传染病",
		},
		CallBack: s.ParseDetail,
		//ProxyEnable: true,
	})
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
		pkg.WithMongoPipeline(),
	).Run()
}

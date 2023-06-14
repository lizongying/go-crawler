package main

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/pipelines"
	"github.com/lizongying/go-crawler/pkg/spider"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	*spider.BaseSpider

	collectionBaiduBaike string
}

func (s *Spider) ParseDetail(ctx context.Context, response *pkg.Response) (err error) {
	s.Logger.Info(response.Request.Request.Header)
	extra := response.Request.Extra.(*ExtraDetail)
	s.Logger.Info("Detail", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	x, err := response.Xpath()
	if err != nil {
		s.Logger.Error(err)
		return
	}

	content := x.FindNodeOne("//div[contains(@class, 'J-content')]").FindStrOne("string(.)")
	s.Logger.Info(string(response.BodyBytes))
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
		s.Logger.Error(err)
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

func NewSpider(baseSpider *spider.BaseSpider, logger *logger.Logger) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		logger.Error(err)
		return
	}

	baseSpider.Name = "baidu-baike"
	baseSpider.SetMiddleware(new(Middleware), 9)
	baseSpider.SetPipeline(new(pipelines.MongoPipeline), 141)
	spider = &Spider{
		BaseSpider:           baseSpider,
		collectionBaiduBaike: "baidu_baike",
	}

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}

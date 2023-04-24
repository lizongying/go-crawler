package main

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"github.com/lizongying/go-crawler/pkg/spider"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	*spider.BaseSpider
}

func (s *Spider) RequestImages(ctx context.Context, request *pkg.Request) (err error) {
	extra := request.Extra.(*ExtraTest)
	s.Logger.Info("extra", utils.JsonStr(extra))

	if ctx == nil {
		ctx = context.Background()
	}

	response, err := s.Request(ctx, request)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	s.Logger.Info("len", len(response.BodyBytes))
	extra.SetName("test2")
	s.Logger.Info("extra", utils.JsonStr(extra))

	//item := pkg.Item{
	//	Collection: "test",
	//	Id:         utils.StrMd5(extra.Name),
	//	Data:       extra,
	//}
	//err = s.YieldItem(&item)
	//if err != nil {
	//	s.Logger.Error(err)
	//	return err
	//}

	for i := 0; i < 20; i++ {
		item := pkg.Item{
			Collection: "test",
			Id:         i,
			Data:       extra,
		}
		_ = s.YieldItem(&item)
	}

	return
}

func (s *Spider) ParseImages(ctx context.Context, response *pkg.Response) (err error) {
	request := response.Request
	s.Logger.Info("Images", utils.JsonStr(request))

	if ctx == nil {
		ctx = context.Background()
	}

	s.Logger.Info("len", len(response.BodyBytes))

	return
}

func (s *Spider) RequestImagesAsync(ctx context.Context, request *pkg.Request) (err error) {
	s.Logger.Info("Images", utils.JsonStr(request))
	request.CallBack = s.ParseImages

	if ctx == nil {
		ctx = context.Background()
	}

	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	//s.Logger.Info("len", len(body))

	return
}

func (s *Spider) TestImages(_ context.Context) (err error) {
	request := new(pkg.Request)
	request.Url = "https://chinese.aljazeera.net/wp-content/uploads/2023/03/1-126.jpg"
	request.Extra = &ExtraTest{
		Image: new(pkg.Image),
	}
	err = s.RequestImages(nil, request)
	return
}

func NewSpider(baseSpider *spider.BaseSpider, logger *logger.Logger) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		logger.Error(err)
		return
	}
	baseSpider.Name = "test"
	baseSpider.SetMiddleware(middlewares.NewImageMiddleware(logger), 3)
	spider = &Spider{
		BaseSpider: baseSpider,
	}

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}

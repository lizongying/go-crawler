package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type DumpMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	spider pkg.Spider
	stats  pkg.Stats
}

func (m *DumpMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.stats = spider.GetStats()
	return
}

func (m *DumpMiddleware) ProcessItem(c *pkg.Context) (err error) {
	m.stats.IncItemTotal()

	item := c.Item
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		//m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		//m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	m.logger.Debug("referer", item.GetReferer())
	m.logger.Info("data", utils.JsonStr(data))

	//m.stats.IncItemSuccess()
	err = c.NextItem()
	return
}

func (m *DumpMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(DumpMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	return m
}

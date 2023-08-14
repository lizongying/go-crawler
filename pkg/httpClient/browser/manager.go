package browser

import (
	"context"
	"errors"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/lizongying/go-crawler/pkg"
)

const PoolSize = 2

type Launcher struct {
	*launcher.Launcher
	managed bool
}

type Manager struct {
	pool   chan *Browser
	logger pkg.Logger
	spider pkg.Spider
}

func (m *Manager) FromSpider(spider pkg.Spider) *Manager {
	if m == nil {
		return new(Manager).FromSpider(spider)
	}

	m.logger = spider.GetLogger()
	config := spider.GetCrawler().GetConfig()

	poolSize := PoolSize
	m.pool = make(chan *Browser, poolSize)
	for i := 0; i < poolSize; i++ {
		browser, err := NewBrowser(m.logger, config.GetProxy(), config.GetRequestTimeout())
		if err != nil {
			m.logger.Error(err)
			continue
		}
		m.pool <- browser
	}

	return m
}

func (m *Manager) Pop(ctx context.Context) (browser *Browser, err error) {
	var ok bool
	select {
	case <-ctx.Done():
		err = errors.New("manager timeout")
		return

	case browser, ok = <-m.pool:
		if !ok {
			err = errors.New("manager closed")
			return
		}
		return
	}
}

func (m *Manager) Put(b *Browser) {
	m.pool <- b
}

func (m *Manager) Close() {
	for len(m.pool) > 0 {
		if b := <-m.pool; b != nil {
			ctx := context.Background()
			if err := b.Close(ctx); err != nil {
				m.logger.Error(err)
				continue
			}
		}
	}
	close(m.pool)
}

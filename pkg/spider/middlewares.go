package spider

import (
	"github.com/lizongying/go-crawler/pkg"
	"reflect"
	"sort"
)

func (s *BaseSpider) GetMiddlewares() (middlewares map[uint8]string) {
	middlewares = make(map[uint8]string)
	for _, v := range s.middlewares {
		middlewares[v.GetOrder()] = v.GetName()
	}

	return
}

func (s *BaseSpider) SetMiddleware(middleware pkg.Middleware, order uint8) pkg.Spider {
	middlewareNew := middleware.FromCrawler(s)
	name := reflect.TypeOf(middleware).Elem().String()
	middlewareNew.SetName(name)
	middlewareNew.SetOrder(order)
	for k, v := range s.middlewares {
		if v.GetName() == name && v.GetOrder() != order {
			s.DelMiddleware(k)
			break
		}
	}

	middlewaresNew := append(s.middlewares, middlewareNew)

	sort.Slice(middlewaresNew, func(i, j int) bool {
		return middlewaresNew[i].GetOrder() < middlewaresNew[j].GetOrder()
	})

	s.middlewares = middlewaresNew
	return s
}

func (s *BaseSpider) DelMiddleware(index int) {
	if index < 0 {
		return
	}
	if index >= len(s.middlewares) {
		return
	}

	s.middlewares = append(s.middlewares[:index], s.middlewares[index+1:]...)
	return
}

func (s *BaseSpider) CleanMiddlewares() {
	s.middlewares = make([]pkg.Middleware, 0)
}

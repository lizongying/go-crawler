package spider

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
)

func (s *BaseSpider) GetMiddlewares() (middlewares map[int]string) {
	middlewares = make(map[int]string)
	for k, v := range s.middlewares {
		middlewares[k] = v.GetName()
	}

	return
}

func (s *BaseSpider) ReplaceMiddlewares(middlewares map[int]pkg.Middleware) (err error) {
	middlewaresNameMap := make(map[string]struct{})
	middlewaresOrderMap := make(map[int]struct{})
	for k, v := range middlewares {
		if _, ok := middlewaresNameMap[v.GetName()]; ok {
			err = errors.New("middleware name duplicate")
			s.Logger.Error(err)
			return
		}
		middlewaresNameMap[v.GetName()] = struct{}{}
		if _, ok := middlewaresOrderMap[k]; ok {
			err = errors.New("middleware order duplicate")
			s.Logger.Error(err)
			return
		}
		middlewaresOrderMap[k] = struct{}{}
	}

	s.middlewares = middlewares

	return
}

func (s *BaseSpider) SetMiddleware(middleware pkg.Middleware, order int) {
	for k, v := range s.middlewares {
		if v.GetName() == middleware.GetName() && k != order {
			delete(s.middlewares, k)
			break
		}
	}

	s.middlewares[order] = middleware

	return
}

func (s *BaseSpider) DelMiddleware(name string) {
	for k, v := range s.middlewares {
		if v.GetName() == name {
			delete(s.middlewares, k)
			break
		}
	}

	return
}

func (s *BaseSpider) CleanMiddlewares() {
	s.middlewares = make(map[int]pkg.Middleware)

	return
}

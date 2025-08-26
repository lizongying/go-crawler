package {{.Name}}_spider

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Middleware struct {
	pkg.UnimplementedMiddleware
}

func (m *Middleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	return
}

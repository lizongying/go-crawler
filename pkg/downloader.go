package pkg

type Downloader interface {
	GetMiddlewares() Middlewares
	SetMiddlewares(Middlewares) Downloader
	Download(Context, Request) (Response, error)
}

package pkg

type Downloader interface {
	Download(Context, Request) (Response, error)
}

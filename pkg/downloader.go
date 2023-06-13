package pkg

import (
	"context"
	"errors"
)

type Downloader struct {
	Item               Item
	middlewares        []Middleware
	processRequestFns  []func(*Request) error
	processResponseFns []func(*Response) error

	processItemIndex uint8

	ctx context.Context

	logger Logger
}

func (d *Downloader) SetMiddlewares(middlewares []Middleware) {
	d.middlewares = middlewares
	var processRequestFns []func(*Request) error
	var processResponseFns []func(*Response) error
	for _, v := range middlewares {
		processRequestFns = append(processRequestFns, v.ProcessRequest)
		processResponseFns = append(processResponseFns, v.ProcessResponse)
	}
	d.processRequestFns = processRequestFns
	d.processResponseFns = processResponseFns
}

func (d *Downloader) ProcessRequest(request *Request) (err error) {
	for k, v := range d.processRequestFns {
		name := d.middlewares[k].GetName()
		d.logger.Debug("enter", name, "processRequest")
		e := v(request)
		err = errors.Join(err, e)
		d.logger.Debug("exit", name, "processRequest")
		if errors.Is(e, ErrIgnoreRequest) {
			break
		}
	}
	return
}

func (d *Downloader) ProcessResponse(response *Response) (err error) {
	for k, v := range d.processResponseFns {
		name := d.middlewares[k].GetName()
		d.logger.Debug("enter", name, "ProcessResponse")
		e := v(response)
		err = errors.Join(err, e)
		d.logger.Debug("exit", name, "ProcessResponse")
		if errors.Is(e, ErrIgnoreResponse) {
			break
		}
	}
	return
}

func (d *Downloader) FromCrawler(spider Spider) *Downloader {
	if d == nil {
		return new(Downloader).FromCrawler(spider)
	}
	d.logger = spider.GetLogger()
	return d
}

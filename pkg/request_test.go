package pkg

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"
)

type A struct{}

func (a *A) C(context.Context, *Response) error {
	fmt.Println(222222)
	return nil
}

func TestRequest_Marshal(t *testing.T) {
	proxy, _ := url.Parse("127.0.0.1")
	//retryMaxTimes := uint8(10)
	request := Request{
		SpendTime: time.Minute,
		Proxy:     proxy,
		//RetryMaxTimes: &retryMaxTimes,
		Platform: []Platform{Ipad, Iphone},
		CallBack: (&A{}).C,
	}
	r, e := request.Marshal()
	t.Logf("%+v %+v", r, e)
}

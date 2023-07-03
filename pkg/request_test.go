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
	fmt.Println(111111)
	return nil
}
func (a *A) E(context.Context, *Response, error) {
	fmt.Println(222222)
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

func TestRequestJson_Unmarshal(t *testing.T) {
	proxy, _ := url.Parse("127.0.0.1")
	retryMaxTimes := uint8(10)
	request := RequestJson{
		SpendTime:     uint(time.Minute),
		Proxy:         proxy.String(),
		RetryMaxTimes: &retryMaxTimes,
		Platform:      []string{string(Ipad), string(Iphone)},
		CallBack:      "C",
		ErrBack:       "E",
	}
	callbacks := make(map[string]Callback)
	callbacks["C"] = (&A{}).C
	errbacks := make(map[string]Errback)
	errbacks["E"] = (&A{}).E
	request.SetCallbacks(callbacks)
	request.SetErrbacks(errbacks)
	r, e := request.ToRequest()
	t.Logf("%+v %+v", r, e)
}

package main

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info(string(response.GetBodyBytes()))
	return
}

// TestOk go run cmd/testBrowserSpider/*.go -c example.yml -n test-browser -f TestOk -m dev
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://www.sohu.com/").
		SetClient(pkg.ClientBrowser).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestAjax go run cmd/testBrowserSpider/*.go -c example.yml -n test-browser -f TestAjax -m dev
func (s *Spider) TestAjax(ctx pkg.Context, _ string) (err error) {
	b, _ := json.Marshal(map[string]string{"originationAirportCode": "LGA", "destinationAirportCode": "SLC", "departureDate": "2023-08-10", "returnDate": "2023-08-11", "departureTimeOfDay": "ALL_DAY", "returnTimeOfDay": "ALL_DAY", "adultPassengersCount": "1", "tripType": "roundtrip", "fareType": "USD", "passengerType": "ADULT", "adultsCount": "1", "int": "HOMEQBOMAIR", "reset": "true", "application": "air-booking", "site": "southwest"})

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://www.southwest.com/api/air-booking/v1/air-booking/page/air/booking/shopping").
		SetMethod(pkg.POST).
		SetHeaders(map[string]string{"accept": "application/json, text/javascript, */*; q=0.01",
			"content-type":         "application/json",
			"authorization":        "null null",
			"x-api-idtoken":        "null",
			"x-api-key":            "l7xx944d175ea25f4b9c903a583ea82a1c4c",
			"x-app-id":             "air-booking",
			"x-channel-id":         "southwest",
			"x-swa-di-dtid":        "a06ce16ee1f534f03b82babd3ec7a670ea61",
			"x-swa-di-pid":         "7726946845745073",
			"x-swa-di-ue":          "",
			"x-swa-di-uid":         "56ea542ba5f73b08673847766bf7c0c7f085",
			"x-swa-di-usid":        "3ff654fb14a1e4c8f4c63b4eaf8ccc3a4a84",
			"x-user-experience-id": "07dbd15f-3744-4e1a-87a9-4ca7f580c5b3"}).
		SetBody(string(b)).
		SetAjax(true).
		SetReferrer("https://www.southwest.com").
		SetClient(pkg.ClientBrowser).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("test-browser"),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mockServers"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info(response.BodyStr())
	return
}

// TestSohu go run cmd/testBrowserSpider/*.go -c example.yml -n test-browser -f TestSohu -m once
func (s *Spider) TestSohu(ctx pkg.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://www.sohu.com/").
		SetClient(pkg.ClientBrowser).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestOk go run cmd/testBrowserSpider/*.go -c example.yml -n test-browser -f TestOk -m once
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServers.UrlOk)).
		SetClient(pkg.ClientBrowser).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestAjax go run cmd/testBrowserSpider/*.go -c example.yml -n test-browser -f TestAjax -m once
func (s *Spider) TestAjax(ctx pkg.Context, _ string) (err error) {
	b, _ := json.Marshal(map[string]string{"adultPassengersCount": "1", "adultsCount": "1", "departureDate": "2023-08-15", "departureTimeOfDay": "ALL_DAY", "destinationAirportCode": "SLC", "fareType": "USD", "from": "", "int": "HOMEQBOMAIR", "originationAirportCode": "LGA", "passengerType": "ADULT", "promoCode": "", "reset": "true", "returnDate": "2023-08-18", "returnTimeOfDay": "ALL_DAY", "to": "", "tripType": "roundtrip", "application": "air-booking", "site": "southwest"})

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://www.southwest.com/api/air-booking/v1/air-booking/page/air/booking/shopping").
		SetMethod(pkg.POST).
		SetHeaders(map[string]string{
			"content-type": "application/json",
			"x-api-key":    "l7xx944d175ea25f4b9c903a583ea82a1c4c",
			"x-app-id":     "air-booking",
			"x-channel-id": "southwest",
		}).
		SetBodyBytes(b).
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
		pkg.WithHost("https://localhost:8081"),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run(
		pkg.WithMockServerRoute(mockServers.NewRouteOk),
	)
}

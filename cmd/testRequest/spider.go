package main

import (
	"github.com/lizongying/go-crawler/pkg/request"
	"net/http"
)

func main() {
	_ = request.NewRequest().
		SetUrl("").
		SetMethod(http.MethodPost).
		SetHeaders(map[string]string{
			"content-type": "application/json",
			"x-api-key":    "l7xx944d175ea25f4b9c903a583ea82a1c4c",
			"x-app-id":     "air-booking",
			"x-channel-id": "southwest",
		}).
		SetBodyStr("")
}

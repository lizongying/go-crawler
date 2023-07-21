package pkg

import (
	"net/http"
	"net/url"
)

type Meta struct {
	Cookies  []*http.Cookie
	Referrer *url.URL
}

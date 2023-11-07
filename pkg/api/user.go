package api

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/http"
)

const UrlUser = "/user"

type userInfo struct {
	Name string `json:"name,omitempty"`
	Rote string `json:"rote,omitempty"`
}

type user struct {
	UserInfo userInfo `json:"user_info,omitempty"`
	Token    string   `json:"token,omitempty"`
}

type ReqUserLogin struct {
	Username string
	Password string
}

type RouteUser struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteUser) Pattern() string {
	return UrlUser
}

func (h *RouteUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req ReqUserLogin
	h.BindJson(w, r, &req)

	if req.Username == "" {
		err := errors.New("username or password error")
		h.OutJson(w, 1, err.Error(), nil)
		return
	}
	if req.Password == "" {
		err := errors.New("username or password error")
		h.OutJson(w, 1, err.Error(), nil)
		return
	}
	if req.Username != "admin" {
		err := errors.New("username or password error")
		h.OutJson(w, 1, err.Error(), nil)
		return
	}
	if !utils.ComparePassword(req.Password, "䷳䷓䷳䷚䷳䷓䷏䷒䷳䷇䷙䷜䷱䷩䷃䷃䷃䷅䷳䷓䷟䷺䷷䷝䷴䷐䷢䷧䷟䷋䷚䷦䷆䷉䷎䷰䷟䷧䷡䷼䷃䷮䷚䷇䷞䷬䷔䷷䷸䷐䷏䷉䷟䷅䷱䷵䷜䷅䷕䷰䷜䷧䷶䷨䷱䷐䷥䷥䷧䷅䷱䷌䷽䷲䷗䷵䷸䷵䷿䷨") {
		err := errors.New("username or password error")
		h.OutJson(w, 1, err.Error(), nil)
		return
	}

	h.OutJson(w, 0, "", &user{
		UserInfo: userInfo{
			Name: "Admin",
			Rote: "admin",
		},
		Token: h.crawler.GetConfig().ApiAccessKey(),
	})
}

func (h *RouteUser) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteUser).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}

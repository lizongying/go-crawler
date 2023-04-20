package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/internal"
	"github.com/lizongying/go-crawler/internal/app"
	"github.com/lizongying/go-crawler/internal/logger"
	"github.com/lizongying/go-crawler/internal/spider"
	"github.com/lizongying/go-crawler/internal/utils"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const Video = "EgIQAQ%253D%253D"

type YoutubeSpider struct {
	*spider.BaseSpider

	collectionYoutubeUser string

	urlSearch    string
	urlSearchApi string
	urlUserApi   string
	urlVideos    string

	apiKey          string
	initialDataRe   *regexp.Regexp
	apiKeyRe        *regexp.Regexp
	emailRe         *regexp.Regexp
	urlRe           *regexp.Regexp
	floatRe         *regexp.Regexp
	intRe           *regexp.Regexp
	publishedTimeRe *regexp.Regexp
}

func (s *YoutubeSpider) RequestSearch(ctx context.Context, request *internal.Request) (err error) {
	extra := request.Extra.(*ExtraSearch)
	s.Logger.Info("Search", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	keyword := url.QueryEscape(extra.Keyword)
	request.Url = fmt.Sprintf(s.urlSearch, keyword)
	if extra.Sp == Video {
		request.Url += fmt.Sprintf("&sp=%s", Video)
	}
	response, err := s.Request(ctx, request)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	r := s.initialDataRe.FindSubmatch(response.BodyBytes)
	if len(r) != 2 {
		err = errors.New("not find content")
		s.Logger.Error(err)
		return
	}
	var respSearch RespSearch
	err = json.Unmarshal(r[1], &respSearch)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	token := ""
	for _, v := range respSearch.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents {
		continuationCommand := v.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand
		if continuationCommand.Request == "CONTINUATION_REQUEST_TYPE_SEARCH" {
			token = continuationCommand.Token
		} else {
			for _, v1 := range v.ItemSectionRenderer.Contents {
				if v1.VideoRenderer.VideoID == "" {
					continue
				}

				runs := v1.VideoRenderer.OwnerText.Runs
				if len(runs) < 1 {
					s.Logger.Error("runs err")
					continue
				}
				id := strings.TrimPrefix(runs[0].NavigationEndpoint.BrowseEndpoint.CanonicalBaseURL, "/@")
				e := s.RequestVideos(ctx, &internal.Request{
					ProxyEnable: true,
					UniqueKey:   id,
					Extra: &ExtraUser{
						KeyWord:  extra.Keyword,
						Id:       id,
						Key:      runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID,
						UserName: runs[0].Text,
					},
				})
				if e != nil {
					s.Logger.Error(e)
					continue
				}
			}
		}
	}

	r = s.apiKeyRe.FindSubmatch(response.BodyBytes)
	if len(r) != 2 {
		err = errors.New("not find api-key")
		s.Logger.Error(err)
		return
	}

	s.apiKey = string(r[1])

	if extra.MaxPage > 0 && extra.Page >= extra.MaxPage {
		s.Logger.Info("max page")
		return
	}
	err = s.RequestSearchApi(ctx, &internal.Request{
		ProxyEnable: true,
		Extra: &ExtraSearch{
			Keyword:       extra.Keyword,
			Sp:            extra.Sp,
			Page:          extra.Page + 1,
			MaxPage:       extra.MaxPage,
			NextPageToken: token,
		},
	})
	if err != nil {
		s.Logger.Error(err)
		return
	}

	return
}

func (s *YoutubeSpider) RequestSearchApi(ctx context.Context, request *internal.Request) (err error) {
	extra := request.Extra.(*ExtraSearch)
	s.Logger.Info("SearchApi", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	request.Method = "POST"
	request.Url = fmt.Sprintf(s.urlSearchApi, s.apiKey)
	request.BodyStr = fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"US","clientName":"WEB","clientVersion":"2.20230327.01.00"}},"continuation":"%s"}`, extra.NextPageToken)
	response, err := s.Request(ctx, request)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	var respSearch RespSearchApi
	err = json.Unmarshal(response.BodyBytes, &respSearch)
	if err != nil {
		s.Logger.Error(err)
		return
	}

	token := ""
	onResponseReceivedCommands := respSearch.OnResponseReceivedCommands
	if len(onResponseReceivedCommands) < 1 {
		err = errors.New("onResponseReceivedCommands err")
		s.Logger.Error(err)
		return
	}

	for _, v := range onResponseReceivedCommands[0].AppendContinuationItemsAction.ContinuationItems {
		continuationCommand := v.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand
		if continuationCommand.Request == "CONTINUATION_REQUEST_TYPE_SEARCH" {
			token = continuationCommand.Token
		} else {
			for _, v1 := range v.ItemSectionRenderer.Contents {
				if v1.VideoRenderer.VideoID == "" {
					continue
				}

				runs := v1.VideoRenderer.OwnerText.Runs
				if len(runs) < 1 {
					s.Logger.Error("runs err")
					continue
				}
				id := strings.TrimPrefix(runs[0].NavigationEndpoint.BrowseEndpoint.CanonicalBaseURL, "/@")
				e := s.RequestVideos(ctx, &internal.Request{
					ProxyEnable: true,
					UniqueKey:   id,
					Extra: &ExtraUser{
						KeyWord:  extra.Keyword,
						Id:       id,
						Key:      runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID,
						UserName: runs[0].Text,
					},
				})
				if e != nil {
					s.Logger.Error(e)
					continue
				}
			}
		}
	}

	if token != "" {
		if extra.MaxPage > 0 && extra.Page >= extra.MaxPage {
			s.Logger.Info("max page")
			return
		}
		err = s.RequestSearchApi(ctx, &internal.Request{
			ProxyEnable: true,
			Extra: &ExtraSearch{
				Keyword:       extra.Keyword,
				Sp:            extra.Sp,
				Page:          extra.Page + 1,
				MaxPage:       extra.MaxPage,
				NextPageToken: token,
			},
		})
		if err != nil {
			s.Logger.Error(err)
			return
		}
	}

	return
}

func (s *YoutubeSpider) RequestVideos(ctx context.Context, request *internal.Request) (err error) {
	extra := request.Extra.(*ExtraUser)
	s.Logger.Info("Videos", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	request.Url = fmt.Sprintf(s.urlVideos, extra.Id)
	response, err := s.Request(ctx, request)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	r := s.initialDataRe.FindSubmatch(response.BodyBytes)
	if len(r) != 2 {
		err = errors.New("not find content")
		s.Logger.Error(err)
		return
	}
	var respVideos RespVideos
	err = json.Unmarshal(r[1], &respVideos)
	if err != nil {
		s.Logger.Error(err)
		return
	}

	viewAvg := 0
	viewTotal := 0
	ok := false
	begin := time.Now().AddDate(0, -3, 0)
	for _, v := range respVideos.Contents.TwoColumnBrowseResultsRenderer.Tabs {
		if v.TabRenderer.Title != "Videos" {
			continue
		}

		i := 0
		for _, v1 := range v.TabRenderer.Content.RichGridRenderer.Contents {
			video := v1.RichItemRenderer.Content.VideoRenderer

			videoID := video.VideoID
			if videoID == "" {
				continue
			}

			viewCountText := video.ViewCountText.SimpleText
			viewCount := 0
			if viewCountText != "" && viewCountText != "No views" {
				viewCountInt, e := strconv.Atoi(strings.Join(s.intRe.FindAllString(viewCountText, -1), ""))
				if e != nil {
					s.Logger.Error(e, "viewCount", viewCountText)
					continue
				}
				viewCount = viewCountInt
			}

			t := time.Now().Unix()
			publishedTime := s.publishedTimeRe.FindStringSubmatch(video.PublishedTimeText.SimpleText)
			if len(publishedTime) == 3 {
				i1, _ := strconv.Atoi(publishedTime[1])
				switch publishedTime[2] {
				case "year":
					t -= int64(i1 * 60 * 60 * 24 * 30 * 365)
				case "month":
					t -= int64(i1 * 60 * 60 * 24 * 30)
				case "week":
					t -= int64(i1 * 60 * 60 * 24 * 7)
				case "day":
					t -= int64(i1 * 60 * 60 * 24)
				case "hour":
					t -= int64(i1 * 60 * 60)
				case "minute":
					t -= int64(i1 * 60)
				case "second":
					t -= int64(i1)
				default:
				}
			}
			if time.Unix(t, 0).After(begin) {
				ok = true
			}

			i++
			viewTotal += viewCount
			viewAvg = viewTotal / i
			if i > 6 {
				break
			}
		}
	}

	if !ok {
		s.Logger.Warning("out date")
		return
	}

	subscriber := respVideos.Header.C4TabbedHeaderRenderer.SubscriberCountText.SimpleText
	index := strings.Index(subscriber, " ")
	followers := 0
	if index > 0 {
		followersText := subscriber[0:index]
		followers64, e := strconv.ParseFloat(strings.Join(s.floatRe.FindAllString(followersText, -1), ""), 64)
		if e != nil {
			s.Logger.Error(e, "followers64", subscriber)
		}
		if strings.HasSuffix(followersText, "T") {
			followers = int(followers64 * 1000 * 1000 * 1000 * 1000)
		} else if strings.HasSuffix(followersText, "G") {
			followers = int(followers64 * 1000 * 1000 * 1000)
		} else if strings.HasSuffix(followersText, "M") {
			followers = int(followers64 * 1000 * 1000)
		} else if strings.HasSuffix(followersText, "K") {
			followers = int(followers64 * 1000)
		} else {
			followers = int(followers64)
		}
	}

	description := strings.TrimSpace(respVideos.Metadata.ChannelMetadataRenderer.Description)
	email := ""
	emails := s.emailRe.FindAllString(description, -1)
	if len(emails) > 0 {
		email = emails[0]
	}

	link := ""
	urls := s.urlRe.FindAllString(description, -1)
	if len(urls) > 0 {
		link = urls[0]
	}

	if viewAvg > 1000 {
		data := DataUser{
			Id:          extra.Id,
			UserName:    extra.UserName,
			Description: description,
			Link:        link,
			Email:       email,
			Followers:   followers,
			ViewAvg10:   viewAvg,
			Keyword:     extra.KeyWord,
		}
		item := internal.Item{
			Collection: s.collectionYoutubeUser,
			Id:         extra.Id,
			UniqueKey:  extra.Id,
			Data:       &data,
		}
		err = s.YieldItem(&item)
		if err != nil {
			s.Logger.Error(err)
			return err
		}
	}

	return
}

func (s *YoutubeSpider) RequestUserApi(ctx context.Context, request *internal.Request) (err error) {
	extra := request.Extra.(*ExtraUser)
	s.Logger.Info("UserApi", utils.JsonStr(extra))
	if ctx == nil {
		ctx = context.Background()
	}

	request.Method = "POST"
	request.Url = fmt.Sprintf(s.urlUserApi, s.apiKey)
	request.BodyStr = fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"US","clientName":"WEB","clientVersion":"2.20230327.01.00"}},"browseId":"%s"}`, extra.Key)
	response, err := s.Request(ctx, request)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	var respUser RespUserApi
	err = json.Unmarshal(response.BodyBytes, &respUser)
	if err != nil {
		s.Logger.Error(err)
		return
	}

	viewAvg := 0
	viewTotal := 0
	ok := false
	begin := time.Now().AddDate(0, -3, 0)
	for _, v := range respUser.Contents.TwoColumnBrowseResultsRenderer.Tabs {
		if v.TabRenderer.Title != "Home" {
			continue
		}

		for _, v1 := range v.TabRenderer.Content.SectionListRenderer.Contents {
			for _, v2 := range v1.ItemSectionRenderer.Contents {
				i := 0
				for _, v3 := range v2.ShelfRenderer.Content.HorizontalListRenderer.Items {
					videoID := v3.GridVideoRenderer.VideoID
					if videoID == "" {
						continue
					}

					viewCountText := v3.GridVideoRenderer.ViewCountText.SimpleText
					viewCount := 0
					if viewCountText != "" && viewCountText != "No views" {
						viewCountInt, e := strconv.Atoi(strings.Join(s.intRe.FindAllString(viewCountText, -1), ""))
						if e != nil {
							s.Logger.Error(e, "viewCount", viewCountText)
							continue
						}
						viewCount = viewCountInt
					}

					t := time.Now().Unix()
					publishedTime := s.publishedTimeRe.FindStringSubmatch(v3.GridVideoRenderer.PublishedTimeText.SimpleText)
					if len(publishedTime) == 3 {
						i1, _ := strconv.Atoi(publishedTime[1])
						switch publishedTime[2] {
						case "year":
							t -= int64(i1 * 60 * 60 * 24 * 30 * 365)
						case "month":
							t -= int64(i1 * 60 * 60 * 24 * 30)
						case "week":
							t -= int64(i1 * 60 * 60 * 24 * 7)
						case "day":
							t -= int64(i1 * 60 * 60 * 24)
						case "hour":
							t -= int64(i1 * 60 * 60)
						case "minute":
							t -= int64(i1 * 60)
						case "second":
							t -= int64(i1)
						default:
						}
					}
					if time.Unix(t, 0).After(begin) {
						ok = true
					}

					i++
					viewTotal += viewCount
					viewAvg = viewTotal / i
					if i > 10 {
						break
					}
				}
			}
		}
	}

	if !ok {
		s.Logger.Warning("out date")
		return
	}

	subscriber := respUser.Header.C4TabbedHeaderRenderer.SubscriberCountText.SimpleText
	index := strings.Index(subscriber, " ")
	followers := 0
	if index > 0 {
		followersText := subscriber[0:index]
		followers64, e := strconv.ParseFloat(strings.Join(s.floatRe.FindAllString(followersText, -1), ""), 64)
		if e != nil {
			s.Logger.Error(e, "followers64", subscriber)
		}
		if strings.HasSuffix(followersText, "T") {
			followers = int(followers64 * 1000 * 1000 * 1000 * 1000)
		} else if strings.HasSuffix(followersText, "G") {
			followers = int(followers64 * 1000 * 1000 * 1000)
		} else if strings.HasSuffix(followersText, "M") {
			followers = int(followers64 * 1000 * 1000)
		} else if strings.HasSuffix(followersText, "K") {
			followers = int(followers64 * 1000)
		} else {
			followers = int(followers64)
		}
	}

	description := strings.TrimSpace(respUser.Metadata.ChannelMetadataRenderer.Description)
	email := ""
	r := s.emailRe.FindAllString(description, -1)
	if len(r) > 0 {
		email = r[0]
	}

	link := ""
	urls := s.urlRe.FindAllString(description, -1)
	if len(urls) > 0 {
		link = urls[0]
	}

	if viewAvg > 1000 && viewAvg < 100000 {
		data := DataUser{
			Id:          extra.Id,
			UserName:    extra.UserName,
			Description: description,
			Link:        link,
			Email:       email,
			Followers:   followers,
			ViewAvg10:   viewAvg,
			Keyword:     extra.KeyWord,
		}
		item := internal.Item{
			Collection: s.collectionYoutubeUser,
			Id:         data.Id,
			Data:       &data,
		}
		err = s.YieldItem(&item)
		if err != nil {
			s.Logger.Error(err)
			return err
		}
	}

	return
}

func (s *YoutubeSpider) Test(_ context.Context) (err error) {
	err = s.RequestVideos(nil, &internal.Request{
		ProxyEnable: true,
		Extra: &ExtraUser{
			Id: "sierramarie",
		},
	})
	return
}

func (s *YoutubeSpider) FromKeyword(_ context.Context) (err error) {
	for _, v := range []string{
		"veja",
		"tote bag",
	} {
		err = s.RequestSearch(nil, &internal.Request{
			ProxyEnable: true,
			Extra: &ExtraSearch{
				Keyword: v,
				Sp:      Video,
				Page:    1,
				//MaxPage: 1,
			},
		})
	}

	return
}

func NewYoutubeSpider(baseSpider *spider.BaseSpider, logger *logger.Logger) (spider internal.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("spider is empty")
		logger.Error(err)
		return
	}

	baseSpider.Name = "youtube"
	baseSpider.TimeoutRequest = time.Second * 30
	baseSpider.SetMiddleware(NewYoutubeMiddleware(logger), 90)
	spider = &YoutubeSpider{
		BaseSpider:            baseSpider,
		collectionYoutubeUser: "youtube_user_202304131",
		urlSearch:             "https://www.youtube.com/results?search_query=%s",
		urlSearchApi:          "https://www.youtube.com/youtubei/v1/search?key=%s",
		urlUserApi:            "https://www.youtube.com/youtubei/v1/browse?key=%s",
		urlVideos:             "https://www.youtube.com/@%s/videos",

		apiKey:          "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8",
		initialDataRe:   regexp.MustCompile(`ytInitialData = (.+);</script>`),
		apiKeyRe:        regexp.MustCompile(`"INNERTUBE_API_KEY":"([^"]+)`),
		emailRe:         regexp.MustCompile(`(\w+[-+.]*\w+@\w+[-.]*\w+\.\w+[-.]*\w+)`),
		urlRe:           regexp.MustCompile(`(?i)\b((?:https?://|www\d{0,3}[.]|[a-z0-9.-]+[.][a-z]{2,4}/)(?:[^\s()<>]+|\(([^\s()<>]+|(\([^\s()<>]+\)))*\))+(?:\(([^\s()<>]+|(\([^\s()<>]+\)))*\)|[^\s\` + "`" + `!()\[\]{};:'".,<>?«»“”‘’]))`),
		floatRe:         regexp.MustCompile(`[\d.]`),
		intRe:           regexp.MustCompile(`\d`),
		publishedTimeRe: regexp.MustCompile(`(\d+)\s*(year|month|week|day|hour|minute|second)`),
	}

	return
}

func main() {
	app.NewApp(NewYoutubeSpider).Run()
}

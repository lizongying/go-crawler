package test_item_spider

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strconv"
)

type Spider struct {
	pkg.Spider
	logger         pkg.Logger
	collectionTest string
	tableTest      string
	topicTest      string
	fileNameTest   string
}

func (s *Spider) ParseMysql(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.Extra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.UnsafeJSON(extra))
	s.logger.Info("response", response.Text())

	if extra.Count > 0 {
		return
	}

	if err = s.YieldItem(ctx, items.NewItemMysql(s.tableTest, true).
		SetUniqueKey("1").
		SetId("3").
		SetData(&DataOk{
			Id: "3",
			A:  0,
			B:  2,
			C:  "",
			D:  "2",
		})); err != nil {
		s.logger.Error(err)
		return
	}

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseMysql)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) ParseKafka(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.Extra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.UnsafeJSON(extra))
	s.logger.Info("response", response.Text())

	if extra.Count > 0 {
		return
	}

	if err = s.YieldItem(ctx, items.NewItemKafka(s.topicTest).
		SetUniqueKey("1").
		SetId("3").
		SetData(&DataOk{
			Id: "3",
			A:  0,
			B:  2,
			C:  "",
			D:  "2",
		})); err != nil {
		s.logger.Error(err)
		return
	}

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseKafka)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) ParseMongo(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.Extra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.UnsafeJSON(extra))
	s.logger.Info("response", response.Text())

	if extra.Count > 0 {
		return
	}

	if err = s.YieldItem(ctx, items.NewItemMongo(s.collectionTest, true).
		SetUniqueKey("1").
		SetId(extra.Count).
		SetData(&DataOk{
			Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
			Count: extra.Count,
		})); err != nil {
		s.logger.Error(err)
		return
	}

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseMongo)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) ParseCsv(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.Extra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.UnsafeJSON(extra))
	s.logger.Info("response", response.Text())

	if extra.Count > 2 {
		return
	}

	if err = s.YieldItem(ctx, items.NewItemCsv(s.fileNameTest).
		SetUniqueKey("1").
		SetId(extra.Count).
		SetData(&DataOk{
			Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
			Count: extra.Count,
		})); err != nil {
		s.logger.Error(err)
		return
	}

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseCsv)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) ParseJsonl(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.Extra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.UnsafeJSON(extra))
	s.logger.Info("response", response.Text())

	if extra.Count > 2 {
		return
	}

	if err = s.YieldItem(ctx, items.NewItemJsonl(s.fileNameTest).
		SetUniqueKey("1").
		SetId(extra.Count).
		SetData(&DataOk{
			Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
			Count: extra.Count,
		})); err != nil {
		s.logger.Error(err)
		return
	}

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseJsonl)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) ParseSqlite(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.Extra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.UnsafeJSON(extra))
	s.logger.Info("response", response.Text())

	if extra.Count > 2 {
		return
	}

	if err = s.YieldItem(ctx, items.NewItemSqlite(s.tableTest, true).
		SetUniqueKey(strconv.Itoa(extra.Count)).
		SetId(strconv.Itoa(extra.Count)).
		SetData(&DataOk{
			Id: strconv.Itoa(extra.Count),
			A:  0,
			B:  2,
			C:  "",
			D:  "2",
		})); err != nil {
		s.logger.Error(err)
		return
	}

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseSqlite)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestMongo go run cmd/test_item_spider/*.go -c dev.yml -n test-item -f TestMongo -m once
func (s *Spider) TestMongo(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseMongo)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestMysql go run cmd/test_item_spider/*.go -c dev.yml -n test-item -f TestMysql -m once
func (s *Spider) TestMysql(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseMysql)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestKafka go run cmd/test_item_spider/*.go -c dev.yml -n test-item -f TestKafka -m once
func (s *Spider) TestKafka(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseKafka)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestCsv go run cmd/test_item_spider/*.go -c dev.yml -n test-item -f TestCsv -m once
func (s *Spider) TestCsv(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseCsv)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestJsonl go run cmd/test_item_spider/*.go -c dev.yml -n test-item -f TestJsonl -m once
func (s *Spider) TestJsonl(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseJsonl)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestSqlite go run cmd/test_item_spider/*.go -c example.yml -n test-item -f TestSqlite -m once
// CREATE TABLE IF NOT EXISTS test (
//
//	id INTEGER PRIMARY KEY AUTOINCREMENT,
//	count INTEGER NOT NULL,
//	a INTEGER NOT NULL,
//	b INTEGER NOT NULL,
//	c TEXT NOT NULL,
//	d TEXT NOT NULL)
func (s *Spider) TestSqlite(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseSqlite)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider:         baseSpider,
		logger:         baseSpider.GetLogger(),
		collectionTest: "test",
		tableTest:      "test",
		topicTest:      "test",
		fileNameTest:   "test",
	}
	spider.WithOptions(
		pkg.WithName("test-item"),
		pkg.WithHost("https://localhost:8081"),

		//pkg.WithMongoPipeline(),
		//pkg.WithCsvPipeline(),
		pkg.WithJsonLinesPipeline(),
		//pkg.WithMysqlPipeline(),
		//pkg.WithKafkaPipeline(),
		//pkg.WithSqlitePipeline(),
	)

	return
}

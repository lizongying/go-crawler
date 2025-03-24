package statistics

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	pb "github.com/lizongying/go-crawler/pkg/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client  pb.StatisticsClient
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (s *Client) crawlerChanged(crawler pkg.Context) (err error) {
	ctx := context.Background()
	_, err = s.client.CrawlerChanged(ctx, &pb.Crawler{})
	return
}
func (s *Client) spiderChanged(spider pkg.Context) (err error) {
	ctx := context.Background()
	_, err = s.client.SpiderChanged(ctx, &pb.Spider{})
	return
}
func (s *Client) jobChanged(job pkg.Context) (err error) {
	ctx := context.Background()
	_, err = s.client.JobChanged(ctx, &pb.Job{})
	return
}
func (s *Client) taskChanged(task pkg.Task) (err error) {
	ctx := context.Background()
	_, err = s.client.TaskChanged(ctx, &pb.Task{})
	return
}
func (s *Client) requestChanged(request pkg.Request) (err error) {
	ctx := context.Background()
	_, err = s.client.RequestChanged(ctx, &pb.Request{})
	return
}
func (s *Client) itemChanged(item pkg.Item) (err error) {
	ctx := context.Background()
	_, err = s.client.ItemChanged(ctx, &pb.Item{})
	return
}
func (s *Client) FromCrawler(crawler pkg.Crawler) (c *Client) {
	if s == nil {
		return new(Client).FromCrawler(crawler)
	}

	s.crawler = crawler
	s.logger = crawler.GetLogger()

	var host string
	var port int

	var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],
                "waitForReady": true,

                "retryPolicy": {
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]
        }`

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(retryPolicy))
	if err != nil {
		s.logger.Error(err)
		return
	}

	s.client = pb.NewStatisticsClient(conn)

	signal := s.crawler.GetSignal()
	signal.RegisterCrawlerChanged(s.crawlerChanged)
	signal.RegisterSpiderChanged(s.spiderChanged)
	signal.RegisterJobChanged(s.jobChanged)
	signal.RegisterTaskChanged(s.taskChanged)
	signal.RegisterRequestChanged(s.requestChanged)
	signal.RegisterItemChanged(s.itemChanged)
	return s
}

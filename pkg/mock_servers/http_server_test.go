package mock_servers

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/loggers"
	"go.uber.org/fx"
	"os"
	"testing"
)

func Run(routeFn func(pkg.Logger) pkg.Route) {
	_ = os.Setenv("CRAWLER_CONFIG_FILE", "/Users/lizongying/IdeaProjects/go-crawler/dev.yml")
	fx.New(
		fx.Provide(
			cli.NewCli,
			NewHttpServer,
			fx.Annotate(
				loggers.NewLogger,
				fx.As(new(pkg.Logger)),
			),
			config.NewConfig,
		),
		fx.Invoke(func(logger pkg.Logger, mockServer pkg.MockServer, shutdowner fx.Shutdowner) {
			mockServer.AddRoutes(routeFn(logger))
			_ = mockServer.Run()

			return
		}),
	).Run()
}

// go test -v ./pkg/mock_servers/*.go -run TestNewRouteRobotsTxt
func TestNewRouteRobotsTxt(t *testing.T) {
	Run(NewRouteRobotsTxt)
}

// go test -v ./pkg/mock_servers/*.go -run TestOk
func TestOk(t *testing.T) {
	Run(NewRouteOk)
}

// go test -v ./pkg/mock_servers/*.go -run TestHtml
func TestHtml(t *testing.T) {
	Run(NewRouteHtml)
}

// go test -v ./pkg/mock_servers/*.go -run TestHello
func TestHello(t *testing.T) {
	Run(NewRouteHello)
}

// go test -v ./pkg/mock_servers/*.go -run TestBadGateway
// curl https://localhost:8081/bad-gateway -k
// curl https://localhost:8081/bad-gateway --cacert static/tls/ca.crt
func TestBadGateway(t *testing.T) {
	Run(NewRouteBadGateway)
}

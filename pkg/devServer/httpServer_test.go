package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"go.uber.org/fx"
	"os"
	"testing"
)

// go test -v ./pkg/devServer/*.go -run NewHttpServer
func TestNewHttpServer(t *testing.T) {
	_ = os.Setenv("CRAWLER_CONFIG_FILE", "/Users/lizongying/IdeaProjects/go-crawler/dev.yml")
	fx.New(
		fx.Provide(
			cli.NewCli,
			NewHttpServer,
			fx.Annotate(
				logger.NewLogger,
				fx.As(new(pkg.Logger)),
			),
			config.NewConfig,
		),
		fx.Invoke(func(logger pkg.Logger, devServer pkg.DevServer, shutdowner fx.Shutdowner) {
			devServer.AddRoutes(NewRouteRobotsTxt(logger))
			_ = devServer.Run()

			return
		}),
	).Run()
}

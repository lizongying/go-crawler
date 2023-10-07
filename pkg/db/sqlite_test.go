package db

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/loggers"
	"go.uber.org/fx"
	"log"
	"os"
	"testing"
)

// go test -v ./pkg/db/*.go -run TestNewSqlite
func TestNewSqlite(t *testing.T) {
	_ = os.Setenv("CRAWLER_CONFIG_FILE", "/Users/lizongying/IdeaProjects/go-crawler/dev.yml")
	fx.New(
		fx.Provide(
			cli.NewCli,
			fx.Annotate(
				loggers.NewLogger,
				fx.As(new(pkg.Logger)),
			),
			fx.Annotate(
				NewSqlite,
				fx.As(new(pkg.Sqlite)),
			),
			config.NewConfig,
		),
		fx.Invoke(func(logger pkg.Logger, sqlite pkg.Sqlite, shutdowner fx.Shutdowner) {
			var err error
			logger.Infof("Client %+v\n", sqlite.Client())
			db := sqlite.Client()
			_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, NAME TEXT)")
			_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")

			rows, err := db.Query("SELECT id, name FROM users")
			for rows.Next() {
				var id int
				var name string
				err := rows.Scan(&id, &name)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("ID: %d, Name: %s\n", id, name)
			}

			err = shutdowner.Shutdown()
			if err != nil {
				logger.Error(err)
				return
			}

			return
		}),
	).Run()
}

package db

import (
	"context"
	"database/sql"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

type Sqlite struct {
	Config *config.Sqlite
	*sql.DB
	logger pkg.Logger
}

func (s *Sqlite) Client() *sql.DB {
	return s.DB
}

func NewSqlite(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (sqlite *Sqlite, err error) {
	sqlite = new(Sqlite)
	sqlite.logger = logger

	for _, v := range config.Sqlite {
		path := v.Path
		if path == "" {
			logger.Warn("path is empty")
			continue
		}

		sqlite.Config = v

		var db *sql.DB
		db, err = sql.Open("sqlite3", path)
		if err != nil {
			logger.Error(err)
			return
		}

		err = db.Ping()
		if err != nil {
			logger.Error(err)
			return
		}

		sqlite.DB = db
		break
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) (err error) {
			if sqlite.DB == nil {
				return
			}

			err = sqlite.DB.Close()
			if err != nil {
				logger.Error(err)
				return
			}
			return
		},
	})
	return
}

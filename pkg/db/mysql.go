package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"go.uber.org/fx"
)

func NewMysql(config *config.Config, logger *logger.Logger, lc fx.Lifecycle) (db *sql.DB, err error) {
	uri := config.Mysql.Example.Uri
	if uri == "" {
		err = errors.New("uri is empty")
		logger.Error(err)
		return
	}

	database := config.Mysql.Example.Database
	if database == "" {
		err = errors.New("database is empty")
		logger.Error(err)
		return
	}

	db, err = sql.Open("mysql", fmt.Sprintf("%s/%s", uri, database))
	if err != nil {
		logger.Error(err)
		return
	}

	err = db.Ping()
	if err != nil {
		logger.Error(err)
		return
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) (err error) {
			err = db.Close()
			if err != nil {
				logger.Error(err)
				return
			}
			return
		},
	})
	return
}

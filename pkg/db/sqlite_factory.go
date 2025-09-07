package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
	"sync"
)

type SqliteFactory struct {
	config sync.Map
	logger pkg.Logger

	clients sync.Map
}

func (s *SqliteFactory) GetClient(name string) (db *sql.DB, err error) {
	if v, ok := s.clients.Load(name); ok {
		return v.(*sql.DB), nil
	}

	c, ok := s.config.Load(name)
	if !ok {
		return nil, fmt.Errorf("sqlite config %s not found", name)
	}

	conf := c.(pkg.Sqlite)
	path := conf.Path
	if path == "" {
		s.logger.Warn("path is empty, using in-memory database")
		path = ":memory:"
	}

	db, err = sql.Open("sqlite3", path)
	if err != nil {
		s.logger.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		s.logger.Error(err)
		return
	}

	actual, loaded := s.clients.LoadOrStore(name, db)
	if loaded {
		_ = db.Close()
	}

	return actual.(*sql.DB), nil
}

func (s *SqliteFactory) Close(_ context.Context) error {
	s.clients.Range(func(key, value interface{}) bool {
		_ = value.(*sql.DB).Close()
		return true
	})
	return nil
}

func NewSqliteFactory(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (sqliteFactory *SqliteFactory, err error) {
	sqliteFactory = &SqliteFactory{
		logger: logger,
	}
	for _, i := range config.SqliteList {
		sqliteFactory.config.Store(i.Name, i)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) (err error) {
			return sqliteFactory.Close(ctx)
		},
	})
	return
}

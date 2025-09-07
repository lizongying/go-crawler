package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"go.uber.org/fx"
	"sync"
)

type MysqlFactory struct {
	config sync.Map
	logger pkg.Logger

	clients sync.Map
}

func (m *MysqlFactory) GetClient(name string) (db *sql.DB, err error) {
	if v, ok := m.clients.Load(name); ok {
		return v.(*sql.DB), nil
	}

	c, ok := m.config.Load(name)
	if !ok {
		return nil, fmt.Errorf("mysql config %s not found", name)
	}

	conf := c.(pkg.Mysql)

	uri := conf.Uri
	if uri == "" {
		m.logger.Warn("uri is empty")
		return
	}

	database := conf.Database
	if database == "" {
		m.logger.Warn("database is empty")
		return
	}

	db, err = sql.Open("mysql", fmt.Sprintf("%s/%s", uri, database))
	if err != nil {
		m.logger.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		m.logger.Error(err)
		return
	}

	actual, loaded := m.clients.LoadOrStore(name, db)
	if loaded {
		_ = db.Close()
	}

	return actual.(*sql.DB), nil
}

func (m *MysqlFactory) Close(_ context.Context) error {
	m.clients.Range(func(key, value interface{}) bool {
		_ = value.(*sql.DB).Close()
		return true
	})
	return nil
}

func NewMysqlFactory(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (mysqlFactory *MysqlFactory, err error) {
	mysqlFactory = &MysqlFactory{
		logger: logger,
	}
	for _, i := range config.MysqlList {
		mysqlFactory.config.Store(i.Name, i)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) (err error) {
			return mysqlFactory.Close(ctx)
		},
	})
	return
}

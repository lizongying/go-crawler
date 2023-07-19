package pipelines

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"reflect"
	"strings"
	"time"
)

type MysqlPipeline struct {
	pkg.UnimplementedPipeline
	crawler pkg.Crawler
	stats   pkg.Stats
	logger  pkg.Logger
	mysql   *sql.DB
	timeout time.Duration
}

func (m *MysqlPipeline) ProcessItem(ctx context.Context, item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}
	if item.GetName() != pkg.ItemMysql {
		m.logger.Warn("item not support", pkg.ItemMysql)
		return
	}
	itemMysql, ok := item.(*items.ItemMysql)
	if !ok {
		m.logger.Warn("item parsing failed with", pkg.ItemMysql)
		return
	}

	if itemMysql.GetTable() == "" {
		err = errors.New("table is empty")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	if m.crawler.GetMode() == "test" {
		m.logger.Debug("current mode don't need save")
		m.stats.IncItemIgnore()
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	refType := reflect.TypeOf(itemMysql.GetData()).Elem()
	refValue := reflect.ValueOf(itemMysql.GetData()).Elem()
	var columns []string
	var values []any
	for i := 0; i < refType.NumField(); i++ {
		column := refType.Field(i).Tag.Get("column")
		if column == "" {
			column = refType.Field(i).Name
		}
		columns = append(columns, fmt.Sprintf("%s=?", column))
		value := refValue.Field(i).Interface()
		values = append(values, value)
	}

	s := fmt.Sprintf(`INSERT %s SET %s`, itemMysql.GetTable(), strings.Join(columns, ","))
	stmt, err := m.mysql.PrepareContext(ctx, s)
	if err != nil {
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}
	res, err := stmt.ExecContext(ctx, values...)
	if err != nil {
		e, o := err.(*mysql.MySQLError)
		if !o {
			m.logger.Error(e)
			m.stats.IncItemError()
			return
		}

		if itemMysql.GetUpdate() && !reflect.ValueOf(itemMysql.GetId()).IsZero() && e.Number == 1062 {
			s = fmt.Sprintf(`UPDATE %s SET %s WHERE id=?`, itemMysql.GetTable(), strings.Join(columns, ","))
			values = append(values, itemMysql.GetId())
			stmt, err = m.mysql.PrepareContext(ctx, s)
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				return
			}

			res, err = stmt.ExecContext(ctx, values...)
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				return
			}

			_, err = res.RowsAffected()
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				return
			}

			m.logger.Info(itemMysql.GetTable(), "update success", itemMysql.GetId())
		} else {
			m.logger.Error(err)
			m.stats.IncItemError()
			return
		}
	} else {
		id, e := res.LastInsertId()
		if e != nil {
			m.logger.Error(e)
			m.stats.IncItemError()
			return
		}

		m.logger.Info(itemMysql.GetTable(), "insert success", id)
	}

	m.stats.IncItemSuccess()
	return
}

func (m *MysqlPipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(MysqlPipeline).FromCrawler(crawler)
	}

	m.crawler = crawler
	m.stats = crawler.GetStats()
	m.logger = crawler.GetLogger()
	m.mysql = crawler.GetMysql()
	m.timeout = time.Minute
	return m
}

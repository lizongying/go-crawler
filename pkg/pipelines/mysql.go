package pipelines

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/lizongying/go-crawler/pkg"
	"reflect"
	"strings"
	"time"
)

type MysqlPipeline struct {
	pkg.UnimplementedPipeline
	info    *pkg.SpiderInfo
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

	itemMysql, ok := item.(*pkg.ItemMysql)
	if !ok {
		m.logger.Warn("item not support mysql")
		return
	}

	if itemMysql.Table == "" {
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

	if m.info.Mode == "test" {
		m.logger.Debug("current mode don't need save")
		m.stats.IncItemIgnore()
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	refType := reflect.TypeOf(itemMysql.Data).Elem()
	refValue := reflect.ValueOf(itemMysql.Data).Elem()
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

	s := fmt.Sprintf(`INSERT %s SET %s`, itemMysql.Table, strings.Join(columns, ","))
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

		if itemMysql.Update && !reflect.ValueOf(itemMysql.Id).IsZero() && e.Number == 1062 {
			s = fmt.Sprintf(`UPDATE %s SET %s WHERE id=?`, itemMysql.Table, strings.Join(columns, ","))
			values = append(values, itemMysql.Id)
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

			m.logger.Info(itemMysql.Table, "update success", itemMysql.Id)
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

		m.logger.Info(itemMysql.Table, "insert success", id)
	}

	m.stats.IncItemSuccess()
	return
}

func (m *MysqlPipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(MysqlPipeline).FromCrawler(crawler)
	}

	m.info = crawler.GetInfo()
	m.stats = crawler.GetStats()
	m.logger = crawler.GetLogger()
	m.mysql = crawler.GetMysql()
	m.timeout = time.Minute
	return m
}

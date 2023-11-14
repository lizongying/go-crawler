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
	env     string
	logger  pkg.Logger
	mysql   *sql.DB
	timeout time.Duration
}

func (m *MysqlPipeline) ProcessItem(item pkg.Item) (err error) {
	spider := m.Spider()
	task := item.GetContext().GetTask()

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	if item.Name() != pkg.ItemMysql {
		m.logger.Warn("item not support", pkg.ItemMysql)
		return
	}

	itemMysql, ok := item.GetItem().(*items.ItemMysql)
	if !ok {
		m.logger.Warn("item parsing failed with", pkg.ItemMysql)
		return
	}

	if itemMysql.GetTable() == "" {
		err = errors.New("table is empty")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	data := item.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	item.GetContext().WithItemProcessed(true)

	if m.env == "dev" {
		m.logger.Debug("current mode don't need save")
		task.IncItemIgnore()
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	refType := reflect.TypeOf(itemMysql.Data()).Elem()
	refValue := reflect.ValueOf(itemMysql.Data()).Elem()
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
		task.IncItemError()
		return
	}
	res, err := stmt.ExecContext(ctx, values...)
	if err != nil {
		var e *mysql.MySQLError
		o := errors.As(err, &e)
		if !o {
			m.logger.Error(err)
			task.IncItemError()
			return
		}

		if itemMysql.GetUpdate() && !reflect.ValueOf(itemMysql.Id()).IsZero() && e.Number == 1062 {
			s = fmt.Sprintf(`UPDATE %s SET %s WHERE id=?`, itemMysql.GetTable(), strings.Join(columns, ","))
			values = append(values, itemMysql.Id())
			stmt, err = m.mysql.PrepareContext(ctx, s)
			if err != nil {
				m.logger.Error(err)
				task.IncItemError()
				return
			}

			res, err = stmt.ExecContext(ctx, values...)
			if err != nil {
				m.logger.Error(err)
				task.IncItemError()
				return
			}

			_, err = res.RowsAffected()
			if err != nil {
				m.logger.Error(err)
				task.IncItemError()
				return
			}

			m.logger.Info(itemMysql.GetTable(), "update success", itemMysql.Id())
		} else {
			m.logger.Error(err)
			task.IncItemError()
			return
		}
	} else {
		id, e := res.LastInsertId()
		if e != nil {
			m.logger.Error(e)
			task.IncItemError()
			return
		}

		m.logger.Info(itemMysql.GetTable(), "insert success", id)
	}

	item.GetContext().WithItemStatus(pkg.ItemStatusSuccess)
	spider.GetCrawler().GetSignal().ItemChanged(item)
	task.IncItemSuccess()
	return
}

func (m *MysqlPipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(MysqlPipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	crawler := spider.GetCrawler()
	m.env = spider.GetConfig().GetEnv()
	m.logger = spider.GetLogger()
	m.mysql = crawler.GetMysql()
	if m.mysql == nil {
		err = errors.New("mysql nil")
		return
	}
	m.timeout = time.Minute
	return
}

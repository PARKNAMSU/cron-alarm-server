package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/jmoiron/sqlx"
	"nspark-cron-alarm.com/cron-alarm-server/app/config"
)

/*
db engine 타입.

description:

	확장성 있는 db 연결을 위해 해당 타입으로 DBMS 종류 정의
*/
type dbEngine = string

/* db connect 시 설정 데이터 타입 */
type dbConfig struct {
	user             string
	password         string
	host             string
	database         string
	charset          string
	maxAllowedPacket uint
}

/*
외부에서 사용할 DB Connect.

description:

	내부 인자 직접 변경 및 사용할 수 없게 private 하게 설정
*/
type CustomDB struct {
	engine        dbEngine // DBMS 타입
	conn          *sqlx.DB // connection 객체
	config        dbConfig // db 설정
	tx            *sqlx.Tx // transaction 객체
	isTransaction bool     // transaction 사용 여부
}

var (
	mysqlEngine dbEngine = "mysql"
)

func (c *CustomDB) connect() {
	if c.conn != nil {
		return
	}
	connector := ""

	switch c.engine {
	case mysqlEngine:
		connector = getMysqlConnect(c.config)
	default:
	}

	db, err := sqlx.Connect(
		string(c.engine),
		connector,
	)

	if err != nil {
		log.Panicln(err.Error())
		return
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetConnMaxIdleTime(time.Minute)

	maxIdleConn := 10
	maxOpenConn := 10
	switch config.ENVIRONMENT {
	case "production":
		maxIdleConn = 30
		maxOpenConn = 30
	case "staging":
		maxIdleConn = 20
		maxOpenConn = 20
	}

	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	c.conn = db
}

func (c *CustomDB) Close() {
	if c.conn == nil {
		return
	}
	c.conn.Close()
}

func (c *CustomDB) transaction() {
	if tx, err := c.conn.BeginTxx(context.TODO(), nil); err != nil {
		fmt.Println("[err]:[setTransaction]:[", err.Error(), "]")
		c.tx = tx
	}
}

func (c *CustomDB) Commit() {
	if c.tx == nil {
		return
	}
	if err := c.tx.Commit(); err != nil {
		fmt.Println("[err]:[Commit]:[", err.Error(), "]")
	}
	c.tx = nil
}

func (c *CustomDB) Rollback() {
	if c.tx == nil {
		return
	}
	if err := c.tx.Rollback(); err != nil {
		fmt.Println("[err]:[Rollback]:[", err.Error(), "]")
	}
	c.tx = nil
}

func (c *CustomDB) QueryExecute(query string, queryParams ...any) (sql.Result, error) {
	if c.isTransaction {
		return c.tx.Exec(query, queryParams...)
	}
	return c.conn.Exec(query, queryParams...)
}

func (c *CustomDB) QuerySelect(data any, query string, queryParams ...any) {
	if c.isTransaction {
		c.tx.Select(data, query, queryParams...)
	}

	if reflect.TypeOf(data).Kind() != reflect.Slice {
		var list []any

		err := c.conn.Select(&list, query, queryParams...)

		fmt.Println("[err]:[QuerySelect]:[", err.Error(), "]")

		if len(list) > 0 {
			data = &list[0]
		}
		return
	}

	err := c.conn.Select(data, query, queryParams...)

	if err != nil {
		fmt.Println("[err]:[QuerySelect]:[", err.Error(), "]")
	}

}

func (c *CustomDB) NamedQueryExecute(query string, params any) (sql.Result, error) {
	if c.isTransaction {
		return c.tx.NamedExec(query, params)
	}
	return c.conn.NamedExec(query, params)
}

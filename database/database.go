package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	user     = "root"
	host     = "localhost"
	database = "manufacture_supplier"
	port     = 3307
	password = "12345678"
)

var db *sqlx.DB

// Connect 连接数据库
func init() {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, password, host, port, database)
	var error error
	db, error = sqlx.Connect("mysql", dataSourceName)
	if error != nil {
		errMsg := fmt.Sprintf("mysql connect failed: %s\n", error)
		panic(errMsg)
	}

	db.SetMaxIdleConns(100)
}

// NewDB 构造数据库句柄
func NewDB() *sqlx.DB {
	return db
}

package entities

import (
	"database/sql"

	

	"github.com/jinzhu/gorm"
	// gorm mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var gormDb *gorm.DB

func init() {
	

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	gormDb = db
}

// SQLExecer interface for supporting sql.DB and sql.Tx to do sql statement
type SQLExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// DaoSource Data Access Object Source
type DaoSource struct {
	// if DB, each statement execute sql with random conn.
	// if Tx, all statements use the same conn as the Tx's connection
	SQLExecer
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
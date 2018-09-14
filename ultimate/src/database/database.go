package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"ultimate/src/infrastructure"
)

var database *sql.DB

func Connection() *sql.DB  {

	var err error
	database , err = sql.Open("mysql","mysql:mysql#root@tcp(47.92.165.70:3306)/mydb?charset=utf8")
	infrastructure.Error(err)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)
	err = database.Ping()
	infrastructure.Error(err)
	return database
}

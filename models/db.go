package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dataSourceName string){
	var err error
	db, err = sql.Open("mysql", dataSourceName)

	defer db.Close()

	if err != nil{
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
}
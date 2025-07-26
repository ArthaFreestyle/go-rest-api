package app

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB{
	db,err := sql.Open("mysql","root@tcp(localhost:3306)/go_restful_api")
	if err != nil{
		panic(err)
	}
	
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}
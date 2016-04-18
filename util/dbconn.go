package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetDbConnetion() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/sy_vendor?charset=utf8")
	if err != nil {
		fmt.Println("\nconnection mysql error")
	}
	//defer db.Close()
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Println("\nopen mysql error ", err)
	}
	return db
}

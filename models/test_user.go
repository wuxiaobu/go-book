package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func PrintUser() {
	db, err := sql.Open("mysql", "root:wzx3410@tcp(127.0.0.1:3306)/company?charset=utf8")
	defer db.Close()
	if err != nil {
		log.Println(err)
		return
	}
	stmt, err := db.Prepare("select id, name from company limit 3")
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	defer stmt.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("create table  if not exists foo (id integer)")
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}

	_, err = db.Exec("insert into foo(id) values(123)")
	if err != nil {
		fmt.Println("Failed to insert record:", err)
		return
	}

	rows, err := db.Query("select count(id) from foo")
	if err != nil {
		fmt.Println("Failed to select records:", err)
	}
	defer rows.Close()

	rows.Next()

	var result int
	rows.Scan(&result)
	fmt.Println("Rowcount: ", result)
}

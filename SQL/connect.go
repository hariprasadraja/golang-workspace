package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connectSQL()
	defer db.Close()

	query := "SELECT * From shop;"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var data struct {
			article int
			dealer  string
			price   string
		}

		rows.Scan(&data.article, &data.dealer, &data.price)
		fmt.Printf("%+v\n", data)
	}
}

func connectSQL() *sql.DB {
	db, err := sql.Open("mysql", "root:hariprasad@tcp(127.0.0.1:3306)/menagerie")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic("sorry, failed to connect db")
	}

	return db
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("hoge")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("SELECT name FROM todo WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("2").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

// TODO sqlite3
// TODO create table todo
// TODO
// TODO

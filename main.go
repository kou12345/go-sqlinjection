package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
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

	// rows, err := stmt.Query("2")
	// database/sqlパッケージがエスケープ処理しているから大丈夫
	rows, err := stmt.Query("'1' OR '1' = '1'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		name := ""
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

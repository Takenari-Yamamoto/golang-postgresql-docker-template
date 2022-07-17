package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age int
}

func main() {

	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// 複数データの取得
	cmd := "DELETE FROM persons WHERE name = ?"
	_ ,err := Db.Exec(cmd,"hanako")

	if err != nil {
		log.Fatalln(err)
	}

}

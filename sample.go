package main

import (
	"database/sql"
	"fmt"
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
	cmd := "SELECT * FROM persons"
	// Query は条件に合うものを全て取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next(){
		var p Person
		err := rows.Scan(&p.Name, &p.Age)

		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

}

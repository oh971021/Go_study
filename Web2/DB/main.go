package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open이라는 함수를 사용한다. 첫번째로 드라이버명 (우리는 mysql 사용)
	// 두번째로는 Connection
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id string
	var name string
	rows, err := db.Query("SELECT id, name FROM test1 WHERE id >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
}

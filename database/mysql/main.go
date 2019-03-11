package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@:123456@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var (
		id   int
		name string
	)

	rows, err := db.Query("SELECT * FROM user WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// 必须把rows 里面的内容读完，否则连接不会被释放
	for rows.Next() {
		err := rows.Scan(&id, name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

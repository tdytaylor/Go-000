学习笔记

package main

import (
"database/sql"
"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
db, err := sql.Open("mysql", "user:password@/dbname")
if err != nil {
panic(err)
}
// See "Important settings" section.
db.SetConnMaxLifetime(time.Minute * 3)
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(10)
}

func dao(name string) (interface{}, error) {
db, err := sql.Open("mysql", "user:password@/dbname")

	type Person struct {
		name string
	}

	var person = Person{}

	err = db.QueryRow("select name from users where id = ?", 1).Scan(&person)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return person, nil
}

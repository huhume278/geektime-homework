package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	name, err := querySomething()
	if err != nil {
		log.Fatal(err)
	}
	println("name:", name)

}

func querySomething() (string, error) {
	db, err := sql.Open("mysql", "user:pwd@tcp(localhost:3306)/go_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var name string
	err = db.QueryRow("select name from users where id = ?", 1).Scan(name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			return "", err
		}
	}
	return name, nil
}

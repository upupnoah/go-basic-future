package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	_, err := sql.Open("postgres", "user=root dbname=gocds password=rootroot sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
}

//import (
//"database/sql"
//_ "github.com/lib/pq"
//)
//
//func main() {
//	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	age := 21
//	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
//	...
//}

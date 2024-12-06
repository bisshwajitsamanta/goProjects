package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

var db *sql.DB
var o sync.Once

func main() {
	Run()
}

// Run - Init function has testability issues, o.Do makes sure that it runs at least once
func Run() {
	o.Do(func() {
		log.Println("Initializing database ...")
		var err error
		db, err = sql.Open("sqlite3", "./mysql.db")
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(db)
	})
}

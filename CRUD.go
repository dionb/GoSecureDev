package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.Conn

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!"))
}

func main() {
	db, err = sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

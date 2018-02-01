package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "hello:world@tcp(127.0.0.1:3306)/gosecdev")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", pingHandler)
	http.HandleFunc("/user", userHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	log.Println(http.ListenAndServe(":8080", nil))
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!"))
}

//STARTselect OMIT
func userHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method { // OMIT
	case "GET": // OMIT
		id := req.URL.Query().Get("id")
		var username string
		err := db.QueryRow("SELECT username FROM users WHERE id=?", id).Scan(&username) // HL
		if err != nil {
			rw.WriteHeader(500)
			log.Println(err.Error())
			rw.Write([]byte("An unknown error has occurred"))
			return
		}
		rw.Write([]byte(username))
		//STOPselect OMIT
	case "POST":
		//STARTinsert OMIT

		//STOPinsert OMIT
	}
}

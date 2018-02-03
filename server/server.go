package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

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
	http.HandleFunc("/play", playHandler)
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
		// func userHandler(rw http.ResponseWriter, req *http.Request) {
		reqBytes, _ := ioutil.ReadAll(req.Body)
		data := make(map[string]string) // HL
		json.Unmarshal(reqBytes, &data) // HL
		if name, ok := data["username"]; ok {
			_, err := db.Query("Insert into users (username) values (?)", name) // HL
			if err != nil {                                                     // OMIT
				log.Println(err.Error())                          // OMIT
				rw.WriteHeader(500)                               // OMIT
				rw.Write([]byte("A database error has occurred")) // OMIT
				return                                            // OMIT
			} // OMIT
			// error handling as before
			rw.WriteHeader(http.StatusOK)
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		//STOPinsert OMIT
	}
}

//STARTexec OMIT
func playHandler(rw http.ResponseWriter, req *http.Request) {
	reqBytes, _ := ioutil.ReadAll(req.Body)
	data := make(map[string]string)
	json.Unmarshal(reqBytes, &data)
	if name, ok := data["playlist"]; ok {
		rw.WriteHeader(http.StatusOK)
		cmd := exec.Command("ffplay", name) // HL
		err := cmd.Run()
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
	rw.WriteHeader(http.StatusBadRequest)
}

//STOPexec OMIT

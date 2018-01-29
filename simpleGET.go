package main

import (
	"log"
	"net/http"
)

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!"))
}

func main() {
	http.HandleFunc("/", pingHandler)
	log.Println(http.ListenAndServe(":8080", nil))
}

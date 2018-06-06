package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println(http.ListenAndServe("0.0.0.0:9999", nil))
}

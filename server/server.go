package server

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", pingHandler)
	log.Println(http.ListenAndServe(":8080", nil))
}

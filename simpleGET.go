package main

import (
	"log"
	"net/http"

	"github.com/gocraft/web"
)

func pingHandler(rw web.ResponseWriter, req *web.Request) {
	rw.Write([]byte("pong"))
}

func main() {
	router := web.New(struct{}{})
	router.Get("ping", pingHandler)
	log.Println(http.ListenAndServe(":8080", router))
}

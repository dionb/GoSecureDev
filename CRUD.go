package main

import (
	"net/http"
)

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!"))
}

func main() {

}

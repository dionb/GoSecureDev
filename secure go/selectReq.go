package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//START OMIT
	resp, _ := http.Get("http://localhost:8080/user?id=2")

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	//STOP OMIT
}

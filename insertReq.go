package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//START OMIT
	user := struct {
		Name string `json:"username"`
	}{
		"George",
	}
	reqBytes, _ := json.Marshal(user) // HL
	buf := bytes.NewBuffer(reqBytes)
	resp, _ := http.Post("http://localhost:8080/user", "application/json", buf)
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	//STOP OMIT
}

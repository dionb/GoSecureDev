package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func main() {
	//START OMIT
	body := map[string]string{
		"playlist": "public/streamtest/master.m3u8",
	}
	reqBytes, _ := json.Marshal(body)
	buf := bytes.NewBuffer(reqBytes)
	http.Post("http://localhost:8080/play", "application/json", buf)
	//STOP OMIT
}

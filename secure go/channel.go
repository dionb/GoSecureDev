package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

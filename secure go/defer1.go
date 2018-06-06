package main

import "fmt"

func main() {
	// START OMIT
	defer fmt.Println("A")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// STOP OMIT
}

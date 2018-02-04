package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// START OMIT
	hash := []byte("")
	password := []byte("password")
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("passwords match")
	// STOP OMIT
}

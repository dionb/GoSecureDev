package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// START OMIT

	hashCost := 10
	pass := []byte("CorrectHorseBatteryStaple")
	hash, err := bcrypt.GenerateFromPassword(pass, hashCost)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(hash))

	// STOP OMIT
}

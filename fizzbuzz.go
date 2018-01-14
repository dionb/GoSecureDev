package main

import (
	"errors"
	"fmt"
)

func main() {
	fizzbuzz(22)
}

// START OMIT
func fizzbuzz(n int) (highest int, err error) {
	subs := map[int]string{
		3: "fizz",
		7: "buzz",
	}
	if lessThanSmallestKey(n, subs) {
		return 0, errors.New("error: input is less than smallest key")
	}
	for i := 1; i <= n; i++ {
		found := false
		for j, k := range subs {
			if i%j == 0 {
				fmt.Print(k)
				highest = i
				found = true
			}
		}
		if !found {
			fmt.Print(i)
		}
		fmt.Println("")
	}
	return
}

// STOP OMIT

func lessThanSmallestKey(n int, subs map[int]string) bool {
	for i := range subs {
		if i < n {
			return false
		}
	}
	return true
}

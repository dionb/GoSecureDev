package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// START OMIT
	type ColorGroup struct {
		ID      int `json:"id"`
		Name    string
		Colors  []string
		hidden  string
		Empty   string `json:",omitempty"`
		Omitted string `json:"-"`
	}
	group := ColorGroup{
		ID:      1,
		Name:    "Reds",
		Colors:  []string{"Crimson", "Red", "Ruby", "Maroon"},
		hidden:  "You can't see me",
		Omitted: "Omitted",
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	// STOP OMIT
}

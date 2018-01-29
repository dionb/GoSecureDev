package main_test

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/golang/example/stringutil"
)

// func main() {
// cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
// cmd.Env = append(os.Environ(), "BE_CRASHER=1")
// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr
// err := cmd.Run()
// if err != nil {
// 	log.Println(err.Error())
// }
// }

// START OMIT
func TestReverse(t *testing.T) {
	test := "Hello World"
	expect := "dlroW olleH"
	if stringutil.Reverse(test) != expect {
		fmt.Printf("Expected '%s' got '%s'", test, expect)
		t.Fail()
	}
	t.Log("Reverse succeeded")
}

func TestRandom(t *testing.T) {
	data := make([]byte, 7)
	expect := []byte("REALLY?")
	rand.Read(data)
	if bytes.Compare(data, expect) != 0 {
		fmt.Printf("Expected %v got %v \n", expect, data)
		t.Fail()
	}
}

// STOP OMIT

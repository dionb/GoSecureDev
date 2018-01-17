package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

var numCooks, numServers = 5, 2

type pizza string

func makePizza(bench chan<- pizza, closing chan bool, working *sync.WaitGroup) {
	var p pizza
	defer working.Done()
	for {
		p = pizza(uuid.New().String()[:2])
		<-time.After(time.Second * 2)
		fmt.Printf("Pizza         %s is ready\n", p)
		select {
		case bench <- p:
		case <-closing:
			fmt.Printf("throwing out %s\n", p)
			return
		case <-time.After(time.Second * 10):
			fmt.Printf("Pizza         %s got cold, I'm off to make a new one\n", p)
		}
	}
}

func servePizza(bench <-chan pizza, closing chan bool, working *sync.WaitGroup) {
	defer working.Done()
	for {
		<-time.After(time.Millisecond * time.Duration(rand.Int63n(9000)))
		select {
		case <-closing:
			return
		case p := <-bench:
			fmt.Printf("Serving pizza %s to a customer\n", p)
		}
	}
}

func main() {
	closing := make(chan bool)
	bench := make(chan pizza)
	working := &sync.WaitGroup{}
	reader := bufio.NewReader(os.Stdin)

	working.Add(numCooks + numServers)
	fmt.Println("Press enter to close the pizza store")

	for i := 0; i < numCooks; i++ {
		go makePizza(bench, closing, working)
	}
	for i := 0; i < numServers; i++ {
		go servePizza(bench, closing, working)
	}

	reader.ReadString('\n')
	fmt.Println("We're closing, everyone out!")
	close(closing)
	working.Wait()
	fmt.Println("Goodnight!")
}

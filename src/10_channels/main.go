package main

import (
	"fmt"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println(num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, a int, b int) {
	result <- a + b
}

func main() {
	// This will through deadlock error because we are trying to send value to unbuffered channel without a concurrent receiver
	/* messageChan := make(chan string) // create a channel
	messageChan <- "Hello" // add "Hello" to channel
	fmt.Println(<-messageChan) // read from channel */

	// This will work as there is a concurrent receiver (thread go routine)
	/* numChan := make(chan int)
	go processNum(numChan)
	for {
		numChan <- rand.Intn(100)
	} */

	// Receiving from channel
	result := make(chan int)
	go sum(result, 1, 2)
	fmt.Println(<-result)
}

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

func task(done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println("Processing....")
}

func worker(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Email received: ", email)
	}
	fmt.Println(<-emailChan)
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
	/* result := make(chan int)
	go sum(result, 1, 2)
	fmt.Println(<-result) */

	// *** Signaling that a task completed in channels --> Go routing synchronizer
	/* done := make(chan bool)
	go task(done)
	<-done // this will block the call until it feds true in it */

	// Worker using channels - Queue Systems
	emailChan := make(chan string, 100) // buffered channel as we had already giving the size so it's a buffered channel no blocking
	done := make(chan bool)
	go worker(emailChan, done)
	for i := 0; i < 100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("All emails sent") // See logs after this line worker starts processing
	close(emailChan)
	<-done
}

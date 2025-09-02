// Channels are used to pass data between go routines and main thread, and can also pass in between go routines
package channels

import (
	"fmt"
	"time"
)

// Note: In function ar
// <-chan T : It means read only channel
// chan<- T : It means write only channel
// chan T   : Can both read and write channel
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Running...")
		}
	}
}

func Channels() {
	// 1. recieve data from a channel and print it - output: data
	/* myChannel := make(chan string)
	go func() {
		myChannel <- "data"
	}()
	msg := <-myChannel
	fmt.Println(msg) */

	// 2. prints the data coming from the fastest channel and exist - output: data1 or data2
	/* myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data - 1"
	}()
	go func() {
		anotherChannel <- "data - 2"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	} */

	// 3. Unbuffered vs Buffered Channels
	// Unbuffered Channel - It's synchronous, when a go routine sends data into an unbuffered channel it will go into pending state until the data in the channel is received by the reciever - think of it like(not exactly) a channel with size one(not a correct depection but can think of it like that)
	// Buffered Channel - It's asynchronous until the channel is full, If it's full it will also go into waiting stage and wait for the receiver to read
	/* charsChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}
	for _, s := range chars {
		charsChannel <- s
	}
	close(charsChannel)

	for result := range charsChannel {
		fmt.Println(result)
	} */

	// 4. Handling go routine leak with done channel - It makes the parent channel kill the useless go routines

	done := make(chan bool)
	go doWork(done)
	time.Sleep(time.Second * 3)
	close(done)
}

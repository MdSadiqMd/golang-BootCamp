package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // defer keyword func will run after the all tasks in function are done
	fmt.Println("Task ", id, " is running")
}

// go routine means we run the task using different threads
// Everything is good but to know whether the concurrent asynchronous multithreading task using go routine had completed running or not we use waitgroups
func main() {
	var waitgroup sync.WaitGroup
	for i := 0; i < 100; i++ {
		waitgroup.Add(1)
		go task(i, &waitgroup) // This go keyword is used to create go routine, this will run paralllely thus we are unable to see anything printing thus to see we add a sleep
	}
	waitgroup.Wait()
	// time.Sleep(time.Second * 2)
}

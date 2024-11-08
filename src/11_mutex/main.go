package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // This mutex will be used to prevent race conditions
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock() // It's better to lock only the critical section
	p.views += 1
}

func main() {
	myPost := post{views: 0}
	// This is ok as it runs synchronously
	/* for i:=0; i<100; i++ {
		myPost.incrementViews()
	} */

	// To run asynchronously and tracking using waitgroups
	// When running concurrently we undergo race conditions thus we use mutexes
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		myPost.incrementViews(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)
}

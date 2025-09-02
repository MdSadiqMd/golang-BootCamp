// context is used to manage concurrent operations in request lifecycle
// It is immutable, cannot by cancelled by child procceses, child proccess can create a new context from existing context and can pass it to their child process, this they can delete
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// parent context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	generator := func(dataItem string, stream chan interface{}) {
		for {
			select {
			case <-ctx.Done():
				return
			case stream <- dataItem:
			}
		}
	}

	infiniteApples := make(chan interface{})
	go generator("🍏", infiniteApples)

	infiniteOranges := make(chan interface{})
	go generator("🍊", infiniteOranges)

	infiniteWatermelon := make(chan interface{})
	go generator("🍉", infiniteWatermelon)

	wg.Add(1)
	go func1(ctx, &wg, infiniteApples)

	func2 := genericFunc
	func3 := genericFunc

	wg.Add(1)
	go func2(ctx, &wg, infiniteOranges)
	wg.Add(1)
	go func3(ctx, &wg, infiniteWatermelon)

	wg.Wait()
}

func func1(ctx context.Context, parentWg *sync.WaitGroup, stream <-chan interface{}) {
	defer parentWg.Done()
	var wg sync.WaitGroup

	doWork := func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-stream:
				if !ok {
					fmt.Println("channel closed")
					return
				}
				fmt.Println(d)
			}
		}
	}

	// creating context in child process from parent context
	newCtx, cancel := context.WithTimeout(ctx, time.Second*3) // when we use WithTimeout it cancels it's childs context after some time
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doWork(newCtx)
	}
	wg.Wait()
}

func genericFunc(ctx context.Context, wg *sync.WaitGroup, stream <-chan interface{}) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case d, ok := <-stream:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println(d)
		}
	}
}

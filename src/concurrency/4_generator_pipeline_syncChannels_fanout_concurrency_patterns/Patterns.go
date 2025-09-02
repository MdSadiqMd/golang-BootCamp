// generator - generates random data in the given threshold
// pipeline - dividing a multiple process into stages, all running synchronously and each stage is handled by a go routine
// As we're using Unbuffered channels the go routines will be in sync, such that if one stage takes time to process, the step before waits to send the message
// To improve the performance we use fanour architecture, which means we give spawn multiple go routines of primeFinder and make them as workers to process the data from the channel and again we use "fan In" to push all the data into one channel
package generatorpipelinesyncchannelsfanoutconcurrencypatterns

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"time"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case v, ok := <-stream:
				if !ok {
					return
				}
				taken <- v
			}
		}
	}()
	return taken
}

func primeFinder(done <-chan int, randIntSteam <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		if randomInt < 2 {
			return false
		}
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt, ok := <-randIntSteam:
				if !ok {
					return
				}
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()
	return primes
}

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)
	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:

			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()
	return fannedInStream
}

func ConcurrencyPattern() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	randNumFecther := func() int {
		n, err := rand.Int(rand.Reader, big.NewInt(5000000))
		if err != nil {
			return 0
		}
		return int(n.Int64())
	}

	// fan out
	CPUCount := runtime.NumCPU()
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ { // for loop to start go routines
		primeFinderChannels[i] = primeFinder(done, repeatFunc(
			done,
			randNumFecther,
		))
	}

	// fan in
	fannedInStream := fanIn(done, primeFinderChannels...)

	for rando := range take(done,
		primeFinder(
			done,
			fannedInStream,
		), 10) {
		fmt.Println(rando)
	}
	fmt.Println(time.Since(start))
}

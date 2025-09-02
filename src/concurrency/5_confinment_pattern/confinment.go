// When multiple go routines access a same resource, it might lead to an race condition
// Thus we use mutexes, we add a lock to the resource, and the go routine need to access it should acquire the lock
// But mutex make code slow, as only one go routine can acess the critical reource at a given time, it might make systems slow, suppose it's a company account in a fintech application
package confinment

import (
	"fmt"
	"sync"
	"time"
)

func process(data int) int {
	time.Sleep(time.Second * 2)
	return data * 2
}

// using mutex - works but slow

/* var lock sync.Mutex
func processData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()
	processedData := process(data)

	lock.Lock()
	*result = append(*result, processedData)
	lock.Unlock()
} */

// Using confinment - It uses pointers to maintain order, rather than taking values - no need of mutex
func processData(wg *sync.WaitGroup, resultDestination *int, data int) {
	defer wg.Done()
	processedData := process(data)
	*resultDestination = processedData
}

func Confinment() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data)
	}
	wg.Wait()

	fmt.Println(result)
	fmt.Println(time.Since(start))
}

// Go concurrency follows "Fork-Join Model", which means it create a child process from the main process via a fork, and they run independently
package routines

import (
	"fmt"
	"time"
)

func someFunc(param string) {
	fmt.Println(param)
}

func Concurrency() {
	// 1. the someFunc will block the execution of Main log - output: 1 Main
	/* someFunc("1")
	fmt.Println("Main") */

	// 2. the someFunc will get spin off in another thread and Main will be printed, as the code will be exicted we will be unable to see "1" - output: Main
	/* go someFunc("1")
	fmt.Println("Main") */

	// 3. the someFunc's will be run by multiple threads, the order is not guarenteed - output: 3 1 2 Main
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")
	time.Sleep(time.Second * 2)
	fmt.Println("Main")
}

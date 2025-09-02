// pipeline is nothing but, rather computing all at once, we compute the values in a format in each step via channels and move to next
package pipeline

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	channel := make(chan int)
	go func() {
		for _, n := range nums {
			channel <- n
		}
		close(channel) // we close channel here, because if we don't the slice will be empty and the process will not stop, as we're using a unbuffered channel
	}()
	return channel
}

func square(channel <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range channel {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func Pipeline() {
	nums := []int{2, 3, 4, 5, 6}
	firstStage := sliceToChannel(nums)
	secondStage := square(firstStage)
	for n := range secondStage {
		fmt.Println(n)
	}
}

package main

import (
	"fmt"
)

func main() {
	a, b, c := add(1, 2)
	fmt.Println(a, b, c)

	d := example(sub)
	fmt.Println(d)

	e := sum(1, 2, 3, 4, 5)
	fmt.Println(e)

	f := adder()
	fmt.Println(f(3))
}

func add(a, b int) (int, int, bool) { // if different type of parameter func add(a int,b string) int
	return a, b, true
}

func sub(a, b int) int {
	return a - b
}

// passing function as parameter to a function
func example(fn func(a int, b int) int) int {
	return fn(1, 2)
}

// variadic functions --> can take any number of arguments
func sum(a ...int) int { // if we want to receieve any type of arguments we use sum(a ...interface{})
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

// closures
func adder() func(int) int { // adder will return a function which retruns a int
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

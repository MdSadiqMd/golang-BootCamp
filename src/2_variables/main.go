package main

import "fmt"

func main() {
	var name string = "Sadiq"
	fmt.Println("Hello, " + name + "!")

	// constants - Immutable need to define at start
	const (
		port = 8080
		host = "localhost"
	)
	// port++; // not allowed
	fmt.Println("Host: ", host, " Port: ", port)

	// for loop - no while loop in golang
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// while using for loop
	j := 1
	for j < 5 {
		fmt.Println(j)
		j++
	}
}

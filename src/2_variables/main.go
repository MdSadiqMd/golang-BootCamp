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

	// range in go
	for i := range 11 {
		fmt.Println(i)
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("unknown")
		}
	}
	whoAmI(1)
	whoAmI("hello")
	whoAmI(true)
}

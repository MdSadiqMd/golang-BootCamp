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
	a := []int{1, 2, 3}
	for i, ele := range a { // range returns index and value
		fmt.Println(i, ele)
	}
	for i, b := range "hello" {
		fmt.Println(i, b) // prints index and ascii value
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

	var nums [4]int // or nums := [4]int{1,2}
	nums[0] = 1
	nums[1] = 2
	fmt.Println(nums)

	// slices - dynamic array
	var nums2 []int // if u not enter size in array then it is a slice
	fmt.Println("By default empty slice have nil value: ", nums2 == nil)
	nums2 = append(nums2, 0)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	fmt.Println(nums2)
	// we can add initial value in slices
	nums3 := make([]int, 3) // make slice of size three --> It will also increse size automatically if we add elements
	fmt.Println(nums3)

	// maps - key value pair
	var map1 map[string]int // or map1 := make(map[string]int) --> Syntax --> variable map[value type]key type
	map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	fmt.Println(map1)
	delete(map1, "one")
	fmt.Println(map1)
	clear(map1)
	fmt.Println(map1)

	_, ok := map1["one"]
	if ok {
		fmt.Println("one is present")
	} else {
		fmt.Println("one is not present")
	}
}

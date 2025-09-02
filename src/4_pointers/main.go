package main

import (
	"fmt"
)

func main() {
	var age int // Initialise a variable with type Int
	age = 5     // Assign a value to it, this get stored in memory
	fmt.Println(age)
	var pointer *int      // Initiliase a pointer in the memory with the type you're gonna reference value too, It's int because "age" is int
	fmt.Println(&age)     // Print the address of the variable "age"
	pointer = &age        // Assign the address of the variable "age" to pointer
	fmt.Println(*pointer) // Print the value to which the pointer is pointing to

	*pointer = 10 // edit the value of the variable via pointer - It's called dereferencing
	fmt.Println(pointer, *pointer, age)
}

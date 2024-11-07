package main

import (
	"fmt"
)

func changeNum(num int) {
	num = 2
}

func changeNumByValue(num *int) {
	*num = 2
}

func main() {
	num := 1
	changeNum(num) // This is pass by value so value won't change
	fmt.Println(num)
	changeNumByValue(&num) // This is pass by reference so we can change value
	fmt.Println(num)
}

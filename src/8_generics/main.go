package main

import (
	"fmt"
)

// With generics we can give any type as input when we don't know the type at compile time
func printSlice[T int | string, V string](items []T, name V) { // or func printSlice[T any](s []T) // or func printSlice[T interface{}](s []T) // or func printSlice[T comparable](s []T)
	for _, item := range items {
		fmt.Println(item, name)
	}
}

func main() {
	nums := []int{1, 2, 3}
	printSlice(nums, "name")
	names := []string{"sadiq", "Mohammad"}
	printSlice(names, "Sadiq")
}

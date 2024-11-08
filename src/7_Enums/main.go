package main

import "fmt"

// Creating an custom type
type OrderStatus string

// Enum is like type in typescript defined set of predefined values which are only allowed
// There is no Enum in go we acheive it using const
/* const ( // This will be compiled as integer as we have given iota
	Received  OrderStatus = iota // <name> <type> = <initial value>
	Confirmed                    // This Confirmed value will be 1 as upper Received given value as iota
	Shipped
	Delivered
) */

// If we want to add values other than indices/integers
const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Shipped               = "shipped"
	Delivered             = "delivered"
)

func changeStatus(status OrderStatus) {
	fmt.Println("Changed status to: ", status)
}
func main() {
	changeStatus(Shipped)
}

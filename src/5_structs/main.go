package main

import (
	"fmt"
	"time"
)

type Order struct { // struct is a collection of fields --> Syntax --> name struct
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func main() {
	myOrder := Order{ // (or) var myOrder Order
		id:     "123",
		amount: 100,
		status: "pending",
		// createdAt: time.Now(),
	}
	myOrder.id = "456"
	myOrder.amount = 200
	myOrder.status = "completed"
	myOrder.createdAt = time.Now()
	fmt.Println(myOrder)

	// Accesing struct using function
	myOrder.changeStatus("test Status")
	fmt.Println(myOrder)

	myOrder.getOrderStatus()
	fmt.Println(myOrder)

	// Constructor
	newOrder := NewOrder("123", 100, "pending")
	fmt.Println(newOrder)
}

// connecting fucntion with struct
// Note: Pass by Value not reference and the remainign de referencing in o.status=status will be handled by go
func (o *Order) changeStatus(status string) { // The receiver type is Order in go we use convention like <first letter of struct> struct Name`
	o.status = status
}

// Note: No need to give pointer as we are not changing the value
func (o Order) getOrderStatus() string {
	return o.status
}

// Constructor are not present in go thus we make them
func NewOrder(id string, amount float32, status string) *Order {
	newOrder := Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &newOrder
}

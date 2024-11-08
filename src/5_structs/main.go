package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type Order struct { // struct is a collection of fields --> Syntax --> name struct
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer  customer // or can just write customer
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

	// Making struct directly this is used in suituation when we only want to use one time
	language := struct {
		name   string
		isGood bool
	}{"sadiq", true}
	fmt.Println(language)

	// Constructor
	newOrder := NewOrder("123", 100, "pending")
	fmt.Println(newOrder)

	// Struct with struct embedding
	newOrder2 := Order{
		id:     "123",
		amount: 100,
		status: "pending",
	}
	fmt.Println(newOrder2)

	newOrder3 := Order{
		id:     "123",
		amount: 100,
		status: "pending",
		customer: customer{
			name:  "sadiq",
			phone: "123",
		},
	}
	fmt.Println(newOrder3)
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

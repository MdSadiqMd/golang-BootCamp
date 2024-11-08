package main

import (
	"fmt"
)

type paymenter interface { // usually we attach "er" in last of interface name
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Paid using razorpay: ", amount)
}

type stripe struct{}

func (r stripe) pay(amount float32) {
	fmt.Println("Paid using stripe: ", amount)
}

func main() {
	fmt.Println("Hello, World!")
	/* razorpayPaymentGw:=payment{
		gateway: razorpay{},
	}
	razorpayPaymentGw.makePayment(100) */
	// or
	// razorpayGateway := razorpay{}
	stripeGateway := stripe{}
	newPayment := payment{
		gateway: stripeGateway,
	}
	newPayment.makePayment(100)
}

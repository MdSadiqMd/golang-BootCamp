package main

import "fmt"

type purchasable interface {
	caluclatePrice() int64
}

var cart []purchasable

func addToCart(products ...purchasable) {
	cart = append(cart, products...)
}

func getCartTotal() int64 {
	var total int64
	for _, product := range cart {
		total += product.caluclatePrice()
	}
	return total
}

func main() {
	myShirt := Shirt{ProductDetails{Price: 5000, Brand: "bata"}, "XL", "Blue"}
	myMonitor := Monitor{ProductDetails{Price: 30000, Brand: "LG"}, "32 inch", "4k"}
	myDrink := Drinks{ProductDetails{Price: 200, Brand: "limcon"}, "2000", "red"}

	addToCart(myShirt, myMonitor, myDrink)
	fmt.Println(getCartTotal())
}

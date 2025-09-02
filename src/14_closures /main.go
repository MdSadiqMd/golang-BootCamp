// Here we have a function called as activateGiftCard, which returns a function that detects the amount called debitFunc
// now we can call the debitFunc and detect the amounts but, the function only returns the debitFunc right?how can we get the access of "amount" variable?
// Turns out golang sees that debitFunc is inheriently dependent on "amount" variable so it stores the varaible and the function in an enclosed unit in the memory called as "clousure"
// And every instance we create from the function is an different entity, so the "amount" variable is different for each other, thus we get answers as 90 and 95
package main

import "fmt"

func activateGiftCard() func(int) int {
	amount := 100
	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
	return debitFunc
}

func main() {
	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()
	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard2(5))
}

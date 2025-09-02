package main

type Drinks struct {
	ProductDetails
	Year string
	Kind string
}

func (s Drinks) caluclatePrice() int64 {
	drinkTax := float64(s.Price) * .23
	return s.Price + int64(drinkTax)
}

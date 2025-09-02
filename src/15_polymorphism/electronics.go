package main

type Monitor struct {
	ProductDetails
	Size       string
	Resolution string
}

func (s Monitor) caluclatePrice() int64 {
	electronicsTax := float64(s.Price) * .30
	return s.Price + int64(electronicsTax)
}

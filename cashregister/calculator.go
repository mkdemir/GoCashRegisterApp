package cashregister

func calculatePrice(item Item) float64 {
	discountedPrice := item.Price - (item.Price * float64(item.Discount) / 100)
	return discountedPrice
}

// TotalPrice calculates the total.
func TotalPrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += calculatePrice(item)
	}
	return total
}

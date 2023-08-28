package main

import (
	"GoCashRegisterApp/cashregister"
	"fmt"
)

func main() {
	item1 := cashregister.Item{
		Name:     "Elma",
		Price:    0.75,
		Discount: 10,
	}
	item2 := cashregister.Item{
		Name:  "Portakal",
		Price: 1.00,
	}

	items := []cashregister.Item{item1, item2}

	for _, item := range items {
		fmt.Println(item.Description())
	}

	total := cashregister.TotalPrice(items)
	fmt.Printf("Toplam Fiyat: %.2f TL\n", total)
}

package cashregister

import "fmt"

type Item struct {
	Name     string
	Price    float64
	Discount int
}

func (item Item) Description() string {
	if item.Discount > 0 {
		discountedPrice := calculatePrice(item)
		return fmt.Sprintf("%s - %.2f TL (%d%% indirimle %.2f TL)", item.Name, item.Price, item.Discount, discountedPrice)
	}
	return fmt.Sprintf("%s - %.2f TL", item.Name, item.Price)
}

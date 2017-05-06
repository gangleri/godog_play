package main

func NewBasket() *Basket {
	return &Basket{
		products: make(map[string]float64),
	}
}

type Basket struct {
	products map[string]float64
}

func (b *Basket) AddItem(productName string, price float64) {
	b.products[productName] = price
}

func (b *Basket) GetBasketSize() int {
	return len(b.products)
}

func (b *Basket) GetBasketTotal() float64 {
	basketTotal := 0.00
	shippingPrice := 0.00

	for _, value := range b.products {
		basketTotal += value
	}

	basketTotal = basketTotal * 1.2

	if basketTotal <= 10 {
		shippingPrice = 3
	} else {
		shippingPrice = 2
	}

	return basketTotal + shippingPrice
}
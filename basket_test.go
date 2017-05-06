package main

import (
	"github.com/DATA-DOG/godog"
	"fmt"
)

type shopping struct {
	shelf *Shelf
	basket *Basket
}

func (sh *shopping) addToBasket(productName string) (err error) {
	sh.basket.AddItem(productName, sh.shelf.GetProductPrice(productName))
	return
}

func (sb *shopping) addProduct(productName string, price float64) (err error) {
	sb.shelf.AddProduct(productName, price)
	return
}

func (sb *shopping) iShouldHaveProductsInTheBasket(productCount int) error {
	if sb.basket.GetBasketSize() != productCount {
		return fmt.Errorf(
			"expected %d products but there are %d",
			sb.basket.GetBasketSize(),
			productCount,
		)
	}

	return nil
}

func (sb *shopping) theOverallBasketPriceShouldBe(basketTotal float64) error {
	if sb.basket.GetBasketTotal() != basketTotal {
		return fmt.Errorf(
			"expected basket total to be %.2f but it is %.2f",
			basketTotal,
			sb.basket.GetBasketTotal(),
		)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {

	sh := &shopping{}

	s.BeforeScenario(func(interface{}) {
		sh.shelf = NewShelf()
		sh.basket = NewBasket()
	})

	s.Step(`^there is a "([a-zA-Z\s]+)", which costs £(\d+)$`, sh.addProduct)
	s.Step(`^I add the "([^"]*)" to the basket$`, sh.addToBasket)
	s.Step(`^I should have (\d+) products in the basket$`, sh.iShouldHaveProductsInTheBasket)
	s.Step(`^the overall basket price should be £(\d+)$`, sh.theOverallBasketPriceShouldBe)
}

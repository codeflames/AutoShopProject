package main

import (
	"fmt"

	"github.com/codeflames/AutoShopProject/shop"
)

func main() {

	shop1 := shop.NewShop()

	item1 := &shop.Product{
		Id:              1,
		Name:            "Civic 2019",
		ProductPrice:    100.00,
		ProductQuantity: 10,
		ProductStatus:   true,
		Car: shop.Car{
			Model:   "Civic",
			Brand:   "Honda",
			Year:    2019,
			Color:   "Red",
			IsUsed:  false,
			Mileage: 0,
		},
	}

	item2 := &shop.Product{
		Id:              2,
		Name:            "Corolla 2018",
		ProductPrice:    200.00,
		ProductQuantity: 5,
		ProductStatus:   true,
		Car: shop.Car{
			Model:   "Corolla",
			Brand:   "Toyota",
			Year:    2018,
			Color:   "Blue",
			IsUsed:  false,
			Mileage: 0,
		},
	}

	shop1.AddProducts(item1, item2)

	shop1.ListSoldProducts()

	fmt.Printf("Total unsold products price: %v", shop1.TotalUnsoldProductsPrice())
	fmt.Println()

	shop1.ListProducts()

	shop1.SellProduct(1, 2)

	shop1.ListSoldProducts()

	shop1.ListProducts()

	fmt.Println(shop1.TotalSales())

	shop1.SellProduct(1, 3)

	shop1.ListSoldProducts()

	shop1.ListProducts()

	fmt.Println(shop1.TotalSales())

	shop1.SellProduct(2, 2)

	shop1.ListSoldProducts()

	shop1.ListProducts()

	fmt.Println(shop1.TotalSales())
	shop1.SellProduct(2, 2)

	shop1.ListSoldProducts()

	shop1.ListProducts()

	fmt.Println(shop1.TotalSales())

	shop1.SellProduct(2, 2)

	shop1.ListSoldProducts()

	shop1.ListProducts()

	fmt.Println(shop1.TotalSales())
}

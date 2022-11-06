package shop

import "fmt"

type Car struct {
	Model   string
	Brand   string
	Year    int
	Color   string
	IsUsed  bool
	Mileage int
}

// a product struct that implements the ShopProduct interface
type Product struct {
	Id              int
	Name            string
	ProductPrice    float64
	ProductQuantity int
	ProductStatus   bool
	Car
}

// The store class should have attributes like
// Number of products in the store that are still up for sale
// Adding an Item to the store
// Listing all product items in the store
// Sell an item
// Show a list of sold items and the total ProductPrice

// store struct
type store struct {
	products []Product
	sold     []Product
}

// NewShop creates a new shop
func NewShop() *store {
	return &store{}
}

// add a product to a store
func (s *store) AddProduct(p *Product) {
	//check if product already exists and update it
	for _, prod := range s.products {
		if prod.Name == p.Name {
			fmt.Println("Product already exists")
			fmt.Println("Updating product")
			s.UpdateProduct(p)
			return
		}
	}
	s.products = append(s.products, *p)
}

// update product in store
func (s *store) UpdateProduct(p *Product) {
	for i, prod := range s.products {
		if prod.Name == p.Name {
			s.products[i] = *p
			return
		}
	}
}

// add multiple products to the store
func (s *store) AddProducts(ps ...*Product) {
	for _, p := range ps {
		s.AddProduct(p)
	}
}

// ListProducts lists all products in the store
func (s *store) ListProducts() {
	//check if store has products
	if len(s.products) == 0 {
		fmt.Println("No products in store")
		return
	} else {
		fmt.Println("All Available Products:")

		for _, p := range s.products {
			p.Display()
		}
		fmt.Println("====================================")
	}
}

// SellProduct sells a product and quantity is reduced by the quantity sold

// add product to sold list with the quantity sold
func (s *store) SellProduct(id int, q int) {
	for i, p := range s.products {
		if p.Id == id {
			if p.ProductStatus {
				if p.ProductQuantity >= q {

					p.ProductQuantity -= q
					s.products[i] = p
					s.checkSoldList(&p, q)
				} else {
					fmt.Println("Not enough quantity")
				}

				// check if product quantity is 0, if yes, update product status to false
				if p.ProductQuantity == 0 {
					p.updateStatus(false)
				}

				return
			} else {
				fmt.Println("Product out of stock")
				return
			}
		}
	}
	fmt.Println("Product not found")
}

// check if sold list has a product, if yes, add the quantity sold to the product in the list else add the product and the quantity sold to the list
func (s *store) checkSoldList(p *Product, q int) {
	for i, prod := range s.sold {
		if prod.Name == p.Name {
			s.sold[i].ProductQuantity += q
			return
		}
	}
	*p = Product{Id: p.Id, Name: p.Name, ProductPrice: p.ProductPrice, ProductQuantity: q}
	s.sold = append(s.sold, *p)
}

// ListSoldProducts lists all sold products
func (s *store) ListSoldProducts() {
	fmt.Println("Sold Products:")
	for _, p := range s.sold {
		p.Display()
	}
	fmt.Println("====================================")
}

// Total price of product left to sell in the store
func (p *store) TotalUnsoldProductsPrice() float64 {
	var total float64
	for _, prod := range p.products {
		total += prod.ProductPrice * float64(prod.ProductQuantity)
	}
	return total
}

// TotalSales returns the total sales
func (s *store) TotalSales() float64 {
	var total float64
	for _, p := range s.sold {
		total += p.Price() * float64(p.Quantity())
	}
	return total
}

// ProductId returns the product id
func (p *Product) ProductId() int {
	return p.Id
}

// Display displays the product
func (p *Product) Display() {
	fmt.Printf("Product: %s, ProductPrice: %.2f, ProductQuantity: %d", p.Name, p.ProductPrice, p.ProductQuantity)
	fmt.Println()
}

// update product status
func (p *Product) updateStatus(status bool) {
	p.ProductStatus = status
}

// ProductQuantity returns the ProductQuantity of the product
func (p *Product) Quantity() int {
	return p.ProductQuantity
}

// Price returns the Price of the product
func (p *Product) Price() float64 {
	return p.ProductPrice
}

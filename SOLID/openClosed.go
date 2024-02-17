package main

import "fmt"

// OCP : Open for extension, Closed for modifications
// Open for functions
type Size int

type Color int

const (
	small Size = iota
	medium
	large
)

const (
	red Color = iota
	green
	blue
)

type Product struct {
	Name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(product []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, val := range product {
		if val.color == color {
			result = append(result, &product[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	banana := Product{"Banana", green, medium}
	orange := Product{"Orange", blue, medium}

	products := []Product{apple, banana, orange}

	f := Filter{}

	for _, val := range f.FilterByColor(products, green) {
		fmt.Printf("%s %d ", val.Name, val.color)
	}
}

package main

import "fmt"

//Open Closed Principle

//This principle states that a class should be open for extension and closed for modification
//lets say we have a class of products, and we added a method in class for filtering these products
//we added method to filter by size and now your manager comes in and says please add
//filter by color as well now we copy/paste filter by size but passing color
//then your manager comes in again and says please add method for filter by color and size
//here we are violating open/closed principle

//PROBLEM

type Size int

const (
	small Size = iota
	medium
	large
)

type Color int

const (
	red Color = iota
	blue
	green
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.size == size && product.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

//Here this approach will make us add multiple filter method each time we require a new type of filtration

//SOLUTION
//OpenClosed Principle can also be called Specification Pattern

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c *ColorSpecification) IsSatisfied(p *Product) bool {
	return c.color == p.color
}

type SizeSpecification struct {
	size Size
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

type BetterFilter struct{}

func (betterFilter *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if spec.IsSatisfied(&product) {
			result = append(result, &products[i])
		}
	}
	return result
}

//Now lets say we want to satisfy two specification
//In this case we use composite pattern

type AndSpecification struct {
	first, second Specification
}

func (compositeSpec AndSpecification) IsSatisfied(product *Product) bool {
	return compositeSpec.first.IsSatisfied(product) && compositeSpec.second.IsSatisfied(product)
}

// MAIN FUNCTION
func main() {
	apple := Product{name: "Apple", color: green, size: small}
	tree := Product{name: "Tree", color: green, size: large}
	house := Product{name: "House", color: blue, size: large}

	products := []Product{apple, tree, house}
	filter := Filter{}
	fmt.Printf("Green Products [old] : \n")

	for _, product := range filter.FilterByColor(products, green) {
		fmt.Printf("- %s is green\n", product.name)
	}

	fmt.Printf("Green Products [new] : \n")
	greenSpec := ColorSpecification{color: green}
	bf := BetterFilter{}
	for _, product := range bf.Filter(products, &greenSpec) {
		fmt.Printf("- %s is green\n", product.name)
	}

	fmt.Printf("Large Green Products [new] : \n")
	largeSpec := SizeSpecification{size: large}
	lgSpec := AndSpecification{&greenSpec, &largeSpec}
	for _, product := range bf.Filter(products, &lgSpec) {
		fmt.Printf("- %s is green\n", product.name)
	}
}
